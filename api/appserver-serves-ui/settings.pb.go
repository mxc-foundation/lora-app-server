// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: settings.proto

package appserver_serves_ui

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSettingsRequest) Reset() {
	*x = GetSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsRequest) ProtoMessage() {}

func (x *GetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{0}
}

type GetSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// when supernode income is lower than expected revenue, warn system owner to increase income
	LowBalanceWarning    string  `protobuf:"bytes,1,opt,name=low_balance_warning,json=lowBalanceWarning,proto3" json:"low_balance_warning,omitempty"`
	DownlinkPrice        float64 `protobuf:"fixed64,2,opt,name=downlink_price,json=downlinkPrice,proto3" json:"downlink_price,omitempty"`
	SupernodeIncomeRatio float64 `protobuf:"fixed64,3,opt,name=supernode_income_ratio,json=supernodeIncomeRatio,proto3" json:"supernode_income_ratio,omitempty"`
	// this is the percentage of reward from supernode income
	StakingPercentage float64 `protobuf:"fixed64,4,opt,name=staking_percentage,json=stakingPercentage,proto3" json:"staking_percentage,omitempty"`
	// this is the percentage of expected reward from staking amount
	StakingExpectedRevenuePercentage float64 `protobuf:"fixed64,5,opt,name=staking_expected_revenue_percentage,json=stakingExpectedRevenuePercentage,proto3" json:"staking_expected_revenue_percentage,omitempty"`
	Compensation                     float64 `protobuf:"fixed64,6,opt,name=compensation,proto3" json:"compensation,omitempty"`
}

func (x *GetSettingsResponse) Reset() {
	*x = GetSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsResponse) ProtoMessage() {}

func (x *GetSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetSettingsResponse) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{1}
}

func (x *GetSettingsResponse) GetLowBalanceWarning() string {
	if x != nil {
		return x.LowBalanceWarning
	}
	return ""
}

func (x *GetSettingsResponse) GetDownlinkPrice() float64 {
	if x != nil {
		return x.DownlinkPrice
	}
	return 0
}

func (x *GetSettingsResponse) GetSupernodeIncomeRatio() float64 {
	if x != nil {
		return x.SupernodeIncomeRatio
	}
	return 0
}

func (x *GetSettingsResponse) GetStakingPercentage() float64 {
	if x != nil {
		return x.StakingPercentage
	}
	return 0
}

func (x *GetSettingsResponse) GetStakingExpectedRevenuePercentage() float64 {
	if x != nil {
		return x.StakingExpectedRevenuePercentage
	}
	return 0
}

func (x *GetSettingsResponse) GetCompensation() float64 {
	if x != nil {
		return x.Compensation
	}
	return 0
}

type ModifySettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowBalanceWarning          *wrappers.Int64Value `protobuf:"bytes,1,opt,name=lowBalanceWarning,proto3" json:"lowBalanceWarning,omitempty"`
	DownlinkFee                *wrappers.Int64Value `protobuf:"bytes,2,opt,name=downlinkFee,proto3" json:"downlinkFee,omitempty"`
	TransactionPercentageShare *wrappers.Int64Value `protobuf:"bytes,3,opt,name=transactionPercentageShare,proto3" json:"transactionPercentageShare,omitempty"`
}

func (x *ModifySettingsRequest) Reset() {
	*x = ModifySettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySettingsRequest) ProtoMessage() {}

func (x *ModifySettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySettingsRequest.ProtoReflect.Descriptor instead.
func (*ModifySettingsRequest) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{2}
}

func (x *ModifySettingsRequest) GetLowBalanceWarning() *wrappers.Int64Value {
	if x != nil {
		return x.LowBalanceWarning
	}
	return nil
}

func (x *ModifySettingsRequest) GetDownlinkFee() *wrappers.Int64Value {
	if x != nil {
		return x.DownlinkFee
	}
	return nil
}

func (x *ModifySettingsRequest) GetTransactionPercentageShare() *wrappers.Int64Value {
	if x != nil {
		return x.TransactionPercentageShare
	}
	return nil
}

type ModifySettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ModifySettingsResponse) Reset() {
	*x = ModifySettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySettingsResponse) ProtoMessage() {}

func (x *ModifySettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySettingsResponse.ProtoReflect.Descriptor instead.
func (*ModifySettingsResponse) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{3}
}

