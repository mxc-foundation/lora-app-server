// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: ns-extra.proto

package nsextra

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SendDelayedProprietaryPayloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MACPayload of the proprietary LoRaWAN frame.
	MacPayload []byte `protobuf:"bytes,1,opt,name=mac_payload,json=macPayload,proto3" json:"mac_payload,omitempty"`
	// MIC of the proprietary LoRaWAN frame (must be 4 bytes).
	Mic []byte `protobuf:"bytes,2,opt,name=mic,proto3" json:"mic,omitempty"`
	// Gateway MAC address(es) to use for transmitting the LoRaWAN frame.
	GatewayMacs [][]byte `protobuf:"bytes,3,rep,name=gateway_macs,json=gatewayMacs,proto3" json:"gateway_macs,omitempty"`
	// Set to true for sending as a gateway, or false for sending as a node.
	// In the latter case the frame will be received by other gateways.
	PolarizationInversion bool `protobuf:"varint,4,opt,name=polarization_inversion,json=polarizationInversion,proto3" json:"polarization_inversion,omitempty"`
	// Frequency (Hz) to use for the transmission.
	Frequency uint32 `protobuf:"varint,5,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Data-rate to use for the transmission.
	Dr uint32 `protobuf:"varint,6,opt,name=dr,proto3" json:"dr,omitempty"`
	// Gateway specific context.
	Context []byte `protobuf:"bytes,7,opt,name=context,proto3" json:"context,omitempty"`
	// The delay will be added to the gateway internal timing, provided by the context object.
	Delay *duration.Duration `protobuf:"bytes,8,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *SendDelayedProprietaryPayloadRequest) Reset() {
	*x = SendDelayedProprietaryPayloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ns_extra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDelayedProprietaryPayloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDelayedProprietaryPayloadRequest) ProtoMessage() {}

func (x *SendDelayedProprietaryPayloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ns_extra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDelayedProprietaryPayloadRequest.ProtoReflect.Descriptor instead.
func (*SendDelayedProprietaryPayloadRequest) Descriptor() ([]byte, []int) {
	return file_ns_extra_proto_rawDescGZIP(), []int{0}
}

func (x *SendDelayedProprietaryPayloadRequest) GetMacPayload() []byte {
	if x != nil {
		return x.MacPayload
	}
	return nil
}

func (x *SendDelayedProprietaryPayloadRequest) GetMic() []byte {
	if x != nil {
		return x.Mic
	}
	return nil
}

func (x *SendDelayedProprietaryPayloadRequest) GetGatewayMacs() [][]byte {
	if x != nil {
		return x.GatewayMacs
	}
	return nil
}

func (x *SendDelayedProprietaryPayloadRequest) GetPolarizationInversion() bool {
	if x != nil {
		return x.PolarizationInversion
	}
	return false
}

func (x *SendDelayedProprietaryPayloadRequest) GetFrequency() uint32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *SendDelayedProprietaryPayloadRequest) GetDr() uint32 {
	if x != nil {
		return x.Dr
	}
	return 0
}

func (x *SendDelayedProprietaryPayloadRequest) GetContext() []byte {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *SendDelayedProprietaryPayloadRequest) GetDelay() *duration.Duration {
	if x != nil {
		return x.Delay
	}
	return nil
}

var File_ns_extra_proto protoreflect.FileDescriptor

var file_ns_extra_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x73, 0x2d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6e, 0x73, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x24, 0x53, 0x65, 0x6e, 0x64, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x72, 0x69, 0x65, 0x74, 0x61, 0x72,
	0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d,
	0x69, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x6d, 0x61,
	0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x4d, 0x61, 0x63, 0x73, 0x12, 0x35, 0x0a, 0x16, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x32, 0x85, 0x01, 0x0a, 0x19, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x72, 0x69, 0x65, 0x74, 0x61, 0x72, 0x79, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2d, 0x2e, 0x6e, 0x73, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x72, 0x69,
	0x65, 0x74, 0x61, 0x72, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x41, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x63, 0x2d,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x70, 0x77, 0x61, 0x6e,
	0x2d, 0x61, 0x70, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6e, 0x73, 0x2d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x3b, 0x6e, 0x73, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ns_extra_proto_rawDescOnce sync.Once
	file_ns_extra_proto_rawDescData = file_ns_extra_proto_rawDesc
)

