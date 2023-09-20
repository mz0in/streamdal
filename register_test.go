package snitch

import (
	"context"
	"sync"
	"testing"

	"github.com/google/uuid"

	"github.com/streamdal/snitch-protos/build/go/protos"
	"github.com/streamdal/snitch-protos/build/go/protos/shared"

	"github.com/streamdal/snitch-go-client/kv/kvfakes"
	"github.com/streamdal/snitch-go-client/logger/loggerfakes"
)

func TestHandleKVCommand(t *testing.T) {
	fakeKV := &kvfakes.FakeIKV{}

	s := &Snitch{
		config: &Config{
			Logger: &loggerfakes.FakeLogger{},
		},
		kv: fakeKV,
	}

	t.Run("create", func(t *testing.T) {
		cmd := &protos.KVCommand{
			Instructions: []*protos.KVInstruction{
				{
					Id:     uuid.New().String(),
					Action: shared.KVAction_KV_ACTION_CREATE,
					Object: &protos.KVObject{
						Key:   "test-key",
						Value: []byte("test-value"),
					},
				},
			},
			Overwrite: false,
		}

		if err := s.handleKVCommand(context.Background(), cmd); err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		if fakeKV.SetCallCount() != 1 {
			t.Errorf("Expected 1 call to Set, got %d", fakeKV.SetCallCount())
		}
	})

	t.Run("delete", func(t *testing.T) {
		cmd := &protos.KVCommand{
			Instructions: []*protos.KVInstruction{
				{
					Id:     uuid.New().String(),
					Action: shared.KVAction_KV_ACTION_DELETE,
					Object: &protos.KVObject{
						Key: "test-key",
					},
				},
			},
		}

		if err := s.handleKVCommand(context.Background(), cmd); err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		if fakeKV.DeleteCallCount() != 1 {
			t.Errorf("Expected 1 call to Delete, got %d", fakeKV.DeleteCallCount())
		}
	})

	t.Run("purge", func(t *testing.T) {
		cmd := &protos.KVCommand{
			Instructions: []*protos.KVInstruction{
				{
					Id:     uuid.New().String(),
					Action: shared.KVAction_KV_ACTION_DELETE_ALL,
				},
			},
		}

		if err := s.handleKVCommand(context.Background(), cmd); err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		if fakeKV.PurgeCallCount() != 1 {
			t.Errorf("Expected 1 call to Purge, got %d", fakeKV.PurgeCallCount())
		}
	})

}

func TestAttachPipeline(t *testing.T) {
	s := &Snitch{
		pipelinesMtx: &sync.RWMutex{},
		pipelines:    make(map[string]map[string]*protos.Command),
		config: &Config{
			Logger: &loggerfakes.FakeLogger{},
		},
	}

	if err := s.attachPipeline(context.Background(), nil); err != ErrEmptyCommand {
		t.Error("Expected ErrEmptyCommand, got nil")
	}

	aud := &protos.Audience{
		OperationName: "test-operation",
		ServiceName:   "test-service",
		OperationType: protos.OperationType_OPERATION_TYPE_PRODUCER,
		ComponentName: "test-component",
	}

	cmd := &protos.Command{
		Audience: aud,
		Command: &protos.Command_AttachPipeline{
			AttachPipeline: &protos.AttachPipelineCommand{
				Pipeline: &protos.Pipeline{
					Id: uuid.New().String(),
				},
			},
		},
	}

	err := s.attachPipeline(context.Background(), cmd)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if len(s.pipelines) != 1 {
		t.Errorf("Expected 1 audience, got %d", len(s.pipelines))
	}

	if len(s.pipelines[audToStr(aud)]) != 1 {
		t.Errorf("Expected 1 pipeline, got %d", len(s.pipelines[audToStr(aud)]))
	}
}

func TestDetachPipeline(t *testing.T) {
	pipelineID := uuid.New().String()

	s := &Snitch{
		pipelinesMtx: &sync.RWMutex{},
		pipelines:    make(map[string]map[string]*protos.Command),
		config: &Config{
			Logger: &loggerfakes.FakeLogger{},
		},
	}

	if err := s.detachPipeline(context.Background(), nil); err != ErrEmptyCommand {
		t.Error("Expected ErrEmptyCommand, got nil")
	}

	aud := &protos.Audience{
		OperationName: "test-operation",
		ServiceName:   "test-service",
		OperationType: protos.OperationType_OPERATION_TYPE_PRODUCER,
		ComponentName: "test-component",
	}

	cmd := &protos.Command{
		Audience: aud,
		Command: &protos.Command_DetachPipeline{
			DetachPipeline: &protos.DetachPipelineCommand{
				PipelineId: pipelineID,
			},
		},
	}

	s.pipelines[audToStr(aud)] = make(map[string]*protos.Command)
	s.pipelines[audToStr(aud)][pipelineID] = cmd

	err := s.detachPipeline(context.Background(), cmd)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if len(s.pipelines) > 0 {
		t.Errorf("Expected 0 audiences in pipeline map, got %d", len(s.pipelines))
	}

}

