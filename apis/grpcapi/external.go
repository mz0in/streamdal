package grpcapi

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/streamdal/snitch-protos/build/go/protos"

	"github.com/streamdal/snitch-server/services/store"
	"github.com/streamdal/snitch-server/types"
	"github.com/streamdal/snitch-server/util"
	"github.com/streamdal/snitch-server/validate"
)

// ExternalServer implements the external GRPC API interface
type ExternalServer struct {
	GRPCAPI
	// Must be implemented in order to satisfy the protos ExternalServer interface
	protos.UnimplementedExternalServer
}

func (g *GRPCAPI) newExternalServer() *ExternalServer {
	return &ExternalServer{
		GRPCAPI: *g,
	}
}

func (s *ExternalServer) GetAll(ctx context.Context, req *protos.GetAllRequest) (*protos.GetAllResponse, error) {
	if err := validate.GetAllRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid get all request")
	}

	liveInfo, err := s.getAllLive(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get live info")
	}

	audiences, err := s.Options.StoreService.GetAudiences(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get audiences")
	}

	pipelines, err := s.getAllPipelines(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get pipelines")
	}

	configs, err := s.Options.StoreService.GetConfig(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get config")
	}

	configStrAudience := util.ConvertConfigStrAudience(configs)

	return &protos.GetAllResponse{
		Live:                   liveInfo,
		Audiences:              audiences,
		Pipelines:              pipelines,
		Config:                 configStrAudience,
		GeneratedAtUnixTsNsUtc: time.Now().UTC().UnixNano(),
	}, nil
}

func (s *ExternalServer) GetAllStream(req *protos.GetAllRequest, server protos.External_GetAllStreamServer) error {
	if err := validate.GetAllRequest(req); err != nil {
		return errors.Wrap(err, "invalid get all request")
	}

	llog := s.log.WithFields(logrus.Fields{
		"method": "GetAllStream",
	})

	// Generate initial GetAllResponse
	resp, err := s.GetAll(server.Context(), req)
	if err != nil {
		llog.Errorf("unable to complete initial GetAll: %s", err)
		return fmt.Errorf("unable to complete initial GetAll: %s", err)
	}

	// Send initial response
	if err := server.Send(resp); err != nil {
		llog.Errorf("unable to send initial GetAll response: %s", err)
		return fmt.Errorf("unable to send initial GetAll response: %s", err)
	}

	// Cleanup after Listen()
	requestID := util.CtxRequestId(server.Context())
	defer s.Options.PubSubService.Close(types.PubSubChangesTopic, requestID)

	sendInProgress := false

MAIN:
	for {
		select {
		case <-server.Context().Done():
			llog.Debug("client closed connection")
			break MAIN
		case <-s.Options.ShutdownContext.Done():
			llog.Debug("server shutting down")
			break MAIN
		case <-s.Options.PubSubService.Listen(types.PubSubChangesTopic, requestID):
			if sendInProgress {
				llog.Debug("send in progress, skipping changes message")
				continue
			}

			sendInProgress = true

			llog.Debug("launching goroutine to send delayed GetAllResponse")

			// Artificial update slowdown -- we do this to avoid spamming the
			// client with updates when there are many concurrent Listen() hits
			go func() {
				defer func() {
					sendInProgress = false
				}()

				time.Sleep(100 * time.Millisecond)

				llog.Debug("received changes message via pubsub")

				// Generate a GetAllResponse
				resp, err := s.GetAll(server.Context(), req)
				if err != nil {
					llog.Errorf("unable to complete GetAll: %s", err)
					return
				}

				// Send response
				if err := server.Send(resp); err != nil {
					llog.Errorf("unable to send GetAll response: %s", err)
					return
				}
			}()
		}
	}

	llog.Debug("closing stream")

	return nil
}

func (s *ExternalServer) getAllLive(ctx context.Context) ([]*protos.LiveInfo, error) {
	liveInfo := make([]*protos.LiveInfo, 0)

	liveData, err := s.Options.StoreService.GetLive(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get live data")
	}

	for _, v := range liveData {
		live := &protos.LiveInfo{
			Audiences: make([]*protos.Audience, 0),
		}

		// If register entry, fill out client info
		if v.Register {
			if err := validate.ClientInfo(v.Value); err != nil {
				s.log.Errorf("getAllLive: unable to validate client info for session id '%s': %v", v.SessionID, err)
				continue
			}

			// There might be no audience - service name + node name etc. should be in ClientInfo though!

			live.Client = v.Value

			liveInfo = append(liveInfo, live)

			// Audiences will get filled out later
		}
	}

	// No register entries == no one connected to any instances of snitch server
	if len(liveData) == 0 {
		return liveInfo, nil
	}

	// Have register entries - fill out audiences for each
	for _, li := range liveInfo {
		// Find all live entry audiences with same session ID (and are NOT register)
		for _, ld := range liveData {
			if *li.Client.XSessionId == ld.SessionID && !ld.Register {
				li.Audiences = append(li.Audiences, ld.Audience)
			}
		}
	}

	return liveInfo, nil
}

