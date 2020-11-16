// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: m2m_gateway.proto

package appserver_serves_m2m

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AppServerGatewayProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac         string               `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	OrgId       int64                `protobuf:"varint,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Name        string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *AppServerGatewayProfile) Reset() {
	*x = AppServerGatewayProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppServerGatewayProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppServerGatewayProfile) ProtoMessage() {}

func (x *AppServerGatewayProfile) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppServerGatewayProfile.ProtoReflect.Descriptor instead.
func (*AppServerGatewayProfile) Descriptor() ([]byte, []int) {
	return file_m2m_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *AppServerGatewayProfile) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *AppServerGatewayProfile) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *AppServerGatewayProfile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppServerGatewayProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppServerGatewayProfile) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetGatewayByMacRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac string `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
}

func (x *GetGatewayByMacRequest) Reset() {
	*x = GetGatewayByMacRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGatewayByMacRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGatewayByMacRequest) ProtoMessage() {}

func (x *GetGatewayByMacRequest) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGatewayByMacRequest.ProtoReflect.Descriptor instead.
func (*GetGatewayByMacRequest) Descriptor() ([]byte, []int) {
	return file_m2m_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *GetGatewayByMacRequest) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

type GetGatewayByMacResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId     int64                    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	GwProfile *AppServerGatewayProfile `protobuf:"bytes,2,opt,name=gw_profile,json=gwProfile,proto3" json:"gw_profile,omitempty"`
}

func (x *GetGatewayByMacResponse) Reset() {
	*x = GetGatewayByMacResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGatewayByMacResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGatewayByMacResponse) ProtoMessage() {}

func (x *GetGatewayByMacResponse) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGatewayByMacResponse.ProtoReflect.Descriptor instead.
func (*GetGatewayByMacResponse) Descriptor() ([]byte, []int) {
	return file_m2m_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *GetGatewayByMacResponse) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *GetGatewayByMacResponse) GetGwProfile() *AppServerGatewayProfile {
	if x != nil {
		return x.GwProfile
	}
	return nil
}

type GetGatewayMacListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayMac []string `protobuf:"bytes,1,rep,name=gateway_mac,json=gatewayMac,proto3" json:"gateway_mac,omitempty"`
}

func (x *GetGatewayMacListResponse) Reset() {
	*x = GetGatewayMacListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGatewayMacListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGatewayMacListResponse) ProtoMessage() {}

func (x *GetGatewayMacListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGatewayMacListResponse.ProtoReflect.Descriptor instead.
func (*GetGatewayMacListResponse) Descriptor() ([]byte, []int) {
	return file_m2m_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *GetGatewayMacListResponse) GetGatewayMac() []string {
	if x != nil {
		return x.GatewayMac
	}
	return nil
}

var File_m2m_gateway_proto protoreflect.FileDescriptor

var file_m2m_gateway_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x32, 0x6d, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x6d, 0x32, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x17, 0x41, 0x70, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2a, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x79, 0x4d, 0x61, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x22, 0x7e, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x0a, 0x67,
	0x77, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x73, 0x5f, 0x6d, 0x32, 0x6d, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x09,
	0x67, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x3c, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x61, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x4d, 0x61, 0x63, 0x32, 0xe1, 0x01, 0x0a, 0x11, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x4d, 0x32, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x79, 0x4d, 0x61, 0x63,
	0x12, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x73, 0x5f, 0x6d, 0x32, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x42, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x73, 0x5f, 0x6d, 0x32, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x42, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x61, 0x63, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x6d, 0x32,
	0x6d, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x61, 0x63, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5a, 0x5a, 0x58, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x63, 0x2d, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x70, 0x77, 0x61, 0x6e, 0x2d, 0x61,
	0x70, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x2d, 0x6d,
	0x32, 0x6d, 0x3b, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x73, 0x5f, 0x6d, 0x32, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_m2m_gateway_proto_rawDescOnce sync.Once
	file_m2m_gateway_proto_rawDescData = file_m2m_gateway_proto_rawDesc
)

func file_m2m_gateway_proto_rawDescGZIP() []byte {
	file_m2m_gateway_proto_rawDescOnce.Do(func() {
		file_m2m_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_m2m_gateway_proto_rawDescData)
	})
	return file_m2m_gateway_proto_rawDescData
}

var file_m2m_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_m2m_gateway_proto_goTypes = []interface{}{
	(*AppServerGatewayProfile)(nil),   // 0: appserver_serves_m2m.AppServerGatewayProfile
	(*GetGatewayByMacRequest)(nil),    // 1: appserver_serves_m2m.GetGatewayByMacRequest
	(*GetGatewayByMacResponse)(nil),   // 2: appserver_serves_m2m.GetGatewayByMacResponse
	(*GetGatewayMacListResponse)(nil), // 3: appserver_serves_m2m.GetGatewayMacListResponse
	(*timestamp.Timestamp)(nil),       // 4: google.protobuf.Timestamp
	(*empty.Empty)(nil),               // 5: google.protobuf.Empty
}
var file_m2m_gateway_proto_depIdxs = []int32{
	4, // 0: appserver_serves_m2m.AppServerGatewayProfile.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: appserver_serves_m2m.GetGatewayByMacResponse.gw_profile:type_name -> appserver_serves_m2m.AppServerGatewayProfile
	1, // 2: appserver_serves_m2m.GatewayM2MService.GetGatewayByMac:input_type -> appserver_serves_m2m.GetGatewayByMacRequest
	5, // 3: appserver_serves_m2m.GatewayM2MService.GetGatewayMacList:input_type -> google.protobuf.Empty
	2, // 4: appserver_serves_m2m.GatewayM2MService.GetGatewayByMac:output_type -> appserver_serves_m2m.GetGatewayByMacResponse
	3, // 5: appserver_serves_m2m.GatewayM2MService.GetGatewayMacList:output_type -> appserver_serves_m2m.GetGatewayMacListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_m2m_gateway_proto_init() }
func file_m2m_gateway_proto_init() {
	if File_m2m_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_m2m_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppServerGatewayProfile); i {
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
		file_m2m_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGatewayByMacRequest); i {
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
		file_m2m_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGatewayByMacResponse); i {
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
		file_m2m_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGatewayMacListResponse); i {
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
			RawDescriptor: file_m2m_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_m2m_gateway_proto_goTypes,
		DependencyIndexes: file_m2m_gateway_proto_depIdxs,
		MessageInfos:      file_m2m_gateway_proto_msgTypes,
	}.Build()
	File_m2m_gateway_proto = out.File
	file_m2m_gateway_proto_rawDesc = nil
	file_m2m_gateway_proto_goTypes = nil
	file_m2m_gateway_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewayM2MServiceClient is the client API for GatewayM2MService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayM2MServiceClient interface {
	GetGatewayByMac(ctx context.Context, in *GetGatewayByMacRequest, opts ...grpc.CallOption) (*GetGatewayByMacResponse, error)
	GetGatewayMacList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetGatewayMacListResponse, error)
}

type gatewayM2MServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayM2MServiceClient(cc grpc.ClientConnInterface) GatewayM2MServiceClient {
	return &gatewayM2MServiceClient{cc}
}

func (c *gatewayM2MServiceClient) GetGatewayByMac(ctx context.Context, in *GetGatewayByMacRequest, opts ...grpc.CallOption) (*GetGatewayByMacResponse, error) {
	out := new(GetGatewayByMacResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_m2m.GatewayM2MService/GetGatewayByMac", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayM2MServiceClient) GetGatewayMacList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetGatewayMacListResponse, error) {
	out := new(GetGatewayMacListResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_m2m.GatewayM2MService/GetGatewayMacList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayM2MServiceServer is the server API for GatewayM2MService service.
type GatewayM2MServiceServer interface {
	GetGatewayByMac(context.Context, *GetGatewayByMacRequest) (*GetGatewayByMacResponse, error)
	GetGatewayMacList(context.Context, *empty.Empty) (*GetGatewayMacListResponse, error)
}

// UnimplementedGatewayM2MServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayM2MServiceServer struct {
}

func (*UnimplementedGatewayM2MServiceServer) GetGatewayByMac(context.Context, *GetGatewayByMacRequest) (*GetGatewayByMacResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayByMac not implemented")
}
func (*UnimplementedGatewayM2MServiceServer) GetGatewayMacList(context.Context, *empty.Empty) (*GetGatewayMacListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayMacList not implemented")
}

func RegisterGatewayM2MServiceServer(s *grpc.Server, srv GatewayM2MServiceServer) {
	s.RegisterService(&_GatewayM2MService_serviceDesc, srv)
}

func _GatewayM2MService_GetGatewayByMac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayByMacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayM2MServiceServer).GetGatewayByMac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_m2m.GatewayM2MService/GetGatewayByMac",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayM2MServiceServer).GetGatewayByMac(ctx, req.(*GetGatewayByMacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayM2MService_GetGatewayMacList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayM2MServiceServer).GetGatewayMacList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_m2m.GatewayM2MService/GetGatewayMacList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayM2MServiceServer).GetGatewayMacList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayM2MService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appserver_serves_m2m.GatewayM2MService",
	HandlerType: (*GatewayM2MServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayByMac",
			Handler:    _GatewayM2MService_GetGatewayByMac_Handler,
		},
		{
			MethodName: "GetGatewayMacList",
			Handler:    _GatewayM2MService_GetGatewayMacList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "m2m_gateway.proto",
}
