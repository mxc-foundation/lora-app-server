// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serviceProfile.proto

package appserver_serves_ui

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateServiceProfileRequest struct {
	// Service-profile object to create.
	ServiceProfile       *ServiceProfile `protobuf:"bytes,1,opt,name=service_profile,json=serviceProfile,proto3" json:"service_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateServiceProfileRequest) Reset()         { *m = CreateServiceProfileRequest{} }
func (m *CreateServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceProfileRequest) ProtoMessage()    {}
func (*CreateServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{0}
}

func (m *CreateServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceProfileRequest.Unmarshal(m, b)
}
func (m *CreateServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceProfileRequest.Merge(m, src)
}
func (m *CreateServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceProfileRequest.Size(m)
}
func (m *CreateServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceProfileRequest proto.InternalMessageInfo

func (m *CreateServiceProfileRequest) GetServiceProfile() *ServiceProfile {
	if m != nil {
		return m.ServiceProfile
	}
	return nil
}

type CreateServiceProfileResponse struct {
	// Service-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceProfileResponse) Reset()         { *m = CreateServiceProfileResponse{} }
func (m *CreateServiceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceProfileResponse) ProtoMessage()    {}
func (*CreateServiceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{1}
}

func (m *CreateServiceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceProfileResponse.Unmarshal(m, b)
}
func (m *CreateServiceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceProfileResponse.Marshal(b, m, deterministic)
}
func (m *CreateServiceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceProfileResponse.Merge(m, src)
}
func (m *CreateServiceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceProfileResponse.Size(m)
}
func (m *CreateServiceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceProfileResponse proto.InternalMessageInfo

func (m *CreateServiceProfileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetServiceProfileRequest struct {
	// Service-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceProfileRequest) Reset()         { *m = GetServiceProfileRequest{} }
func (m *GetServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceProfileRequest) ProtoMessage()    {}
func (*GetServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{2}
}

func (m *GetServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceProfileRequest.Unmarshal(m, b)
}
func (m *GetServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceProfileRequest.Merge(m, src)
}
func (m *GetServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceProfileRequest.Size(m)
}
func (m *GetServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceProfileRequest proto.InternalMessageInfo

func (m *GetServiceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetServiceProfileResponse struct {
	// Service-profile object.
	ServiceProfile *ServiceProfile `protobuf:"bytes,1,opt,name=service_profile,json=serviceProfile,proto3" json:"service_profile,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetServiceProfileResponse) Reset()         { *m = GetServiceProfileResponse{} }
func (m *GetServiceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceProfileResponse) ProtoMessage()    {}
func (*GetServiceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{3}
}

func (m *GetServiceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceProfileResponse.Unmarshal(m, b)
}
func (m *GetServiceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetServiceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceProfileResponse.Merge(m, src)
}
func (m *GetServiceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceProfileResponse.Size(m)
}
func (m *GetServiceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceProfileResponse proto.InternalMessageInfo

func (m *GetServiceProfileResponse) GetServiceProfile() *ServiceProfile {
	if m != nil {
		return m.ServiceProfile
	}
	return nil
}

func (m *GetServiceProfileResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetServiceProfileResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateServiceProfileRequest struct {
	// Service-profile object to update.
	ServiceProfile       *ServiceProfile `protobuf:"bytes,1,opt,name=service_profile,json=serviceProfile,proto3" json:"service_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateServiceProfileRequest) Reset()         { *m = UpdateServiceProfileRequest{} }
func (m *UpdateServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateServiceProfileRequest) ProtoMessage()    {}
func (*UpdateServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{4}
}

func (m *UpdateServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServiceProfileRequest.Unmarshal(m, b)
}
func (m *UpdateServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServiceProfileRequest.Merge(m, src)
}
func (m *UpdateServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateServiceProfileRequest.Size(m)
}
func (m *UpdateServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServiceProfileRequest proto.InternalMessageInfo

func (m *UpdateServiceProfileRequest) GetServiceProfile() *ServiceProfile {
	if m != nil {
		return m.ServiceProfile
	}
	return nil
}

type DeleteServiceProfileRequest struct {
	// Service-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServiceProfileRequest) Reset()         { *m = DeleteServiceProfileRequest{} }
func (m *DeleteServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteServiceProfileRequest) ProtoMessage()    {}
func (*DeleteServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{5}
}

func (m *DeleteServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServiceProfileRequest.Unmarshal(m, b)
}
func (m *DeleteServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServiceProfileRequest.Merge(m, src)
}
func (m *DeleteServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteServiceProfileRequest.Size(m)
}
func (m *DeleteServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServiceProfileRequest proto.InternalMessageInfo

func (m *DeleteServiceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListServiceProfileRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationId       int64    `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServiceProfileRequest) Reset()         { *m = ListServiceProfileRequest{} }
func (m *ListServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ListServiceProfileRequest) ProtoMessage()    {}
func (*ListServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{6}
}

func (m *ListServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceProfileRequest.Unmarshal(m, b)
}
func (m *ListServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *ListServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceProfileRequest.Merge(m, src)
}
func (m *ListServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ListServiceProfileRequest.Size(m)
}
func (m *ListServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceProfileRequest proto.InternalMessageInfo

func (m *ListServiceProfileRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListServiceProfileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListServiceProfileRequest) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

type ServiceProfileListItem struct {
	// Service-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Service-profile name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID of the service-profile.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID of the service-profile.
	NetworkServerId int64 `protobuf:"varint,4,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ServiceProfileListItem) Reset()         { *m = ServiceProfileListItem{} }
func (m *ServiceProfileListItem) String() string { return proto.CompactTextString(m) }
func (*ServiceProfileListItem) ProtoMessage()    {}
func (*ServiceProfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{7}
}

func (m *ServiceProfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceProfileListItem.Unmarshal(m, b)
}
func (m *ServiceProfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceProfileListItem.Marshal(b, m, deterministic)
}
func (m *ServiceProfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceProfileListItem.Merge(m, src)
}
func (m *ServiceProfileListItem) XXX_Size() int {
	return xxx_messageInfo_ServiceProfileListItem.Size(m)
}
func (m *ServiceProfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceProfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceProfileListItem proto.InternalMessageInfo

func (m *ServiceProfileListItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceProfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceProfileListItem) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *ServiceProfileListItem) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *ServiceProfileListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ServiceProfileListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ListServiceProfileResponse struct {
	// Total number of service-profiles.
	TotalCount           int64                     `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Result               []*ServiceProfileListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListServiceProfileResponse) Reset()         { *m = ListServiceProfileResponse{} }
func (m *ListServiceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ListServiceProfileResponse) ProtoMessage()    {}
func (*ListServiceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e9e4e25c67f7cd, []int{8}
}

func (m *ListServiceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceProfileResponse.Unmarshal(m, b)
}
func (m *ListServiceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceProfileResponse.Marshal(b, m, deterministic)
}
func (m *ListServiceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceProfileResponse.Merge(m, src)
}
func (m *ListServiceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ListServiceProfileResponse.Size(m)
}
func (m *ListServiceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceProfileResponse proto.InternalMessageInfo

func (m *ListServiceProfileResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListServiceProfileResponse) GetResult() []*ServiceProfileListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateServiceProfileRequest)(nil), "appserver_serves_ui.CreateServiceProfileRequest")
	proto.RegisterType((*CreateServiceProfileResponse)(nil), "appserver_serves_ui.CreateServiceProfileResponse")
	proto.RegisterType((*GetServiceProfileRequest)(nil), "appserver_serves_ui.GetServiceProfileRequest")
	proto.RegisterType((*GetServiceProfileResponse)(nil), "appserver_serves_ui.GetServiceProfileResponse")
	proto.RegisterType((*UpdateServiceProfileRequest)(nil), "appserver_serves_ui.UpdateServiceProfileRequest")
	proto.RegisterType((*DeleteServiceProfileRequest)(nil), "appserver_serves_ui.DeleteServiceProfileRequest")
	proto.RegisterType((*ListServiceProfileRequest)(nil), "appserver_serves_ui.ListServiceProfileRequest")
	proto.RegisterType((*ServiceProfileListItem)(nil), "appserver_serves_ui.ServiceProfileListItem")
	proto.RegisterType((*ListServiceProfileResponse)(nil), "appserver_serves_ui.ListServiceProfileResponse")
}

func init() { proto.RegisterFile("serviceProfile.proto", fileDescriptor_d9e9e4e25c67f7cd) }

var fileDescriptor_d9e9e4e25c67f7cd = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xe3, 0xc4, 0x52, 0x26, 0x52, 0xa2, 0x2e, 0x6d, 0x48, 0xdc, 0x20, 0x2a, 0x73, 0xa0,
	0x0a, 0x8a, 0x0d, 0x41, 0x1c, 0xe0, 0x56, 0x25, 0xa8, 0x8a, 0xd4, 0x03, 0x72, 0xe1, 0x6c, 0xb9,
	0xf1, 0x24, 0x5a, 0xd5, 0xf6, 0x1a, 0x7b, 0x5d, 0x04, 0xa8, 0x97, 0x1e, 0xf8, 0x39, 0x71, 0xe0,
	0x89, 0x78, 0x06, 0x1e, 0x80, 0x0b, 0x0f, 0x82, 0xbc, 0xde, 0x20, 0x62, 0xec, 0x90, 0x22, 0xc4,
	0x25, 0xf1, 0xee, 0xce, 0xcc, 0xf7, 0xed, 0x37, 0xdf, 0x0e, 0xec, 0x26, 0x18, 0x5f, 0xd0, 0x39,
	0x3e, 0x8b, 0xd9, 0x82, 0xfa, 0x68, 0x46, 0x31, 0xe3, 0x8c, 0xdc, 0x70, 0xa3, 0x28, 0x3b, 0xc0,
	0xd8, 0x11, 0x7f, 0x89, 0x93, 0x52, 0x7d, 0xb0, 0x64, 0x6c, 0xe9, 0xa3, 0xe5, 0x46, 0xd4, 0x72,
	0xc3, 0x90, 0x71, 0x97, 0x53, 0x16, 0x26, 0x79, 0x8a, 0xde, 0xe7, 0x34, 0xc0, 0x84, 0xbb, 0x41,
	0x64, 0xfd, 0xfc, 0x92, 0x47, 0x3b, 0x18, 0x44, 0xfc, 0xb5, 0x25, 0x7e, 0xe5, 0x56, 0x3b, 0xca,
	0xf1, 0x64, 0xb6, 0x71, 0x0e, 0xfb, 0x93, 0x18, 0x5d, 0x8e, 0xa7, 0x6b, 0x74, 0x6c, 0x7c, 0x99,
	0x62, 0xc2, 0xc9, 0x09, 0x74, 0x24, 0x4f, 0x47, 0x26, 0xf6, 0x94, 0x03, 0xe5, 0xb0, 0x35, 0xbe,
	0x63, 0x96, 0x30, 0x35, 0x0b, 0x45, 0xda, 0xeb, 0x77, 0x34, 0x4c, 0x18, 0x94, 0x83, 0x25, 0x11,
	0x0b, 0x13, 0x24, 0x6d, 0xa8, 0x51, 0x4f, 0x00, 0x34, 0xed, 0x1a, 0xf5, 0x8c, 0x21, 0xf4, 0x8e,
	0x91, 0x97, 0x33, 0x2b, 0xc6, 0x7e, 0x53, 0xa0, 0x5f, 0x12, 0x2c, 0x2b, 0xff, 0xd3, 0x7b, 0x90,
	0xc7, 0x00, 0x73, 0x71, 0x0f, 0xcf, 0x71, 0x79, 0xaf, 0x26, 0x0a, 0xe9, 0x66, 0xde, 0xa5, 0x5c,
	0xd7, 0xb3, 0x74, 0x61, 0x3e, 0x5f, 0x75, 0xc3, 0x6e, 0xca, 0xe8, 0x23, 0x9e, 0xa5, 0xa6, 0x91,
	0xb7, 0x4a, 0x55, 0xff, 0x9c, 0x2a, 0xa3, 0x8f, 0x78, 0xd6, 0xaa, 0x17, 0x62, 0xf1, 0x3f, 0x5a,
	0x35, 0x82, 0xfd, 0x29, 0xfa, 0x58, 0x05, 0x56, 0x54, 0x3f, 0x86, 0xfe, 0x09, 0x4d, 0x2a, 0x5a,
	0xb5, 0x0b, 0x0d, 0x9f, 0x06, 0x94, 0x8b, 0x78, 0xd5, 0xce, 0x17, 0xa4, 0x0b, 0x1a, 0x5b, 0x2c,
	0x12, 0xcc, 0x05, 0x54, 0x6d, 0xb9, 0x22, 0x77, 0xa1, 0xc3, 0xe2, 0xa5, 0x1b, 0xd2, 0x37, 0xc2,
	0xe6, 0x0e, 0xf5, 0x84, 0x4c, 0xaa, 0xdd, 0xfe, 0x75, 0x7b, 0x36, 0x35, 0xde, 0xd7, 0xa0, 0xbb,
	0x0e, 0x98, 0x51, 0x98, 0x71, 0x0c, 0x8a, 0xf4, 0x08, 0x81, 0x7a, 0xe8, 0x06, 0x28, 0x90, 0x9a,
	0xb6, 0xf8, 0xde, 0x1a, 0x87, 0x0c, 0x61, 0x27, 0x44, 0xfe, 0x8a, 0xc5, 0xe7, 0x8e, 0x54, 0x91,
	0x7a, 0xbd, 0xba, 0x08, 0xed, 0xc8, 0x83, 0x53, 0xb1, 0x3f, 0x9b, 0x16, 0x9c, 0xd1, 0xf8, 0x7b,
	0x67, 0x68, 0xd7, 0x71, 0xc6, 0x95, 0x02, 0x7a, 0x99, 0xfc, 0xd2, 0xfc, 0xb7, 0xa1, 0xc5, 0x19,
	0x77, 0x7d, 0x67, 0xce, 0xd2, 0x70, 0xd5, 0x05, 0x10, 0x5b, 0x93, 0x6c, 0x87, 0x4c, 0x40, 0x8b,
	0x31, 0x49, 0xfd, 0xac, 0x15, 0xea, 0x61, 0x6b, 0x7c, 0x6f, 0x0b, 0xc7, 0xac, 0xb4, 0xb6, 0x65,
	0xea, 0xf8, 0x4b, 0x03, 0xf6, 0xd6, 0x43, 0xe4, 0x8a, 0x7c, 0x52, 0x40, 0xcb, 0xdf, 0x3d, 0xb9,
	0x5f, 0x5a, 0x79, 0xc3, 0x04, 0xd2, 0x1f, 0x5c, 0x23, 0x23, 0xbf, 0xaf, 0x71, 0x70, 0xf5, 0xf5,
	0xfb, 0xe7, 0x9a, 0x6e, 0xec, 0x89, 0x89, 0x29, 0x8d, 0x3d, 0x5a, 0x0d, 0xbe, 0x27, 0xca, 0x90,
	0x7c, 0x50, 0x40, 0x3d, 0x46, 0x4e, 0x46, 0xa5, 0xc5, 0xab, 0x66, 0x8e, 0x6e, 0x6e, 0x1b, 0x2e,
	0x89, 0x18, 0x82, 0xc8, 0x80, 0xe8, 0xa5, 0x44, 0xac, 0xb7, 0xd4, 0xbb, 0x24, 0x1f, 0x15, 0xd0,
	0xf2, 0x67, 0x5d, 0x21, 0xce, 0x86, 0x37, 0xaf, 0x77, 0x7f, 0xf3, 0xc7, 0xd3, 0x6c, 0xd6, 0x1b,
	0x8f, 0x04, 0xb0, 0xa5, 0x0f, 0x2b, 0x80, 0x0b, 0x83, 0xc2, 0xa4, 0xde, 0x65, 0x26, 0xcb, 0x05,
	0x68, 0xf9, 0xa3, 0xaf, 0xa0, 0xb2, 0x61, 0x22, 0x54, 0x52, 0x91, 0x1a, 0x0c, 0x37, 0x69, 0xf0,
	0x4e, 0x81, 0x7a, 0xe6, 0x27, 0x52, 0x2e, 0x70, 0xe5, 0x64, 0xd1, 0xad, 0xad, 0xe3, 0x65, 0x47,
	0x6e, 0x09, 0x36, 0x37, 0x49, 0xb9, 0x35, 0xce, 0x34, 0x41, 0xfe, 0xe1, 0x8f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0xe9, 0x0a, 0x07, 0x9d, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceProfileServiceClient is the client API for ServiceProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceProfileServiceClient interface {
	// Create creates the given service-profile.
	Create(ctx context.Context, in *CreateServiceProfileRequest, opts ...grpc.CallOption) (*CreateServiceProfileResponse, error)
	// Get returns the service-profile matching the given id.
	Get(ctx context.Context, in *GetServiceProfileRequest, opts ...grpc.CallOption) (*GetServiceProfileResponse, error)
	// Update updates the given serviceprofile.
	Update(ctx context.Context, in *UpdateServiceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the service-profile matching the given id.
	Delete(ctx context.Context, in *DeleteServiceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available service-profiles.
	List(ctx context.Context, in *ListServiceProfileRequest, opts ...grpc.CallOption) (*ListServiceProfileResponse, error)
}

type serviceProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceProfileServiceClient(cc grpc.ClientConnInterface) ServiceProfileServiceClient {
	return &serviceProfileServiceClient{cc}
}

func (c *serviceProfileServiceClient) Create(ctx context.Context, in *CreateServiceProfileRequest, opts ...grpc.CallOption) (*CreateServiceProfileResponse, error) {
	out := new(CreateServiceProfileResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.ServiceProfileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProfileServiceClient) Get(ctx context.Context, in *GetServiceProfileRequest, opts ...grpc.CallOption) (*GetServiceProfileResponse, error) {
	out := new(GetServiceProfileResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.ServiceProfileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProfileServiceClient) Update(ctx context.Context, in *UpdateServiceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.ServiceProfileService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProfileServiceClient) Delete(ctx context.Context, in *DeleteServiceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.ServiceProfileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProfileServiceClient) List(ctx context.Context, in *ListServiceProfileRequest, opts ...grpc.CallOption) (*ListServiceProfileResponse, error) {
	out := new(ListServiceProfileResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.ServiceProfileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceProfileServiceServer is the server API for ServiceProfileService service.
type ServiceProfileServiceServer interface {
	// Create creates the given service-profile.
	Create(context.Context, *CreateServiceProfileRequest) (*CreateServiceProfileResponse, error)
	// Get returns the service-profile matching the given id.
	Get(context.Context, *GetServiceProfileRequest) (*GetServiceProfileResponse, error)
	// Update updates the given serviceprofile.
	Update(context.Context, *UpdateServiceProfileRequest) (*empty.Empty, error)
	// Delete deletes the service-profile matching the given id.
	Delete(context.Context, *DeleteServiceProfileRequest) (*empty.Empty, error)
	// List lists the available service-profiles.
	List(context.Context, *ListServiceProfileRequest) (*ListServiceProfileResponse, error)
}

// UnimplementedServiceProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceProfileServiceServer struct {
}

func (*UnimplementedServiceProfileServiceServer) Create(ctx context.Context, req *CreateServiceProfileRequest) (*CreateServiceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedServiceProfileServiceServer) Get(ctx context.Context, req *GetServiceProfileRequest) (*GetServiceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedServiceProfileServiceServer) Update(ctx context.Context, req *UpdateServiceProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedServiceProfileServiceServer) Delete(ctx context.Context, req *DeleteServiceProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedServiceProfileServiceServer) List(ctx context.Context, req *ListServiceProfileRequest) (*ListServiceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterServiceProfileServiceServer(s *grpc.Server, srv ServiceProfileServiceServer) {
	s.RegisterService(&_ServiceProfileService_serviceDesc, srv)
}

func _ServiceProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.ServiceProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).Create(ctx, req.(*CreateServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.ServiceProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).Get(ctx, req.(*GetServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.ServiceProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).Update(ctx, req.(*UpdateServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.ServiceProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).Delete(ctx, req.(*DeleteServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.ServiceProfileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).List(ctx, req.(*ListServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appserver_serves_ui.ServiceProfileService",
	HandlerType: (*ServiceProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ServiceProfileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServiceProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServiceProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ServiceProfileService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ServiceProfileService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serviceProfile.proto",
}
