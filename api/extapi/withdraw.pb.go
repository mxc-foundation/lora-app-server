// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.13.0
// source: withdraw.proto

package extapi

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetWithdrawFeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of crypto currency
	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *GetWithdrawFeeRequest) Reset() {
	*x = GetWithdrawFeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawFeeRequest) ProtoMessage() {}

func (x *GetWithdrawFeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawFeeRequest.ProtoReflect.Descriptor instead.
func (*GetWithdrawFeeRequest) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{0}
}

func (x *GetWithdrawFeeRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type GetWithdrawFeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Withdraw object.
	WithdrawFee string `protobuf:"bytes,1,opt,name=withdrawFee,proto3" json:"withdrawFee,omitempty"`
	// actual currency of the withdraw fee
	// for BTC, withdraw fee should be in MXC
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *GetWithdrawFeeResponse) Reset() {
	*x = GetWithdrawFeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawFeeResponse) ProtoMessage() {}

func (x *GetWithdrawFeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawFeeResponse.ProtoReflect.Descriptor instead.
func (*GetWithdrawFeeResponse) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{1}
}

func (x *GetWithdrawFeeResponse) GetWithdrawFee() string {
	if x != nil {
		return x.WithdrawFee
	}
	return ""
}

func (x *GetWithdrawFeeResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type GetWithdrawHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId    int64                `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Currency string               `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	From     *timestamp.Timestamp `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	Till     *timestamp.Timestamp `protobuf:"bytes,7,opt,name=till,proto3" json:"till,omitempty"`
}

func (x *GetWithdrawHistoryRequest) Reset() {
	*x = GetWithdrawHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawHistoryRequest) ProtoMessage() {}

func (x *GetWithdrawHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetWithdrawHistoryRequest) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{2}
}

func (x *GetWithdrawHistoryRequest) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *GetWithdrawHistoryRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *GetWithdrawHistoryRequest) GetFrom() *timestamp.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *GetWithdrawHistoryRequest) GetTill() *timestamp.Timestamp {
	if x != nil {
		return x.Till
	}
	return nil
}

type WithdrawHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxStatus    string               `protobuf:"bytes,3,opt,name=tx_status,json=txStatus,proto3" json:"tx_status,omitempty"`
	TxHash      string               `protobuf:"bytes,4,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	DenyComment string               `protobuf:"bytes,5,opt,name=deny_comment,json=denyComment,proto3" json:"deny_comment,omitempty"`
	Amount      string               `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	WithdrawFee string               `protobuf:"bytes,8,opt,name=withdraw_fee,json=withdrawFee,proto3" json:"withdraw_fee,omitempty"`
}

func (x *WithdrawHistory) Reset() {
	*x = WithdrawHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawHistory) ProtoMessage() {}

func (x *WithdrawHistory) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawHistory.ProtoReflect.Descriptor instead.
func (*WithdrawHistory) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{3}
}

func (x *WithdrawHistory) GetTxStatus() string {
	if x != nil {
		return x.TxStatus
	}
	return ""
}

func (x *WithdrawHistory) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *WithdrawHistory) GetDenyComment() string {
	if x != nil {
		return x.DenyComment
	}
	return ""
}

func (x *WithdrawHistory) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *WithdrawHistory) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *WithdrawHistory) GetWithdrawFee() string {
	if x != nil {
		return x.WithdrawFee
	}
	return ""
}

type GetWithdrawHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithdrawHistory []*WithdrawHistory `protobuf:"bytes,2,rep,name=withdraw_history,json=withdrawHistory,proto3" json:"withdraw_history,omitempty"`
}

func (x *GetWithdrawHistoryResponse) Reset() {
	*x = GetWithdrawHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawHistoryResponse) ProtoMessage() {}

func (x *GetWithdrawHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetWithdrawHistoryResponse) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{4}
}

func (x *GetWithdrawHistoryResponse) GetWithdrawHistory() []*WithdrawHistory {
	if x != nil {
		return x.WithdrawHistory
	}
	return nil
}

type ModifyWithdrawFeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency    string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	WithdrawFee string `protobuf:"bytes,2,opt,name=withdraw_fee,json=withdrawFee,proto3" json:"withdraw_fee,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ModifyWithdrawFeeRequest) Reset() {
	*x = ModifyWithdrawFeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyWithdrawFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyWithdrawFeeRequest) ProtoMessage() {}