func (s *ExternalServer) getAllPipelines(ctx context.Context) (map[string]*protos.PipelineInfo, error) {
	gen := make(map[string]*protos.PipelineInfo)

	// Get all pipelines
	allPipelines, err := s.Options.StoreService.GetPipelines(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get pipelines")
	}

	for pipelineID, pipeline := range allPipelines {
		util.StripWASMFields(pipeline)

		gen[pipelineID] = &protos.PipelineInfo{
			Audiences: make([]*protos.Audience, 0),
			Pipeline:  pipeline,
			Paused:    make([]*protos.Audience, 0),
		}
	}

	// Get audience <-> pipeline mappings
	pipelineConfig, err := s.Options.StoreService.GetConfig(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get pipeline config")
	}

	// Update pipeline info with info about attached pipelines
	for aud, pipelineIDs := range pipelineConfig {
		for _, pipelineID := range pipelineIDs {
			if _, ok := gen[pipelineID]; ok {
				gen[pipelineID].Audiences = append(gen[pipelineID].Audiences, aud)
			}
		}
	}

	// Update pipeline info with state info
	pausedPipelines, err := s.Options.StoreService.GetPaused(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get paused pipelines")
	}

	for _, p := range pausedPipelines {
		if _, ok := gen[p.PipelineID]; ok {
			gen[p.PipelineID].Paused = append(gen[p.PipelineID].Paused, p.Audience)
		}
	}

	return gen, nil
}

func (s *ExternalServer) GetPipelines(ctx context.Context, req *protos.GetPipelinesRequest) (*protos.GetPipelinesResponse, error) {
	if err := validate.GetPipelinesRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid get pipelines request")
	}

	// Read all keys in "snitch_pipelines"
	pipelines, err := s.Options.StoreService.GetPipelines(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get pipelines")
	}

	// Convert map to slice
	pipelineSlice := make([]*protos.Pipeline, 0)

	for _, pipeline := range pipelines {
		pipelineSlice = append(pipelineSlice, pipeline)

		// Strip WASM fields (to save on b/w)
		util.StripWASMFields(pipeline)
	}

	return &protos.GetPipelinesResponse{
		Pipelines: pipelineSlice,
	}, nil
}

func (s *ExternalServer) GetPipeline(ctx context.Context, req *protos.GetPipelineRequest) (*protos.GetPipelineResponse, error) {
	if err := validate.GetPipelineRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid get pipeline request")
	}

	pipeline, err := s.Options.StoreService.GetPipeline(ctx, req.PipelineId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get pipelines")
	}

	// Strip WASM fields (to save on b/w)
	util.StripWASMFields(pipeline)

	return &protos.GetPipelineResponse{
		Pipeline: pipeline,
	}, nil
}

func (s *ExternalServer) CreatePipeline(ctx context.Context, req *protos.CreatePipelineRequest) (*protos.CreatePipelineResponse, error) {
	if err := validate.CreatePipelineRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid create pipeline request")
	}

	// Create ID for pipeline
	req.Pipeline.Id = util.GenerateUUID()

	// Populate WASM fields
	if err := util.PopulateWASMFields(req.Pipeline, s.Options.Config.WASMDir); err != nil {
		return nil, errors.Wrap(err, "unable to populate WASM fields")
	}

	if err := s.Options.StoreService.CreatePipeline(ctx, req.Pipeline); err != nil {
		return nil, errors.Wrap(err, "unable to store pipeline")
	}

	return &protos.CreatePipelineResponse{
		Message:    "Pipeline created successfully",
		PipelineId: req.Pipeline.Id,
	}, nil
}

func (s *ExternalServer) UpdatePipeline(ctx context.Context, req *protos.UpdatePipelineRequest) (*protos.StandardResponse, error) {
	if err := validate.UpdatePipelineRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	// Is this a known pipeline?
	if _, err := s.Options.StoreService.GetPipeline(ctx, req.Pipeline.Id); err != nil {
		if err == store.ErrPipelineNotFound {
			return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_NOT_FOUND, err.Error()), nil
		}

		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Re-populate WASM bytes (since they are stripped for UI)
	if err := util.PopulateWASMFields(req.Pipeline, s.Options.Config.WASMDir); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, errors.Wrap(err, "unable to repopulate WASM data").Error()), nil
	}

	// Update pipeline in storage
	if err := s.Options.StoreService.UpdatePipeline(ctx, req.Pipeline); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Pipeline exists - broadcast it as there might be snitch-servers that have
	// a client that has an active registration using this pipeline (and it should
	// get updated)
	if err := s.Options.BusService.BroadcastUpdatePipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: fmt.Sprintf("pipeline '%s' updated", req.Pipeline.Id),
	}, nil
}

