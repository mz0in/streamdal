// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: sp_hostfunc.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HttpRequest_Method int32

const (
	HttpRequest_METHOD_UNSET  HttpRequest_Method = 0
	HttpRequest_METHOD_GET    HttpRequest_Method = 1
	HttpRequest_METHOD_POST   HttpRequest_Method = 2
	HttpRequest_METHOD_PUT    HttpRequest_Method = 3
	HttpRequest_METHOD_DELETE HttpRequest_Method = 4
)

// Enum value maps for HttpRequest_Method.
var (
	HttpRequest_Method_name = map[int32]string{
		0: "METHOD_UNSET",
		1: "METHOD_GET",
		2: "METHOD_POST",
		3: "METHOD_PUT",
		4: "METHOD_DELETE",
	}
	HttpRequest_Method_value = map[string]int32{
		"METHOD_UNSET":  0,
		"METHOD_GET":    1,
		"METHOD_POST":   2,
		"METHOD_PUT":    3,
		"METHOD_DELETE": 4,
	}
)

func (x HttpRequest_Method) Enum() *HttpRequest_Method {
	p := new(HttpRequest_Method)
	*p = x
	return p
}

func (x HttpRequest_Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpRequest_Method) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_hostfunc_proto_enumTypes[0].Descriptor()
}

func (HttpRequest_Method) Type() protoreflect.EnumType {
	return &file_sp_hostfunc_proto_enumTypes[0]
}

func (x HttpRequest_Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpRequest_Method.Descriptor instead.
func (HttpRequest_Method) EnumDescriptor() ([]byte, []int) {
	return file_sp_hostfunc_proto_rawDescGZIP(), []int{0, 0}
}

type HttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  HttpRequest_Method `protobuf:"varint,1,opt,name=method,proto3,enum=protos.HttpRequest_Method" json:"method,omitempty"`
	Url     string             `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Body    []byte             `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Headers map[string]string  `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HttpRequest) Reset() {
	*x = HttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_hostfunc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequest) ProtoMessage() {}

func (x *HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_hostfunc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequest.ProtoReflect.Descriptor instead.
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return file_sp_hostfunc_proto_rawDescGZIP(), []int{0}
}

func (x *HttpRequest) GetMethod() HttpRequest_Method {
	if x != nil {
		return x.Method
	}
	return HttpRequest_METHOD_UNSET
}

func (x *HttpRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HttpRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *HttpRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type HttpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Body    []byte            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HttpResponse) Reset() {
	*x = HttpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_hostfunc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpResponse) ProtoMessage() {}

func (x *HttpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sp_hostfunc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpResponse.ProtoReflect.Descriptor instead.
func (*HttpResponse) Descriptor() ([]byte, []int) {
	return file_sp_hostfunc_proto_rawDescGZIP(), []int{1}
}

func (x *HttpResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HttpResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *HttpResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_sp_hostfunc_proto protoreflect.FileDescriptor

var file_sp_hostfunc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x70, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x66, 0x75, 0x6e, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xbf, 0x02, 0x0a, 0x0b,
	0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a,
	0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x50, 0x55, 0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x22, 0xaf, 0x01,
	0x0a, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73, 0x6e, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_hostfunc_proto_rawDescOnce sync.Once
	file_sp_hostfunc_proto_rawDescData = file_sp_hostfunc_proto_rawDesc
)

func file_sp_hostfunc_proto_rawDescGZIP() []byte {
	file_sp_hostfunc_proto_rawDescOnce.Do(func() {
		file_sp_hostfunc_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_hostfunc_proto_rawDescData)
	})
	return file_sp_hostfunc_proto_rawDescData
}

var file_sp_hostfunc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sp_hostfunc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sp_hostfunc_proto_goTypes = []interface{}{
	(HttpRequest_Method)(0), // 0: protos.HttpRequest.Method
	(*HttpRequest)(nil),     // 1: protos.HttpRequest
	(*HttpResponse)(nil),    // 2: protos.HttpResponse
	nil,                     // 3: protos.HttpRequest.HeadersEntry
	nil,                     // 4: protos.HttpResponse.HeadersEntry
}
var file_sp_hostfunc_proto_depIdxs = []int32{
	0, // 0: protos.HttpRequest.method:type_name -> protos.HttpRequest.Method
	3, // 1: protos.HttpRequest.headers:type_name -> protos.HttpRequest.HeadersEntry
	4, // 2: protos.HttpResponse.headers:type_name -> protos.HttpResponse.HeadersEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sp_hostfunc_proto_init() }
func file_sp_hostfunc_proto_init() {
	if File_sp_hostfunc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sp_hostfunc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sp_hostfunc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_hostfunc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_hostfunc_proto_goTypes,
		DependencyIndexes: file_sp_hostfunc_proto_depIdxs,
		EnumInfos:         file_sp_hostfunc_proto_enumTypes,
		MessageInfos:      file_sp_hostfunc_proto_msgTypes,
	}.Build()
	File_sp_hostfunc_proto = out.File
	file_sp_hostfunc_proto_rawDesc = nil
	file_sp_hostfunc_proto_goTypes = nil
	file_sp_hostfunc_proto_depIdxs = nil
}
