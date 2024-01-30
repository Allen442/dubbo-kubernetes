// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: api/mesh/v1alpha1/mux.proto

package v1alpha1

import (
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Message_LegacyRequest
	//	*Message_LegacyResponse
	//	*Message_Request
	//	*Message_Response
	Value isMessage_Value `protobuf_oneof:"value"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_mesh_v1alpha1_mux_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_mesh_v1alpha1_mux_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_mesh_v1alpha1_mux_proto_rawDescGZIP(), []int{0}
}

func (m *Message) GetValue() isMessage_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Message) GetLegacyRequest() *v2.DiscoveryRequest {
	if x, ok := x.GetValue().(*Message_LegacyRequest); ok {
		return x.LegacyRequest
	}
	return nil
}

func (x *Message) GetLegacyResponse() *v2.DiscoveryResponse {
	if x, ok := x.GetValue().(*Message_LegacyResponse); ok {
		return x.LegacyResponse
	}
	return nil
}

func (x *Message) GetRequest() *v3.DiscoveryRequest {
	if x, ok := x.GetValue().(*Message_Request); ok {
		return x.Request
	}
	return nil
}

func (x *Message) GetResponse() *v3.DiscoveryResponse {
	if x, ok := x.GetValue().(*Message_Response); ok {
		return x.Response
	}
	return nil
}

type isMessage_Value interface {
	isMessage_Value()
}

type Message_LegacyRequest struct {
	LegacyRequest *v2.DiscoveryRequest `protobuf:"bytes,1,opt,name=legacy_request,json=legacyRequest,proto3,oneof"`
}

type Message_LegacyResponse struct {
	LegacyResponse *v2.DiscoveryResponse `protobuf:"bytes,2,opt,name=legacy_response,json=legacyResponse,proto3,oneof"`
}

type Message_Request struct {
	Request *v3.DiscoveryRequest `protobuf:"bytes,3,opt,name=request,proto3,oneof"`
}

type Message_Response struct {
	Response *v3.DiscoveryResponse `protobuf:"bytes,4,opt,name=response,proto3,oneof"`
}

func (*Message_LegacyRequest) isMessage_Value() {}

func (*Message_LegacyResponse) isMessage_Value() {}

func (*Message_Request) isMessage_Value() {}

func (*Message_Response) isMessage_Value() {}

var File_api_mesh_v1alpha1_mux_proto protoreflect.FileDescriptor

var file_api_mesh_v1alpha1_mux_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6d, 0x75, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64,
	0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x6c, 0x65, 0x67, 0x61,
	0x63, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0d, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4a, 0x0a, 0x0f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x6c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x63, 0x0a,
	0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x1c, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_mesh_v1alpha1_mux_proto_rawDescOnce sync.Once
	file_api_mesh_v1alpha1_mux_proto_rawDescData = file_api_mesh_v1alpha1_mux_proto_rawDesc
)

func file_api_mesh_v1alpha1_mux_proto_rawDescGZIP() []byte {
	file_api_mesh_v1alpha1_mux_proto_rawDescOnce.Do(func() {
		file_api_mesh_v1alpha1_mux_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_mesh_v1alpha1_mux_proto_rawDescData)
	})
	return file_api_mesh_v1alpha1_mux_proto_rawDescData
}

var file_api_mesh_v1alpha1_mux_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_mesh_v1alpha1_mux_proto_goTypes = []interface{}{
	(*Message)(nil),              // 0: dubbo.mesh.v1alpha1.Message
	(*v2.DiscoveryRequest)(nil),  // 1: envoy.api.v2.DiscoveryRequest
	(*v2.DiscoveryResponse)(nil), // 2: envoy.api.v2.DiscoveryResponse
	(*v3.DiscoveryRequest)(nil),  // 3: envoy.service.discovery.v3.DiscoveryRequest
	(*v3.DiscoveryResponse)(nil), // 4: envoy.service.discovery.v3.DiscoveryResponse
}
var file_api_mesh_v1alpha1_mux_proto_depIdxs = []int32{
	1, // 0: dubbo.mesh.v1alpha1.Message.legacy_request:type_name -> envoy.api.v2.DiscoveryRequest
	2, // 1: dubbo.mesh.v1alpha1.Message.legacy_response:type_name -> envoy.api.v2.DiscoveryResponse
	3, // 2: dubbo.mesh.v1alpha1.Message.request:type_name -> envoy.service.discovery.v3.DiscoveryRequest
	4, // 3: dubbo.mesh.v1alpha1.Message.response:type_name -> envoy.service.discovery.v3.DiscoveryResponse
	0, // 4: dubbo.mesh.v1alpha1.MultiplexService.StreamMessage:input_type -> dubbo.mesh.v1alpha1.Message
	0, // 5: dubbo.mesh.v1alpha1.MultiplexService.StreamMessage:output_type -> dubbo.mesh.v1alpha1.Message
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_mesh_v1alpha1_mux_proto_init() }
func file_api_mesh_v1alpha1_mux_proto_init() {
	if File_api_mesh_v1alpha1_mux_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_mesh_v1alpha1_mux_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_api_mesh_v1alpha1_mux_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_LegacyRequest)(nil),
		(*Message_LegacyResponse)(nil),
		(*Message_Request)(nil),
		(*Message_Response)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_mesh_v1alpha1_mux_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_mesh_v1alpha1_mux_proto_goTypes,
		DependencyIndexes: file_api_mesh_v1alpha1_mux_proto_depIdxs,
		MessageInfos:      file_api_mesh_v1alpha1_mux_proto_msgTypes,
	}.Build()
	File_api_mesh_v1alpha1_mux_proto = out.File
	file_api_mesh_v1alpha1_mux_proto_rawDesc = nil
	file_api_mesh_v1alpha1_mux_proto_goTypes = nil
	file_api_mesh_v1alpha1_mux_proto_depIdxs = nil
}
