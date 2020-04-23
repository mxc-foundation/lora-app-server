// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverInfo.proto

package appserver_serves_ui

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ServerRegion int32

const (
	ServerRegion_NOT_DEFINED ServerRegion = 0
	ServerRegion_AVERAGE     ServerRegion = 1
	ServerRegion_RESTRICTED  ServerRegion = 2
)

var ServerRegion_name = map[int32]string{
	0: "NOT_DEFINED",
	1: "AVERAGE",
	2: "RESTRICTED",
}

var ServerRegion_value = map[string]int32{
	"NOT_DEFINED": 0,
	"AVERAGE":     1,
	"RESTRICTED":  2,
}

func (x ServerRegion) String() string {
	return proto.EnumName(ServerRegion_name, int32(x))
}

func (ServerRegion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_386dce779fe74cf3, []int{0}
}

type GetAppserverVersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppserverVersionResponse) Reset()         { *m = GetAppserverVersionResponse{} }
func (m *GetAppserverVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetAppserverVersionResponse) ProtoMessage()    {}
func (*GetAppserverVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_386dce779fe74cf3, []int{0}
}

func (m *GetAppserverVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppserverVersionResponse.Unmarshal(m, b)
}
func (m *GetAppserverVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppserverVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetAppserverVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppserverVersionResponse.Merge(m, src)
}
func (m *GetAppserverVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetAppserverVersionResponse.Size(m)
}
func (m *GetAppserverVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppserverVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppserverVersionResponse proto.InternalMessageInfo

func (m *GetAppserverVersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetServerRegionResponse struct {
	ServerRegion         string   `protobuf:"bytes,1,opt,name=server_region,json=serverRegion,proto3" json:"server_region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServerRegionResponse) Reset()         { *m = GetServerRegionResponse{} }
func (m *GetServerRegionResponse) String() string { return proto.CompactTextString(m) }
func (*GetServerRegionResponse) ProtoMessage()    {}
func (*GetServerRegionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_386dce779fe74cf3, []int{1}
}

func (m *GetServerRegionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServerRegionResponse.Unmarshal(m, b)
}
func (m *GetServerRegionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServerRegionResponse.Marshal(b, m, deterministic)
}
func (m *GetServerRegionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServerRegionResponse.Merge(m, src)
}
func (m *GetServerRegionResponse) XXX_Size() int {
	return xxx_messageInfo_GetServerRegionResponse.Size(m)
}
func (m *GetServerRegionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServerRegionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServerRegionResponse proto.InternalMessageInfo

func (m *GetServerRegionResponse) GetServerRegion() string {
	if m != nil {
		return m.ServerRegion
	}
	return ""
}

func init() {
	proto.RegisterEnum("appserver_serves_ui.ServerRegion", ServerRegion_name, ServerRegion_value)
	proto.RegisterType((*GetAppserverVersionResponse)(nil), "appserver_serves_ui.GetAppserverVersionResponse")
	proto.RegisterType((*GetServerRegionResponse)(nil), "appserver_serves_ui.GetServerRegionResponse")
}

func init() { proto.RegisterFile("serverInfo.proto", fileDescriptor_386dce779fe74cf3) }

var fileDescriptor_386dce779fe74cf3 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0x2c,
	0x28, 0x80, 0x08, 0xc6, 0x83, 0xa9, 0xe2, 0xf8, 0xd2, 0x4c, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4,
	0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc,
	0xbc, 0x62, 0x88, 0x16, 0x29, 0xc1, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x7d, 0x30, 0x09, 0x11, 0x52,
	0x32, 0xe7, 0x92, 0x76, 0x4f, 0x2d, 0x71, 0x84, 0x19, 0x15, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f,
	0x17, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc1, 0xc5, 0x5e, 0x06, 0x11, 0x92,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0xec, 0xb8, 0xc4, 0xdd, 0x53, 0x4b, 0x82,
	0xc1, 0xba, 0x82, 0x52, 0xd3, 0x91, 0x35, 0x29, 0x73, 0xf1, 0x42, 0x1d, 0x56, 0x04, 0x96, 0x80,
	0x6a, 0xe5, 0x29, 0x46, 0x52, 0xac, 0x65, 0xc3, 0xc5, 0x83, 0xac, 0x59, 0x88, 0x9f, 0x8b, 0xdb,
	0xcf, 0x3f, 0x24, 0xde, 0xc5, 0xd5, 0xcd, 0xd3, 0xcf, 0xd5, 0x45, 0x80, 0x41, 0x88, 0x9b, 0x8b,
	0xdd, 0x31, 0xcc, 0x35, 0xc8, 0xd1, 0xdd, 0x55, 0x80, 0x51, 0x88, 0x8f, 0x8b, 0x2b, 0xc8, 0x35,
	0x38, 0x24, 0xc8, 0xd3, 0x39, 0xc4, 0xd5, 0x45, 0x80, 0xc9, 0x68, 0x11, 0x13, 0x97, 0x60, 0x30,
	0x3c, 0x44, 0x40, 0xac, 0xcc, 0xe4, 0x54, 0xa1, 0x6e, 0x46, 0x2e, 0x61, 0x2c, 0xbe, 0x11, 0x12,
	0xd3, 0x83, 0x04, 0x0b, 0xc4, 0xcf, 0x49, 0xa5, 0x69, 0x7a, 0xae, 0xa0, 0x20, 0x90, 0x32, 0xd0,
	0xc3, 0x12, 0x86, 0x7a, 0x78, 0xc2, 0x43, 0x49, 0xab, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x2a, 0x42,
	0x4a, 0xe0, 0x10, 0x86, 0xa8, 0xd1, 0xcd, 0xcc, 0x4b, 0xcb, 0xd7, 0x87, 0x9b, 0xa4, 0x0b, 0x0d,
	0x21, 0xa1, 0x7a, 0x2e, 0x7e, 0xb4, 0x10, 0xc2, 0xe9, 0x10, 0x1d, 0x5c, 0x0e, 0xc1, 0x16, 0xbe,
	0x4a, 0x6a, 0x60, 0x47, 0x28, 0x08, 0xc9, 0x61, 0x38, 0x02, 0xca, 0x86, 0x04, 0x7b, 0x12, 0x1b,
	0xd8, 0x16, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0xbe, 0x4f, 0xbd, 0x3c, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerInfoServiceClient is the client API for ServerInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerInfoServiceClient interface {
	// get version
	GetAppserverVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAppserverVersionResponse, error)
	GetServerRegion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetServerRegionResponse, error)
}

type serverInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerInfoServiceClient(cc grpc.ClientConnInterface) ServerInfoServiceClient {
	return &serverInfoServiceClient{cc}
}

func (c *serverInfoServiceClient) GetAppserverVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAppserverVersionResponse, error) {
	out := new(GetAppserverVersionResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.ServerInfoService/GetAppserverVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInfoServiceClient) GetServerRegion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetServerRegionResponse, error) {
	out := new(GetServerRegionResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.ServerInfoService/GetServerRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerInfoServiceServer is the server API for ServerInfoService service.
type ServerInfoServiceServer interface {
	// get version
	GetAppserverVersion(context.Context, *empty.Empty) (*GetAppserverVersionResponse, error)
	GetServerRegion(context.Context, *empty.Empty) (*GetServerRegionResponse, error)
}

// UnimplementedServerInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServerInfoServiceServer struct {
}

func (*UnimplementedServerInfoServiceServer) GetAppserverVersion(ctx context.Context, req *empty.Empty) (*GetAppserverVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppserverVersion not implemented")
}
func (*UnimplementedServerInfoServiceServer) GetServerRegion(ctx context.Context, req *empty.Empty) (*GetServerRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerRegion not implemented")
}

func RegisterServerInfoServiceServer(s *grpc.Server, srv ServerInfoServiceServer) {
	s.RegisterService(&_ServerInfoService_serviceDesc, srv)
}

func _ServerInfoService_GetAppserverVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInfoServiceServer).GetAppserverVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.ServerInfoService/GetAppserverVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInfoServiceServer).GetAppserverVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInfoService_GetServerRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInfoServiceServer).GetServerRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.ServerInfoService/GetServerRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInfoServiceServer).GetServerRegion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appserver_serves_ui.ServerInfoService",
	HandlerType: (*ServerInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppserverVersion",
			Handler:    _ServerInfoService_GetAppserverVersion_Handler,
		},
		{
			MethodName: "GetServerRegion",
			Handler:    _ServerInfoService_GetServerRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverInfo.proto",
}
