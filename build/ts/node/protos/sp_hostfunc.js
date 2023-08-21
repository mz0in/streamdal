// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "sp_hostfunc.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf enum protos.HttpRequest.Method
 */
export var HttpRequest_Method;
(function (HttpRequest_Method) {
    /**
     * @generated from protobuf enum value: METHOD_UNSET = 0;
     */
    HttpRequest_Method[HttpRequest_Method["UNSET"] = 0] = "UNSET";
    /**
     * @generated from protobuf enum value: METHOD_GET = 1;
     */
    HttpRequest_Method[HttpRequest_Method["GET"] = 1] = "GET";
    /**
     * @generated from protobuf enum value: METHOD_POST = 2;
     */
    HttpRequest_Method[HttpRequest_Method["POST"] = 2] = "POST";
    /**
     * @generated from protobuf enum value: METHOD_PUT = 3;
     */
    HttpRequest_Method[HttpRequest_Method["PUT"] = 3] = "PUT";
    /**
     * @generated from protobuf enum value: METHOD_DELETE = 4;
     */
    HttpRequest_Method[HttpRequest_Method["DELETE"] = 4] = "DELETE";
})(HttpRequest_Method || (HttpRequest_Method = {}));
// @generated message type with reflection information, may provide speed optimized methods
class HttpRequest$Type extends MessageType {
    constructor() {
        super("protos.HttpRequest", [
            { no: 1, name: "method", kind: "enum", T: () => ["protos.HttpRequest.Method", HttpRequest_Method, "METHOD_"] },
            { no: 2, name: "url", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "body", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 4, name: "headers", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.HttpRequest
 */
export const HttpRequest = new HttpRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class HttpResponse$Type extends MessageType {
    constructor() {
        super("protos.HttpResponse", [
            { no: 1, name: "code", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "body", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 3, name: "headers", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.HttpResponse
 */
export const HttpResponse = new HttpResponse$Type();
//# sourceMappingURL=sp_hostfunc.js.map