func (x *ModifyWithdrawFeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyWithdrawFeeRequest.ProtoReflect.Descriptor instead.
func (*ModifyWithdrawFeeRequest) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{5}
}

func (x *ModifyWithdrawFeeRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ModifyWithdrawFeeRequest) GetWithdrawFee() string {
	if x != nil {
		return x.WithdrawFee
	}
	return ""
}

func (x *ModifyWithdrawFeeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ModifyWithdrawFeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ModifyWithdrawFeeResponse) Reset() {
	*x = ModifyWithdrawFeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyWithdrawFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyWithdrawFeeResponse) ProtoMessage() {}

func (x *ModifyWithdrawFeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyWithdrawFeeResponse.ProtoReflect.Descriptor instead.
func (*ModifyWithdrawFeeResponse) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{6}
}

func (x *ModifyWithdrawFeeResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetWithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId      int64  `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	EthAddress string `protobuf:"bytes,3,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
	Currency   string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount     string `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GetWithdrawRequest) Reset() {
	*x = GetWithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawRequest) ProtoMessage() {}

func (x *GetWithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawRequest.ProtoReflect.Descriptor instead.
func (*GetWithdrawRequest) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{7}
}

func (x *GetWithdrawRequest) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *GetWithdrawRequest) GetEthAddress() string {
	if x != nil {
		return x.EthAddress
	}
	return ""
}

func (x *GetWithdrawRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *GetWithdrawRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type GetWithdrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetWithdrawResponse) Reset() {
	*x = GetWithdrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_withdraw_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawResponse) ProtoMessage() {}

func (x *GetWithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_withdraw_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawResponse.ProtoReflect.Descriptor instead.
func (*GetWithdrawResponse) Descriptor() ([]byte, []int) {
	return file_withdraw_proto_rawDescGZIP(), []int{8}
}

func (x *GetWithdrawResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_withdraw_proto protoreflect.FileDescriptor

var file_withdraw_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x65, 0x78, 0x74, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x56, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x46, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0xae, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6c, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6c, 0x6c, 0x22, 0xdf, 0x01, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x78, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x78,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x6e, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x5f, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x22, 0x60, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x65, 0x78, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x75, 0x0a, 0x18, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x66, 0x65,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x46, 0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x33, 0x0a, 0x19, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x74, 0x68, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x74, 0x68, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xf5, 0x03, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x12, 0x1d, 0x2e,
	0x65, 0x78, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65,
	0x78, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x2d, 0x66, 0x65, 0x65, 0x12, 0x7a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x65, 0x78, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x86, 0x01, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x12, 0x20, 0x2e, 0x65, 0x78, 0x74, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x78, 0x74, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x26, 0x1a, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x2f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x2d, 0x66, 0x65, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x1a, 0x2e, 0x65, 0x78, 0x74, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x78, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2f, 0x72, 0x65, 0x71, 0x3a, 0x01, 0x2a, 0x42,
	0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78,
	0x63, 0x2d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x70, 0x77,
	0x61, 0x6e, 0x2d, 0x61, 0x70, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x78, 0x74, 0x61, 0x70, 0x69, 0x3b, 0x65, 0x78, 0x74, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_withdraw_proto_rawDescOnce sync.Once
	file_withdraw_proto_rawDescData = file_withdraw_proto_rawDesc
)

func file_withdraw_proto_rawDescGZIP() []byte {
	file_withdraw_proto_rawDescOnce.Do(func() {
		file_withdraw_proto_rawDescData = protoimpl.X.CompressGZIP(file_withdraw_proto_rawDescData)
	})
	return file_withdraw_proto_rawDescData
}

var file_withdraw_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_withdraw_proto_goTypes = []interface{}{
	(*GetWithdrawFeeRequest)(nil),      // 0: extapi.GetWithdrawFeeRequest
	(*GetWithdrawFeeResponse)(nil),     // 1: extapi.GetWithdrawFeeResponse
	(*GetWithdrawHistoryRequest)(nil),  // 2: extapi.GetWithdrawHistoryRequest
	(*WithdrawHistory)(nil),            // 3: extapi.WithdrawHistory
	(*GetWithdrawHistoryResponse)(nil), // 4: extapi.GetWithdrawHistoryResponse
	(*ModifyWithdrawFeeRequest)(nil),   // 5: extapi.ModifyWithdrawFeeRequest
	(*ModifyWithdrawFeeResponse)(nil),  // 6: extapi.ModifyWithdrawFeeResponse
	(*GetWithdrawRequest)(nil),         // 7: extapi.GetWithdrawRequest
	(*GetWithdrawResponse)(nil),        // 8: extapi.GetWithdrawResponse
	(*timestamp.Timestamp)(nil),        // 9: google.protobuf.Timestamp
}
var file_withdraw_proto_depIdxs = []int32{
	9, // 0: extapi.GetWithdrawHistoryRequest.from:type_name -> google.protobuf.Timestamp
	9, // 1: extapi.GetWithdrawHistoryRequest.till:type_name -> google.protobuf.Timestamp
	9, // 2: extapi.WithdrawHistory.timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: extapi.GetWithdrawHistoryResponse.withdraw_history:type_name -> extapi.WithdrawHistory
	0, // 4: extapi.WithdrawService.GetWithdrawFee:input_type -> extapi.GetWithdrawFeeRequest
	2, // 5: extapi.WithdrawService.GetWithdrawHistory:input_type -> extapi.GetWithdrawHistoryRequest
	5, // 6: extapi.WithdrawService.ModifyWithdrawFee:input_type -> extapi.ModifyWithdrawFeeRequest
	7, // 7: extapi.WithdrawService.GetWithdraw:input_type -> extapi.GetWithdrawRequest
	1, // 8: extapi.WithdrawService.GetWithdrawFee:output_type -> extapi.GetWithdrawFeeResponse
	4, // 9: extapi.WithdrawService.GetWithdrawHistory:output_type -> extapi.GetWithdrawHistoryResponse
	6, // 10: extapi.WithdrawService.ModifyWithdrawFee:output_type -> extapi.ModifyWithdrawFeeResponse
	8, // 11: extapi.WithdrawService.GetWithdraw:output_type -> extapi.GetWithdrawResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_withdraw_proto_init() }
func file_withdraw_proto_init() {
	if File_withdraw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_withdraw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawFeeRequest); i {
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
		file_withdraw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawFeeResponse); i {
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
		file_withdraw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawHistoryRequest); i {
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
		file_withdraw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawHistory); i {
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
		file_withdraw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawHistoryResponse); i {
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
		file_withdraw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyWithdrawFeeRequest); i {
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
		file_withdraw_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyWithdrawFeeResponse); i {
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
		file_withdraw_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawRequest); i {
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
		file_withdraw_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawResponse); i {
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
			RawDescriptor: file_withdraw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_withdraw_proto_goTypes,
		DependencyIndexes: file_withdraw_proto_depIdxs,
		MessageInfos:      file_withdraw_proto_msgTypes,
	}.Build()
	File_withdraw_proto = out.File
	file_withdraw_proto_rawDesc = nil
	file_withdraw_proto_goTypes = nil
	file_withdraw_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WithdrawServiceClient is the client API for WithdrawService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WithdrawServiceClient interface {
	// GetWithdrawFee data for current withdraw fee
	GetWithdrawFee(ctx context.Context, in *GetWithdrawFeeRequest, opts ...grpc.CallOption) (*GetWithdrawFeeResponse, error)
	GetWithdrawHistory(ctx context.Context, in *GetWithdrawHistoryRequest, opts ...grpc.CallOption) (*GetWithdrawHistoryResponse, error)
	ModifyWithdrawFee(ctx context.Context, in *ModifyWithdrawFeeRequest, opts ...grpc.CallOption) (*ModifyWithdrawFeeResponse, error)
	// after user clicks withdraw, send withdraw request to cobo directly
	GetWithdraw(ctx context.Context, in *GetWithdrawRequest, opts ...grpc.CallOption) (*GetWithdrawResponse, error)
}

type withdrawServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWithdrawServiceClient(cc grpc.ClientConnInterface) WithdrawServiceClient {
	return &withdrawServiceClient{cc}
}

func (c *withdrawServiceClient) GetWithdrawFee(ctx context.Context, in *GetWithdrawFeeRequest, opts ...grpc.CallOption) (*GetWithdrawFeeResponse, error) {
	out := new(GetWithdrawFeeResponse)
	err := c.cc.Invoke(ctx, "/extapi.WithdrawService/GetWithdrawFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawServiceClient) GetWithdrawHistory(ctx context.Context, in *GetWithdrawHistoryRequest, opts ...grpc.CallOption) (*GetWithdrawHistoryResponse, error) {
	out := new(GetWithdrawHistoryResponse)
	err := c.cc.Invoke(ctx, "/extapi.WithdrawService/GetWithdrawHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawServiceClient) ModifyWithdrawFee(ctx context.Context, in *ModifyWithdrawFeeRequest, opts ...grpc.CallOption) (*ModifyWithdrawFeeResponse, error) {
	out := new(ModifyWithdrawFeeResponse)
	err := c.cc.Invoke(ctx, "/extapi.WithdrawService/ModifyWithdrawFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *withdrawServiceClient) GetWithdraw(ctx context.Context, in *GetWithdrawRequest, opts ...grpc.CallOption) (*GetWithdrawResponse, error) {
	out := new(GetWithdrawResponse)
	err := c.cc.Invoke(ctx, "/extapi.WithdrawService/GetWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WithdrawServiceServer is the server API for WithdrawService service.
type WithdrawServiceServer interface {
	// GetWithdrawFee data for current withdraw fee
	GetWithdrawFee(context.Context, *GetWithdrawFeeRequest) (*GetWithdrawFeeResponse, error)
	GetWithdrawHistory(context.Context, *GetWithdrawHistoryRequest) (*GetWithdrawHistoryResponse, error)
	ModifyWithdrawFee(context.Context, *ModifyWithdrawFeeRequest) (*ModifyWithdrawFeeResponse, error)
	// after user clicks withdraw, send withdraw request to cobo directly
	GetWithdraw(context.Context, *GetWithdrawRequest) (*GetWithdrawResponse, error)
}

// UnimplementedWithdrawServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWithdrawServiceServer struct {
}

func (*UnimplementedWithdrawServiceServer) GetWithdrawFee(context.Context, *GetWithdrawFeeRequest) (*GetWithdrawFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawFee not implemented")
}
func (*UnimplementedWithdrawServiceServer) GetWithdrawHistory(context.Context, *GetWithdrawHistoryRequest) (*GetWithdrawHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawHistory not implemented")
}
func (*UnimplementedWithdrawServiceServer) ModifyWithdrawFee(context.Context, *ModifyWithdrawFeeRequest) (*ModifyWithdrawFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyWithdrawFee not implemented")
}
func (*UnimplementedWithdrawServiceServer) GetWithdraw(context.Context, *GetWithdrawRequest) (*GetWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdraw not implemented")
}

func RegisterWithdrawServiceServer(s *grpc.Server, srv WithdrawServiceServer) {
	s.RegisterService(&_WithdrawService_serviceDesc, srv)
}

func _WithdrawService_GetWithdrawFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServiceServer).GetWithdrawFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extapi.WithdrawService/GetWithdrawFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServiceServer).GetWithdrawFee(ctx, req.(*GetWithdrawFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WithdrawService_GetWithdrawHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServiceServer).GetWithdrawHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extapi.WithdrawService/GetWithdrawHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServiceServer).GetWithdrawHistory(ctx, req.(*GetWithdrawHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WithdrawService_ModifyWithdrawFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyWithdrawFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServiceServer).ModifyWithdrawFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extapi.WithdrawService/ModifyWithdrawFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServiceServer).ModifyWithdrawFee(ctx, req.(*ModifyWithdrawFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WithdrawService_GetWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithdrawServiceServer).GetWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extapi.WithdrawService/GetWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithdrawServiceServer).GetWithdraw(ctx, req.(*GetWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WithdrawService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "extapi.WithdrawService",
	HandlerType: (*WithdrawServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWithdrawFee",
			Handler:    _WithdrawService_GetWithdrawFee_Handler,
		},
		{
			MethodName: "GetWithdrawHistory",
			Handler:    _WithdrawService_GetWithdrawHistory_Handler,
		},
		{
			MethodName: "ModifyWithdrawFee",
			Handler:    _WithdrawService_ModifyWithdrawFee_Handler,
		},
		{
			MethodName: "GetWithdraw",
			Handler:    _WithdrawService_GetWithdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "withdraw.proto",
}