func (x *ModifySettingsResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_settings_proto protoreflect.FileDescriptor

var file_settings_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x73, 0x5f, 0x75, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc4, 0x02, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x6f, 0x77, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x6c, 0x6f, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x2d,
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a,
	0x23, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x20, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x65, 0x6e,
	0x75, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xfe, 0x01, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x6c, 0x6f,
	0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x11, 0x6c, 0x6f, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e,
	0x6b, 0x46, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e,
	0x6b, 0x46, 0x65, 0x65, 0x12, 0x5b, 0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x22, 0x30, 0x0a, 0x16, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0x90, 0x02, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x83, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x1a, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x63, 0x2d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x70, 0x77, 0x61, 0x6e, 0x2d, 0x61, 0x70, 0x70, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x2d, 0x75, 0x69, 0x3b, 0x61, 0x70, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x75, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_settings_proto_rawDescOnce sync.Once
	file_settings_proto_rawDescData = file_settings_proto_rawDesc
)

func file_settings_proto_rawDescGZIP() []byte {
	file_settings_proto_rawDescOnce.Do(func() {
		file_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_settings_proto_rawDescData)
	})
	return file_settings_proto_rawDescData
}

var file_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_settings_proto_goTypes = []interface{}{
	(*GetSettingsRequest)(nil),     // 0: appserver_serves_ui.GetSettingsRequest
	(*GetSettingsResponse)(nil),    // 1: appserver_serves_ui.GetSettingsResponse
	(*ModifySettingsRequest)(nil),  // 2: appserver_serves_ui.ModifySettingsRequest
	(*ModifySettingsResponse)(nil), // 3: appserver_serves_ui.ModifySettingsResponse
	(*wrappers.Int64Value)(nil),    // 4: google.protobuf.Int64Value
}
var file_settings_proto_depIdxs = []int32{
	4, // 0: appserver_serves_ui.ModifySettingsRequest.lowBalanceWarning:type_name -> google.protobuf.Int64Value
	4, // 1: appserver_serves_ui.ModifySettingsRequest.downlinkFee:type_name -> google.protobuf.Int64Value
	4, // 2: appserver_serves_ui.ModifySettingsRequest.transactionPercentageShare:type_name -> google.protobuf.Int64Value
	0, // 3: appserver_serves_ui.SettingsService.GetSettings:input_type -> appserver_serves_ui.GetSettingsRequest
	2, // 4: appserver_serves_ui.SettingsService.ModifySettings:input_type -> appserver_serves_ui.ModifySettingsRequest
	1, // 5: appserver_serves_ui.SettingsService.GetSettings:output_type -> appserver_serves_ui.GetSettingsResponse
	3, // 6: appserver_serves_ui.SettingsService.ModifySettings:output_type -> appserver_serves_ui.ModifySettingsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_settings_proto_init() }
func file_settings_proto_init() {
	if File_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsRequest); i {
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
		file_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsResponse); i {
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
		file_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySettingsRequest); i {
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
		file_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySettingsResponse); i {
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
			RawDescriptor: file_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_settings_proto_goTypes,
		DependencyIndexes: file_settings_proto_depIdxs,
		MessageInfos:      file_settings_proto_msgTypes,
	}.Build()
	File_settings_proto = out.File
	file_settings_proto_rawDesc = nil
	file_settings_proto_goTypes = nil
	file_settings_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SettingsServiceClient is the client API for SettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SettingsServiceClient interface {
	GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error)
	ModifySettings(ctx context.Context, in *ModifySettingsRequest, opts ...grpc.CallOption) (*ModifySettingsResponse, error)
}

type settingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsServiceClient(cc grpc.ClientConnInterface) SettingsServiceClient {
	return &settingsServiceClient{cc}
}

func (c *settingsServiceClient) GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error) {
	out := new(GetSettingsResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.SettingsService/GetSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) ModifySettings(ctx context.Context, in *ModifySettingsRequest, opts ...grpc.CallOption) (*ModifySettingsResponse, error) {
	out := new(ModifySettingsResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.SettingsService/ModifySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServiceServer is the server API for SettingsService service.
type SettingsServiceServer interface {
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)
	ModifySettings(context.Context, *ModifySettingsRequest) (*ModifySettingsResponse, error)
}

// UnimplementedSettingsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSettingsServiceServer struct {
}

func (*UnimplementedSettingsServiceServer) GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
}
func (*UnimplementedSettingsServiceServer) ModifySettings(context.Context, *ModifySettingsRequest) (*ModifySettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySettings not implemented")
}

func RegisterSettingsServiceServer(s *grpc.Server, srv SettingsServiceServer) {
	s.RegisterService(&_SettingsService_serviceDesc, srv)
}

func _SettingsService_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.SettingsService/GetSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetSettings(ctx, req.(*GetSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_ModifySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).ModifySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.SettingsService/ModifySettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).ModifySettings(ctx, req.(*ModifySettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SettingsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appserver_serves_ui.SettingsService",
	HandlerType: (*SettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSettings",
			Handler:    _SettingsService_GetSettings_Handler,
		},
		{
			MethodName: "ModifySettings",
			Handler:    _SettingsService_ModifySettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings.proto",
}
