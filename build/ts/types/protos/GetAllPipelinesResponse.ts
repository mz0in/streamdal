// Original file: ../../protos/pipeline.proto

import type { SetPipelineRequest as _protos_SetPipelineRequest, SetPipelineRequest__Output as _protos_SetPipelineRequest__Output } from '../protos/SetPipelineRequest';
import type { ResponseStatus as _protos_ResponseStatus, ResponseStatus__Output as _protos_ResponseStatus__Output } from '../protos/ResponseStatus';

export interface GetAllPipelinesResponse {
  'pipelines'?: (_protos_SetPipelineRequest)[];
  'status'?: (_protos_ResponseStatus);
  'message'?: (string);
}

export interface GetAllPipelinesResponse__Output {
  'pipelines': (_protos_SetPipelineRequest__Output)[];
  'status': (_protos_ResponseStatus__Output);
  'message': (string);
}