func TestPausePipeline(t *testing.T) {
	pipelineID := uuid.New().String()

	s := &Snitch{
		pipelinesPausedMtx: &sync.RWMutex{},
		pipelinesPaused:    make(map[string]map[string]*protos.Command),
		pipelinesMtx:       &sync.RWMutex{},
		pipelines:          make(map[string]map[string]*protos.Command),
		config: &Config{
			Logger: &loggerfakes.FakeLogger{},
		},
	}

	if err := s.pausePipeline(context.Background(), nil); err != ErrEmptyCommand {
		t.Error("Expected ErrEmptyCommand, got nil")
	}

	aud := &protos.Audience{
		OperationName: "test-operation",
		ServiceName:   "test-service",
		OperationType: protos.OperationType_OPERATION_TYPE_PRODUCER,
		ComponentName: "test-component",
	}

	attachCmd := &protos.Command{
		Audience: aud,
		Command: &protos.Command_AttachPipeline{
			AttachPipeline: &protos.AttachPipelineCommand{
				Pipeline: &protos.Pipeline{
					Id: pipelineID,
				},
			},
		},
	}

	if err := s.pausePipeline(context.Background(), attachCmd); err != ErrPipelineNotActive {

	}

	s.pipelines[audToStr(aud)] = make(map[string]*protos.Command)
	s.pipelines[audToStr(aud)][pipelineID] = attachCmd

	pauseCmd := &protos.Command{
		Audience: aud,
		Command: &protos.Command_PausePipeline{
			PausePipeline: &protos.PausePipelineCommand{
				PipelineId: pipelineID,
			},
		},
	}

	err := s.pausePipeline(context.Background(), pauseCmd)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if len(s.pipelinesPaused) != 1 {
		t.Errorf("Expected 1 audience in paused map, got %d", len(s.pipelinesPaused))
	}

	if len(s.pipelinesPaused[audToStr(aud)]) != 1 {
		t.Errorf("Expected 1 pipeline in paused map, got %d", len(s.pipelinesPaused[audToStr(aud)]))
	}

	// It should no longer be in active pipeline map
	if len(s.pipelines) > 0 {
		t.Errorf("Expected 0 audiences in active map, got %d", len(s.pipelines))
	}
}

func TestResumePipeline(t *testing.T) {
	pipelineID := uuid.New().String()

	s := &Snitch{
		pipelinesPausedMtx: &sync.RWMutex{},
		pipelinesPaused:    make(map[string]map[string]*protos.Command),
		pipelinesMtx:       &sync.RWMutex{},
		pipelines:          make(map[string]map[string]*protos.Command),
		config: &Config{
			Logger: &loggerfakes.FakeLogger{},
		},
	}

	if err := s.resumePipeline(context.Background(), nil); err != ErrEmptyCommand {
		t.Error("Expected ErrEmptyCommand, got nil")
	}

	aud := &protos.Audience{
		OperationName: "test-operation",
		ServiceName:   "test-service",
		OperationType: protos.OperationType_OPERATION_TYPE_PRODUCER,
		ComponentName: "test-component",
	}

	resumeCmd := &protos.Command{
		Audience: aud,
		Command: &protos.Command_ResumePipeline{
			ResumePipeline: &protos.ResumePipelineCommand{
				PipelineId: pipelineID,
			},
		},
	}

	if err := s.resumePipeline(context.Background(), resumeCmd); err != ErrPipelineNotPaused {
		t.Error("Expected ErrPipelineNotPaused, got nil")
	}

	attachCmd := &protos.Command{
		Audience: aud,
		Command: &protos.Command_AttachPipeline{
			AttachPipeline: &protos.AttachPipelineCommand{
				Pipeline: &protos.Pipeline{
					Id: pipelineID,
				},
			},
		},
	}

	s.pipelinesPaused[audToStr(aud)] = make(map[string]*protos.Command)
	s.pipelinesPaused[audToStr(aud)][pipelineID] = attachCmd

	err := s.resumePipeline(context.Background(), resumeCmd)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if len(s.pipelines) != 1 {
		t.Errorf("Expected 1 audience in active map, got %d", len(s.pipelines))
	}

	if len(s.pipelines[audToStr(aud)]) != 1 {
		t.Errorf("Expected 1 pipeline in active map, got %d", len(s.pipelines[audToStr(aud)]))
	}

	// It should no longer be in paused pipeline map
	if len(s.pipelinesPaused) > 0 {
		t.Errorf("Expected 0 audiences in paused map, got %d", len(s.pipelinesPaused))
	}
}
