// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: heartbeat.proto

package appserver_serves_gateway

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MiningRequest sends gateway list to m2m
type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayMac string `protobuf:"bytes,1,opt,name=gateway_mac,json=gatewayMac,proto3" json:"gateway_mac,omitempty"`
	Model      string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	ConfigHash string `protobuf:"bytes,3,opt,name=config_hash,json=configHash,proto3" json:"config_hash,omitempty"`
	OsVersion  string `protobuf:"bytes,4,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	Statistics string `protobuf:"bytes,5,opt,name=statistics,proto3" json:"statistics,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_heartbeat_proto_rawDescGZIP(), []int{0}
}

func (x *HeartbeatRequest) GetGatewayMac() string {
	if x != nil {
		return x.GatewayMac
	}
	return ""
}

func (x *HeartbeatRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *HeartbeatRequest) GetConfigHash() string {
	if x != nil {
		return x.ConfigHash
	}
	return ""
}

func (x *HeartbeatRequest) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *HeartbeatRequest) GetStatistics() string {
	if x != nil {
		return x.Statistics
	}
	return ""
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewFirmwareLink string `protobuf:"bytes,1,opt,name=new_firmware_link,json=newFirmwareLink,proto3" json:"new_firmware_link,omitempty"`
	Config          string `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_heartbeat_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatResponse) GetNewFirmwareLink() string {
	if x != nil {
		return x.NewFirmwareLink
	}
	return ""
}

func (x *HeartbeatResponse) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

var File_heartbeat_proto protoreflect.FileDescriptor

var file_heartbeat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x73, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0xa9, 0x01, 0x0a, 0x10,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x6d, 0x61, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x61,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x73, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x57, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x6e, 0x65, 0x77, 0x5f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x77, 0x46, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x32, 0x78, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x73, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x62, 0x5a, 0x60, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x63, 0x2d, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x70, 0x77, 0x61, 0x6e, 0x2d, 0x61, 0x70,
	0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x2d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x3b, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heartbeat_proto_rawDescOnce sync.Once
	file_heartbeat_proto_rawDescData = file_heartbeat_proto_rawDesc
)

func file_heartbeat_proto_rawDescGZIP() []byte {
	file_heartbeat_proto_rawDescOnce.Do(func() {
		file_heartbeat_proto_rawDescData = protoimpl.X.CompressGZIP(file_heartbeat_proto_rawDescData)
	})
	return file_heartbeat_proto_rawDescData
}

var file_heartbeat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_heartbeat_proto_goTypes = []interface{}{
	(*HeartbeatRequest)(nil),  // 0: appserver_serves_gateway.HeartbeatRequest
	(*HeartbeatResponse)(nil), // 1: appserver_serves_gateway.HeartbeatResponse
}
var file_heartbeat_proto_depIdxs = []int32{
	0, // 0: appserver_serves_gateway.HeartbeatService.Heartbeat:input_type -> appserver_serves_gateway.HeartbeatRequest
	1, // 1: appserver_serves_gateway.HeartbeatService.Heartbeat:output_type -> appserver_serves_gateway.HeartbeatResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heartbeat_proto_init() }
func file_heartbeat_proto_init() {
	if File_heartbeat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heartbeat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_heartbeat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
			RawDescriptor: file_heartbeat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heartbeat_proto_goTypes,
		DependencyIndexes: file_heartbeat_proto_depIdxs,
		MessageInfos:      file_heartbeat_proto_msgTypes,
	}.Build()
	File_heartbeat_proto = out.File
	file_heartbeat_proto_rawDesc = nil
	file_heartbeat_proto_goTypes = nil
	file_heartbeat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HeartbeatServiceClient is the client API for HeartbeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatServiceClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type heartbeatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartbeatServiceClient(cc grpc.ClientConnInterface) HeartbeatServiceClient {
	return &heartbeatServiceClient{cc}
}

func (c *heartbeatServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_gateway.HeartbeatService/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServiceServer is the server API for HeartbeatService service.
type HeartbeatServiceServer interface {
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
}

// UnimplementedHeartbeatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHeartbeatServiceServer struct {
}

func (*UnimplementedHeartbeatServiceServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterHeartbeatServiceServer(s *grpc.Server, srv HeartbeatServiceServer) {
	s.RegisterService(&_HeartbeatService_serviceDesc, srv)
}

func _HeartbeatService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_gateway.HeartbeatService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeartbeatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appserver_serves_gateway.HeartbeatService",
	HandlerType: (*HeartbeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _HeartbeatService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}
