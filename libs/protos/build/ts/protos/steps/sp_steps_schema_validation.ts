// @generated by protobuf-ts 2.9.0 with parameter long_type_string
// @generated from protobuf file "steps/sp_steps_schema_validation.proto" (package "protos.steps", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message protos.steps.SchemaValidationStep
 */
export interface SchemaValidationStep {
    /**
     * @generated from protobuf field: protos.steps.SchemaValidationType type = 1;
     */
    type: SchemaValidationType;
    /**
     * @generated from protobuf field: protos.steps.SchemaValidationCondition condition = 2;
     */
    condition: SchemaValidationCondition;
    /**
     * @generated from protobuf oneof: options
     */
    options: {
        oneofKind: "jsonSchema";
        /**
         * @generated from protobuf field: protos.steps.SchemaValidationJSONSchema json_schema = 101;
         */
        jsonSchema: SchemaValidationJSONSchema;
    } | {
        oneofKind: undefined;
    };
}
/**
 * @generated from protobuf message protos.steps.SchemaValidationJSONSchema
 */
export interface SchemaValidationJSONSchema {
    /**
     * @generated from protobuf field: bytes json_schema = 1;
     */
    jsonSchema: Uint8Array;
}
/**
 * TODO: expand for protobuf, avro, etc.
 *
 * @generated from protobuf enum protos.steps.SchemaValidationType
 */
export enum SchemaValidationType {
    /**
     * @generated from protobuf enum value: SCHEMA_VALIDATION_TYPE_UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: SCHEMA_VALIDATION_TYPE_JSONSCHEMA = 1;
     */
    JSONSCHEMA = 1
}
/**
 * @generated from protobuf enum protos.steps.SchemaValidationCondition
 */
export enum SchemaValidationCondition {
    /**
     * @generated from protobuf enum value: SCHEMA_VALIDATION_CONDITION_UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: SCHEMA_VALIDATION_CONDITION_MATCH = 1;
     */
    MATCH = 1,
    /**
     * TODO: backwards compat, evolve, etc.
     *
     * @generated from protobuf enum value: SCHEMA_VALIDATION_CONDITION_NOT_MATCH = 2;
     */
    NOT_MATCH = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class SchemaValidationStep$Type extends MessageType<SchemaValidationStep> {
    constructor() {
        super("protos.steps.SchemaValidationStep", [
            { no: 1, name: "type", kind: "enum", T: () => ["protos.steps.SchemaValidationType", SchemaValidationType, "SCHEMA_VALIDATION_TYPE_"] },
            { no: 2, name: "condition", kind: "enum", T: () => ["protos.steps.SchemaValidationCondition", SchemaValidationCondition, "SCHEMA_VALIDATION_CONDITION_"] },
            { no: 101, name: "json_schema", kind: "message", oneof: "options", T: () => SchemaValidationJSONSchema }
        ]);
    }
    create(value?: PartialMessage<SchemaValidationStep>): SchemaValidationStep {
        const message = { type: 0, condition: 0, options: { oneofKind: undefined } };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SchemaValidationStep>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SchemaValidationStep): SchemaValidationStep {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* protos.steps.SchemaValidationType type */ 1:
                    message.type = reader.int32();
                    break;
                case /* protos.steps.SchemaValidationCondition condition */ 2:
                    message.condition = reader.int32();
                    break;
                case /* protos.steps.SchemaValidationJSONSchema json_schema */ 101:
                    message.options = {
                        oneofKind: "jsonSchema",
                        jsonSchema: SchemaValidationJSONSchema.internalBinaryRead(reader, reader.uint32(), options, (message.options as any).jsonSchema)
                    };
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SchemaValidationStep, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* protos.steps.SchemaValidationType type = 1; */
        if (message.type !== 0)
            writer.tag(1, WireType.Varint).int32(message.type);
        /* protos.steps.SchemaValidationCondition condition = 2; */
        if (message.condition !== 0)
            writer.tag(2, WireType.Varint).int32(message.condition);
        /* protos.steps.SchemaValidationJSONSchema json_schema = 101; */
        if (message.options.oneofKind === "jsonSchema")
            SchemaValidationJSONSchema.internalBinaryWrite(message.options.jsonSchema, writer.tag(101, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.SchemaValidationStep
 */
export const SchemaValidationStep = new SchemaValidationStep$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SchemaValidationJSONSchema$Type extends MessageType<SchemaValidationJSONSchema> {
    constructor() {
        super("protos.steps.SchemaValidationJSONSchema", [
            { no: 1, name: "json_schema", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<SchemaValidationJSONSchema>): SchemaValidationJSONSchema {
        const message = { jsonSchema: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SchemaValidationJSONSchema>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SchemaValidationJSONSchema): SchemaValidationJSONSchema {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes json_schema */ 1:
                    message.jsonSchema = reader.bytes();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SchemaValidationJSONSchema, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes json_schema = 1; */
        if (message.jsonSchema.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.jsonSchema);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.SchemaValidationJSONSchema
 */
export const SchemaValidationJSONSchema = new SchemaValidationJSONSchema$Type();
