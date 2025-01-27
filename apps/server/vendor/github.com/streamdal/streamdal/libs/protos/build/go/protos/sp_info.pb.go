// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sp_info.proto

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

type ClientType int32

const (
	ClientType_CLIENT_TYPE_UNSET ClientType = 0
	ClientType_CLIENT_TYPE_SDK   ClientType = 1
	ClientType_CLIENT_TYPE_SHIM  ClientType = 2
)

// Enum value maps for ClientType.
var (
	ClientType_name = map[int32]string{
		0: "CLIENT_TYPE_UNSET",
		1: "CLIENT_TYPE_SDK",
		2: "CLIENT_TYPE_SHIM",
	}
	ClientType_value = map[string]int32{
		"CLIENT_TYPE_UNSET": 0,
		"CLIENT_TYPE_SDK":   1,
		"CLIENT_TYPE_SHIM":  2,
	}
)

func (x ClientType) Enum() *ClientType {
	p := new(ClientType)
	*p = x
	return p
}

func (x ClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_sp_info_proto_enumTypes[0].Descriptor()
}

func (ClientType) Type() protoreflect.EnumType {
	return &file_sp_info_proto_enumTypes[0]
}

func (x ClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientType.Descriptor instead.
func (ClientType) EnumDescriptor() ([]byte, []int) {
	return file_sp_info_proto_rawDescGZIP(), []int{0}
}

type LiveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If empty, client has not announced any audiences
	Audiences []*Audience `protobuf:"bytes,1,rep,name=audiences,proto3" json:"audiences,omitempty"`
	Client    *ClientInfo `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *LiveInfo) Reset() {
	*x = LiveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveInfo) ProtoMessage() {}

func (x *LiveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sp_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveInfo.ProtoReflect.Descriptor instead.
func (*LiveInfo) Descriptor() ([]byte, []int) {
	return file_sp_info_proto_rawDescGZIP(), []int{0}
}

func (x *LiveInfo) GetAudiences() []*Audience {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *LiveInfo) GetClient() *ClientInfo {
	if x != nil {
		return x.Client
	}
	return nil
}

type PipelineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// What audience(s) this pipeline is attached to (none if empty)
	Audiences []*Audience `protobuf:"bytes,1,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// Pipeline config
	Pipeline *Pipeline `protobuf:"bytes,2,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
}

func (x *PipelineInfo) Reset() {
	*x = PipelineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineInfo) ProtoMessage() {}

func (x *PipelineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sp_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineInfo.ProtoReflect.Descriptor instead.
func (*PipelineInfo) Descriptor() ([]byte, []int) {
	return file_sp_info_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineInfo) GetAudiences() []*Audience {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *PipelineInfo) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

// Most of this is constructed by client SDKs and provided during Register call
type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientType     ClientType `protobuf:"varint,1,opt,name=client_type,json=clientType,proto3,enum=protos.ClientType" json:"client_type,omitempty"`
	LibraryName    string     `protobuf:"bytes,2,opt,name=library_name,json=libraryName,proto3" json:"library_name,omitempty"`
	LibraryVersion string     `protobuf:"bytes,3,opt,name=library_version,json=libraryVersion,proto3" json:"library_version,omitempty"`
	Language       string     `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	Arch           string     `protobuf:"bytes,5,opt,name=arch,proto3" json:"arch,omitempty"`
	Os             string     `protobuf:"bytes,6,opt,name=os,proto3" json:"os,omitempty"`
	// Filled out by server on GetAll()
	XSessionId   *string `protobuf:"bytes,7,opt,name=_session_id,json=SessionId,proto3,oneof" json:"_session_id,omitempty"`       // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
	XServiceName *string `protobuf:"bytes,8,opt,name=_service_name,json=ServiceName,proto3,oneof" json:"_service_name,omitempty"` // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
	XNodeName    *string `protobuf:"bytes,9,opt,name=_node_name,json=NodeName,proto3,oneof" json:"_node_name,omitempty"`          // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sp_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_sp_info_proto_rawDescGZIP(), []int{2}
}

func (x *ClientInfo) GetClientType() ClientType {
	if x != nil {
		return x.ClientType
	}
	return ClientType_CLIENT_TYPE_UNSET
}

func (x *ClientInfo) GetLibraryName() string {
	if x != nil {
		return x.LibraryName
	}
	return ""
}

func (x *ClientInfo) GetLibraryVersion() string {
	if x != nil {
		return x.LibraryVersion
	}
	return ""
}

func (x *ClientInfo) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *ClientInfo) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *ClientInfo) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *ClientInfo) GetXSessionId() string {
	if x != nil && x.XSessionId != nil {
		return *x.XSessionId
	}
	return ""
}

func (x *ClientInfo) GetXServiceName() string {
	if x != nil && x.XServiceName != nil {
		return *x.XServiceName
	}
	return ""
}

func (x *ClientInfo) GetXNodeName() string {
	if x != nil && x.XNodeName != nil {
		return *x.XNodeName
	}
	return ""
}

var File_sp_info_proto protoreflect.FileDescriptor

var file_sp_info_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0f, 0x73, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x70, 0x5f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x08, 0x4c,
	0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x61, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x22, 0x6c, 0x0a, 0x0c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x22, 0xef, 0x02, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f,
	0x73, 0x12, 0x23, 0x0a, 0x0b, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x0a, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x58, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x58, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x58, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x2a, 0x4e, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x49, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x44, 0x4b, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x49,
	0x4d, 0x10, 0x02, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_info_proto_rawDescOnce sync.Once
	file_sp_info_proto_rawDescData = file_sp_info_proto_rawDesc
)

func file_sp_info_proto_rawDescGZIP() []byte {
	file_sp_info_proto_rawDescOnce.Do(func() {
		file_sp_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_info_proto_rawDescData)
	})
	return file_sp_info_proto_rawDescData
}

var file_sp_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sp_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sp_info_proto_goTypes = []interface{}{
	(ClientType)(0),      // 0: protos.ClientType
	(*LiveInfo)(nil),     // 1: protos.LiveInfo
	(*PipelineInfo)(nil), // 2: protos.PipelineInfo
	(*ClientInfo)(nil),   // 3: protos.ClientInfo
	(*Audience)(nil),     // 4: protos.Audience
	(*Pipeline)(nil),     // 5: protos.Pipeline
}
var file_sp_info_proto_depIdxs = []int32{
	4, // 0: protos.LiveInfo.audiences:type_name -> protos.Audience
	3, // 1: protos.LiveInfo.client:type_name -> protos.ClientInfo
	4, // 2: protos.PipelineInfo.audiences:type_name -> protos.Audience
	5, // 3: protos.PipelineInfo.pipeline:type_name -> protos.Pipeline
	0, // 4: protos.ClientInfo.client_type:type_name -> protos.ClientType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sp_info_proto_init() }
func file_sp_info_proto_init() {
	if File_sp_info_proto != nil {
		return
	}
	file_sp_common_proto_init()
	file_sp_pipeline_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sp_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveInfo); i {
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
		file_sp_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineInfo); i {
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
		file_sp_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
	file_sp_info_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sp_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sp_info_proto_goTypes,
		DependencyIndexes: file_sp_info_proto_depIdxs,
		EnumInfos:         file_sp_info_proto_enumTypes,
		MessageInfos:      file_sp_info_proto_msgTypes,
	}.Build()
	File_sp_info_proto = out.File
	file_sp_info_proto_rawDesc = nil
	file_sp_info_proto_goTypes = nil
	file_sp_info_proto_depIdxs = nil
}