func (s *ExternalServer) DeletePipeline(ctx context.Context, req *protos.DeletePipelineRequest) (*protos.StandardResponse, error) {
	if err := validate.DeletePipelineRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	// Does this pipeline exist?
	if _, err := s.Options.StoreService.GetPipeline(ctx, req.PipelineId); err != nil {
		if err == store.ErrPipelineNotFound {
			return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_NOT_FOUND, err.Error()), nil
		}

		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Pipeline exists, delete it
	if err := s.Options.StoreService.DeletePipeline(ctx, req.PipelineId); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Now broadcast delete
	if err := s.Options.BusService.BroadcastDeletePipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: fmt.Sprintf("pipeline '%s' deleted", req.PipelineId),
	}, nil
}

func (s *ExternalServer) AttachPipeline(ctx context.Context, req *protos.AttachPipelineRequest) (*protos.StandardResponse, error) {
	if err := validate.AttachPipelineRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	// Does this pipeline exist?
	if _, err := s.Options.StoreService.GetPipeline(ctx, req.PipelineId); err != nil {
		if err == store.ErrPipelineNotFound {
			return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_NOT_FOUND, err.Error()), nil
		}

		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	if err := s.Options.StoreService.AttachPipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Pipeline exists, broadcast attach
	if err := s.Options.BusService.BroadcastAttachPipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: fmt.Sprintf("pipeline '%s' attached", req.PipelineId),
	}, nil
}

func (s *ExternalServer) DetachPipeline(ctx context.Context, req *protos.DetachPipelineRequest) (*protos.StandardResponse, error) {
	if err := validate.DetachPipelineRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	// Does this pipeline exist?
	if _, err := s.Options.StoreService.GetPipeline(ctx, req.PipelineId); err != nil {
		if err == store.ErrPipelineNotFound {
			return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_NOT_FOUND, err.Error()), nil
		}

		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Pipeline exists, broadcast delete
	if err := s.Options.BusService.BroadcastDetachPipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// TODO: figure out a better method for this
	// Give time for all snitch server instances to receive detach command and send it to active sessions
	// before removing the config entry, otherwise getActivePipelineUsage() will return incorrect results
	go func() {
		time.Sleep(time.Second * 5)
		// Remove config entry
		if err := s.Options.StoreService.DetachPipeline(ctx, req); err != nil {
			s.log.Error(errors.Wrap(err, "unable to detach pipeline"))
		}
	}()

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: fmt.Sprintf("pipeline '%s' detached", req.PipelineId),
	}, nil
}

func (s *ExternalServer) PausePipeline(ctx context.Context, req *protos.PausePipelineRequest) (*protos.StandardResponse, error) {
	if err := validate.PausePipelineRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	// Does this pipeline exist?
	if _, err := s.Options.StoreService.GetPipeline(ctx, req.PipelineId); err != nil {
		if err == store.ErrPipelineNotFound {
			return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_NOT_FOUND, err.Error()), nil
		}

		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Can attempt to pause; PausePipeline() will noop if pipeline is already paused
	if err := s.Options.StoreService.PausePipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	if err := s.Options.BusService.BroadcastPausePipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: fmt.Sprintf("pipeline '%s' paused", req.PipelineId),
	}, nil
}

func (s *ExternalServer) ResumePipeline(ctx context.Context, req *protos.ResumePipelineRequest) (*protos.StandardResponse, error) {
	if err := validate.ResumePipelineRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	// Does this pipeline exist?
	if _, err := s.Options.StoreService.GetPipeline(ctx, req.PipelineId); err != nil {
		if err == store.ErrPipelineNotFound {
			return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_NOT_FOUND, err.Error()), nil
		}

		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Can attempt to resume; ResumePipeline() will noop if pipeline is already running
	if err := s.Options.StoreService.ResumePipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Pipeline exists, broadcast resume
	if err := s.Options.BusService.BroadcastResumePipeline(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: fmt.Sprintf("pipeline '%s' deleted", req.PipelineId),
	}, nil
}