func file_ns_extra_proto_rawDescGZIP() []byte {
	file_ns_extra_proto_rawDescOnce.Do(func() {
		file_ns_extra_proto_rawDescData = protoimpl.X.CompressGZIP(file_ns_extra_proto_rawDescData)
	})
	return file_ns_extra_proto_rawDescData
}

var file_ns_extra_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ns_extra_proto_goTypes = []interface{}{
	(*SendDelayedProprietaryPayloadRequest)(nil), // 0: nsextra.SendDelayedProprietaryPayloadRequest
	(*duration.Duration)(nil),                    // 1: google.protobuf.Duration
	(*empty.Empty)(nil),                          // 2: google.protobuf.Empty
}
var file_ns_extra_proto_depIdxs = []int32{
	1, // 0: nsextra.SendDelayedProprietaryPayloadRequest.delay:type_name -> google.protobuf.Duration
	0, // 1: nsextra.NetworkServerExtraService.SendDelayedProprietaryPayload:input_type -> nsextra.SendDelayedProprietaryPayloadRequest
	2, // 2: nsextra.NetworkServerExtraService.SendDelayedProprietaryPayload:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ns_extra_proto_init() }
func file_ns_extra_proto_init() {
	if File_ns_extra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ns_extra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDelayedProprietaryPayloadRequest); i {
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
			RawDescriptor: file_ns_extra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ns_extra_proto_goTypes,
		DependencyIndexes: file_ns_extra_proto_depIdxs,
		MessageInfos:      file_ns_extra_proto_msgTypes,
	}.Build()
	File_ns_extra_proto = out.File
	file_ns_extra_proto_rawDesc = nil
	file_ns_extra_proto_goTypes = nil
	file_ns_extra_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NetworkServerExtraServiceClient is the client API for NetworkServerExtraService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServerExtraServiceClient interface {
	// SendDelayedProprietaryPayload send a delayed payload using the 'Proprietary' LoRaWAN message-type.
	SendDelayedProprietaryPayload(ctx context.Context, in *SendDelayedProprietaryPayloadRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServerExtraServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkServerExtraServiceClient(cc grpc.ClientConnInterface) NetworkServerExtraServiceClient {
	return &networkServerExtraServiceClient{cc}
}

func (c *networkServerExtraServiceClient) SendDelayedProprietaryPayload(ctx context.Context, in *SendDelayedProprietaryPayloadRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nsextra.NetworkServerExtraService/SendDelayedProprietaryPayload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServerExtraServiceServer is the server API for NetworkServerExtraService service.
type NetworkServerExtraServiceServer interface {
	// SendDelayedProprietaryPayload send a delayed payload using the 'Proprietary' LoRaWAN message-type.
	SendDelayedProprietaryPayload(context.Context, *SendDelayedProprietaryPayloadRequest) (*empty.Empty, error)
}

// UnimplementedNetworkServerExtraServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServerExtraServiceServer struct {
}

func (*UnimplementedNetworkServerExtraServiceServer) SendDelayedProprietaryPayload(context.Context, *SendDelayedProprietaryPayloadRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDelayedProprietaryPayload not implemented")
}

func RegisterNetworkServerExtraServiceServer(s *grpc.Server, srv NetworkServerExtraServiceServer) {
	s.RegisterService(&_NetworkServerExtraService_serviceDesc, srv)
}

func _NetworkServerExtraService_SendDelayedProprietaryPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDelayedProprietaryPayloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerExtraServiceServer).SendDelayedProprietaryPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsextra.NetworkServerExtraService/SendDelayedProprietaryPayload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerExtraServiceServer).SendDelayedProprietaryPayload(ctx, req.(*SendDelayedProprietaryPayloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServerExtraService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nsextra.NetworkServerExtraService",
	HandlerType: (*NetworkServerExtraServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendDelayedProprietaryPayload",
			Handler:    _NetworkServerExtraService_SendDelayedProprietaryPayload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ns-extra.proto",
}
