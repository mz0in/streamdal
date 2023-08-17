// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "command.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { Pipeline } from "./pipeline.ts";
import { Audience } from "./common.ts";
/**
 * Command is used by snitch-server for sending commands to SDKs
 *
 * @generated from protobuf message protos.Command
 */
export interface Command {
    /**
     * Who is this command intended for?
     * NOTE: Some commands (such as KeepAliveCommand, KVCommand) do NOT use audience and will ignore it
     *
     * @generated from protobuf field: protos.Audience audience = 1;
     */
    audience?: Audience;
    /**
     * @generated from protobuf oneof: command
     */
    command: {
        oneofKind: "attachPipeline";
        /**
         * @generated from protobuf field: protos.AttachPipelineCommand attach_pipeline = 100;
         */
        attachPipeline: AttachPipelineCommand;
    } | {
        oneofKind: "detachPipeline";
        /**
         * @generated from protobuf field: protos.DetachPipelineCommand detach_pipeline = 101;
         */
        detachPipeline: DetachPipelineCommand;
    } | {
        oneofKind: "pausePipeline";
        /**
         * @generated from protobuf field: protos.PausePipelineCommand pause_pipeline = 102;
         */
        pausePipeline: PausePipelineCommand;
    } | {
        oneofKind: "resumePipeline";
        /**
         * @generated from protobuf field: protos.ResumePipelineCommand resume_pipeline = 103;
         */
        resumePipeline: ResumePipelineCommand;
    } | {
        oneofKind: "keepAlive";
        /**
         * @generated from protobuf field: protos.KeepAliveCommand keep_alive = 104;
         */
        keepAlive: KeepAliveCommand;
    } | {
        oneofKind: "kv";
        /**
         * snitch-server will emit this when a user makes changes to the KV store
         * via the KV HTTP API.
         *
         * @generated from protobuf field: protos.KVCommand kv = 105;
         */
        kv: KVCommand;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message protos.AttachPipelineCommand
 */
export interface AttachPipelineCommand {
    /**
     * @generated from protobuf field: protos.Pipeline pipeline = 1;
     */
    pipeline?: Pipeline;
}
/**
 * @generated from protobuf message protos.DetachPipelineCommand
 */
export interface DetachPipelineCommand {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
}
/**
 * @generated from protobuf message protos.PausePipelineCommand
 */
export interface PausePipelineCommand {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
}
/**
 * @generated from protobuf message protos.ResumePipelineCommand
 */
export interface ResumePipelineCommand {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
}
/**
 * Nothing needed in here, just a ping from server to SDK
 *
 * @generated from protobuf message protos.KeepAliveCommand
 */
export interface KeepAliveCommand {
}
/**
 * @generated from protobuf message protos.KVObject
 */
export interface KVObject {
    /**
     * Key regex: /^[a-zA-Z0-9_-:]+$/)
     *
     * @generated from protobuf field: string key = 1;
     */
    key: string;
    /**
     * KV value
     *
     * @generated from protobuf field: bytes value = 2;
     */
    value: Uint8Array;
    /**
     * When was this object created
     *
     * @generated from protobuf field: int64 created_at_unix_ts_nano_utc = 3;
     */
    createdAtUnixTsNanoUtc: bigint;
    /**
     * Last time the object was updated
     *
     * @generated from protobuf field: int64 updated_at_unix_ts_nano_utc = 4;
     */
    updatedAtUnixTsNanoUtc: bigint;
}
/**
 * Used in KVCommand to indicate a series of KV-related actions - ie. create, update, delete
 *
 * @generated from protobuf message protos.KVInstruction
 */
export interface KVInstruction {
    /**
     * Unique ID for this instruction
     *
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * What kind of an action is this?
     *
     * @generated from protobuf field: protos.KVAction action = 2;
     */
    action: KVAction;
    /**
     * KV object
     *
     * @generated from protobuf field: protos.KVObject object = 3;
     */
    object?: KVObject;
    /**
     * When this instruction was requested (usually will be the HTTP API request time)
     *
     * @generated from protobuf field: int64 requested_at_unix_ts_nano_utc = 4;
     */
    requestedAtUnixTsNanoUtc: bigint;
}
/**
 * @generated from protobuf message protos.KVCommand
 */
export interface KVCommand {
    /**
     * @generated from protobuf field: repeated protos.KVInstruction instructions = 1;
     */
    instructions: KVInstruction[];
}
/**
 * @generated from protobuf enum protos.KVAction
 */
export enum KVAction {
    /**
     * protolint:disable:this ENUM_FIELD_NAMES_PREFIX
     *
     * @generated from protobuf enum value: KV_ACTION_UNSET = 0;
     */
    KV_ACTION_UNSET = 0,
    /**
     * protolint:disable:this ENUM_FIELD_NAMES_PREFIX
     *
     * @generated from protobuf enum value: KV_ACTION_CREATE = 1;
     */
    KV_ACTION_CREATE = 1,
    /**
     * protolint:disable:this ENUM_FIELD_NAMES_PREFIX
     *
     * @generated from protobuf enum value: KV_ACTION_UPDATE = 2;
     */
    KV_ACTION_UPDATE = 2,
    /**
     * protolint:disable:this ENUM_FIELD_NAMES_PREFIX
     *
     * @generated from protobuf enum value: KV_ACTION_DELETE = 3;
     */
    KV_ACTION_DELETE = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class Command$Type extends MessageType<Command> {
    constructor() {
        super("protos.Command", [
            { no: 1, name: "audience", kind: "message", T: () => Audience },
            { no: 100, name: "attach_pipeline", kind: "message", oneof: "command", T: () => AttachPipelineCommand },
            { no: 101, name: "detach_pipeline", kind: "message", oneof: "command", T: () => DetachPipelineCommand },
            { no: 102, name: "pause_pipeline", kind: "message", oneof: "command", T: () => PausePipelineCommand },
            { no: 103, name: "resume_pipeline", kind: "message", oneof: "command", T: () => ResumePipelineCommand },
            { no: 104, name: "keep_alive", kind: "message", oneof: "command", T: () => KeepAliveCommand },
            { no: 105, name: "kv", kind: "message", oneof: "command", T: () => KVCommand }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.Command
 */
export const Command = new Command$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AttachPipelineCommand$Type extends MessageType<AttachPipelineCommand> {
    constructor() {
        super("protos.AttachPipelineCommand", [
            { no: 1, name: "pipeline", kind: "message", T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.AttachPipelineCommand
 */
export const AttachPipelineCommand = new AttachPipelineCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DetachPipelineCommand$Type extends MessageType<DetachPipelineCommand> {
    constructor() {
        super("protos.DetachPipelineCommand", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.DetachPipelineCommand
 */
export const DetachPipelineCommand = new DetachPipelineCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PausePipelineCommand$Type extends MessageType<PausePipelineCommand> {
    constructor() {
        super("protos.PausePipelineCommand", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.PausePipelineCommand
 */
export const PausePipelineCommand = new PausePipelineCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ResumePipelineCommand$Type extends MessageType<ResumePipelineCommand> {
    constructor() {
        super("protos.ResumePipelineCommand", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.ResumePipelineCommand
 */
export const ResumePipelineCommand = new ResumePipelineCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KeepAliveCommand$Type extends MessageType<KeepAliveCommand> {
    constructor() {
        super("protos.KeepAliveCommand", []);
    }
}
/**
 * @generated MessageType for protobuf message protos.KeepAliveCommand
 */
export const KeepAliveCommand = new KeepAliveCommand$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVObject$Type extends MessageType<KVObject> {
    constructor() {
        super("protos.KVObject", [
            { no: 1, name: "key", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "value", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 3, name: "created_at_unix_ts_nano_utc", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "updated_at_unix_ts_nano_utc", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.KVObject
 */
export const KVObject = new KVObject$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVInstruction$Type extends MessageType<KVInstruction> {
    constructor() {
        super("protos.KVInstruction", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "action", kind: "enum", T: () => ["protos.KVAction", KVAction] },
            { no: 3, name: "object", kind: "message", T: () => KVObject },
            { no: 4, name: "requested_at_unix_ts_nano_utc", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.KVInstruction
 */
export const KVInstruction = new KVInstruction$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KVCommand$Type extends MessageType<KVCommand> {
    constructor() {
        super("protos.KVCommand", [
            { no: 1, name: "instructions", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KVInstruction }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.KVCommand
 */
export const KVCommand = new KVCommand$Type();