func (s *ExternalServer) CreateNotification(ctx context.Context, req *protos.CreateNotificationRequest) (*protos.StandardResponse, error) {
	if err := validate.CreateNotificationRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	req.Notification.Id = util.StringPtr(util.GenerateUUID())

	if err := s.Options.StoreService.CreateNotificationConfig(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Notification config created",
	}, nil
}
func (s *ExternalServer) UpdateNotification(ctx context.Context, req *protos.UpdateNotificationRequest) (*protos.StandardResponse, error) {
	if err := validate.UpdateNotificationRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	if err := s.Options.StoreService.UpdateNotificationConfig(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Notification config update",
	}, nil
}

func (s *ExternalServer) DeleteNotification(ctx context.Context, req *protos.DeleteNotificationRequest) (*protos.StandardResponse, error) {
	if err := validate.DeleteNotificationRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	if err := s.Options.StoreService.DeleteNotificationConfig(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Notification config deleted",
	}, nil
}

func (s *ExternalServer) GetNotifications(ctx context.Context, req *protos.GetNotificationsRequest) (*protos.GetNotificationsResponse, error) {
	cfgs, err := s.Options.StoreService.GetNotificationConfigs(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get notification configs")
	}

	return &protos.GetNotificationsResponse{Notifications: cfgs}, nil
}

func (s *ExternalServer) GetNotification(ctx context.Context, req *protos.GetNotificationRequest) (*protos.GetNotificationResponse, error) {
	if err := validate.GetNotificationRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	cfg, err := s.Options.StoreService.GetNotificationConfig(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get notification config")
	}

	return &protos.GetNotificationResponse{Notification: cfg}, nil
}

func (s *ExternalServer) AttachNotification(ctx context.Context, req *protos.AttachNotificationRequest) (*protos.StandardResponse, error) {
	if err := validate.AttachNotificationRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	if err := s.Options.StoreService.AttachNotificationConfig(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Notification config attached",
	}, nil
}

func (s *ExternalServer) DetachNotification(ctx context.Context, req *protos.DetachNotificationRequest) (*protos.StandardResponse, error) {
	if err := validate.DetachNotificationRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	if err := s.Options.StoreService.DetachNotificationConfig(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Notification config detached",
	}, nil
}

func (s *ExternalServer) DeleteAudience(ctx context.Context, req *protos.DeleteAudienceRequest) (*protos.StandardResponse, error) {
	if err := validate.DeleteAudienceRequest(req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST, err.Error()), nil
	}

	if err := s.Options.StoreService.DeleteAudience(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	// Broadcast delete to other nodes so that they can emit an event for GetAllStream()
	if err := s.Options.BusService.BroadcastDeleteAudience(ctx, req); err != nil {
		return util.StandardResponse(ctx, protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR, err.Error()), nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Audience deleted",
	}, nil
}

func (s *ExternalServer) Tail(req *protos.TailRequest, server protos.External_TailServer) error {
	// Each tail request gets its own unique ID so that we can receive messages over
	// a unique channel from NATS
	req.Id = util.GenerateUUID()

	if err := validate.StartTailRequest(req); err != nil {
		return errors.Wrap(err, "invalid tail request")
	}

	// Get channel for receiving TailResponse messages that get shipped over NATS.
	// This should exist before TailRequest command is sent to the SDKs so that we are ready to receive
	// messages from NATS
	sdkReceiveChan := s.Options.PubSubService.Listen(req.Id, util.CtxRequestId(server.Context()))

	if err := s.Options.BusService.BroadcastTailRequest(context.Background(), req); err != nil {
		return errors.Wrap(err, "unable to broadcast tail request")
	}

	defer func() {
		req.Type = protos.TailRequestType_TAIL_REQUEST_TYPE_STOP
		if err := s.Options.BusService.BroadcastTailRequest(context.Background(), req); err != nil {
			s.log.Error(errors.Wrap(err, "unable to broadcast stop tail request"))
		}
	}()

	for {
		select {
		case <-server.Context().Done():
			s.log.Debug("frontend closed tail stream")
			return nil
		case <-s.Options.ShutdownContext.Done():
			s.log.Debug("server shutting down, exiting tail stream")
			return nil
		case msg := <-sdkReceiveChan:
			tr, ok := msg.(*protos.TailResponse)
			if !ok {
				s.log.Errorf("unknown message received from connected SDK session: %v", msg)
				continue
			}
			if err := server.Send(tr); err != nil {
				return errors.Wrap(err, "unable to send tail stream response")
			}
		}
	}
}

func (s *ExternalServer) Test(_ context.Context, req *protos.TestRequest) (*protos.TestResponse, error) {
	return &protos.TestResponse{
		Output: "Pong: " + req.Input,
	}, nil
}
