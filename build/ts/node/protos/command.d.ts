import { MessageType } from "@protobuf-ts/runtime";
import { Pipeline } from "./pipeline.js";
import { Audience } from "./common.js";
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
export declare enum KVAction {
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
declare class Command$Type extends MessageType<Command> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.Command
 */
export declare const Command: Command$Type;
declare class AttachPipelineCommand$Type extends MessageType<AttachPipelineCommand> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.AttachPipelineCommand
 */
export declare const AttachPipelineCommand: AttachPipelineCommand$Type;
declare class DetachPipelineCommand$Type extends MessageType<DetachPipelineCommand> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.DetachPipelineCommand
 */
export declare const DetachPipelineCommand: DetachPipelineCommand$Type;
declare class PausePipelineCommand$Type extends MessageType<PausePipelineCommand> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.PausePipelineCommand
 */
export declare const PausePipelineCommand: PausePipelineCommand$Type;
declare class ResumePipelineCommand$Type extends MessageType<ResumePipelineCommand> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.ResumePipelineCommand
 */
export declare const ResumePipelineCommand: ResumePipelineCommand$Type;
declare class KeepAliveCommand$Type extends MessageType<KeepAliveCommand> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.KeepAliveCommand
 */
export declare const KeepAliveCommand: KeepAliveCommand$Type;
declare class KVObject$Type extends MessageType<KVObject> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.KVObject
 */
export declare const KVObject: KVObject$Type;
declare class KVInstruction$Type extends MessageType<KVInstruction> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.KVInstruction
 */
export declare const KVInstruction: KVInstruction$Type;
declare class KVCommand$Type extends MessageType<KVCommand> {
    constructor();
}
/**
 * @generated MessageType for protobuf message protos.KVCommand
 */
export declare const KVCommand: KVCommand$Type;
export {};
