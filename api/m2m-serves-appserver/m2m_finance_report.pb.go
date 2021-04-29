// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: m2m_finance_report.proto

package m2m_serves_appserver

import (
	context "context"
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

type GetSupportedFiatCurrencyListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSupportedFiatCurrencyListRequest) Reset() {
	*x = GetSupportedFiatCurrencyListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_finance_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedFiatCurrencyListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedFiatCurrencyListRequest) ProtoMessage() {}

func (x *GetSupportedFiatCurrencyListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_finance_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedFiatCurrencyListRequest.ProtoReflect.Descriptor instead.
func (*GetSupportedFiatCurrencyListRequest) Descriptor() ([]byte, []int) {
	return file_m2m_finance_report_proto_rawDescGZIP(), []int{0}
}

type FiatCurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *FiatCurrency) Reset() {
	*x = FiatCurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_finance_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiatCurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiatCurrency) ProtoMessage() {}

func (x *FiatCurrency) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_finance_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiatCurrency.ProtoReflect.Descriptor instead.
func (*FiatCurrency) Descriptor() ([]byte, []int) {
	return file_m2m_finance_report_proto_rawDescGZIP(), []int{1}
}

func (x *FiatCurrency) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FiatCurrency) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetSupportedFiatCurrencyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FiatCurrency []*FiatCurrency `protobuf:"bytes,1,rep,name=fiat_currency,json=fiatCurrency,proto3" json:"fiat_currency,omitempty"`
}

func (x *GetSupportedFiatCurrencyListResponse) Reset() {
	*x = GetSupportedFiatCurrencyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_finance_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedFiatCurrencyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedFiatCurrencyListResponse) ProtoMessage() {}

func (x *GetSupportedFiatCurrencyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_finance_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedFiatCurrencyListResponse.ProtoReflect.Descriptor instead.
func (*GetSupportedFiatCurrencyListResponse) Descriptor() ([]byte, []int) {
	return file_m2m_finance_report_proto_rawDescGZIP(), []int{2}
}

func (x *GetSupportedFiatCurrencyListResponse) GetFiatCurrency() []*FiatCurrency {
	if x != nil {
		return x.FiatCurrency
	}
	return nil
}

type GetSupportedCryptoCurrencyListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSupportedCryptoCurrencyListRequest) Reset() {
	*x = GetSupportedCryptoCurrencyListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_finance_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedCryptoCurrencyListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedCryptoCurrencyListRequest) ProtoMessage() {}

func (x *GetSupportedCryptoCurrencyListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_finance_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedCryptoCurrencyListRequest.ProtoReflect.Descriptor instead.
func (*GetSupportedCryptoCurrencyListRequest) Descriptor() ([]byte, []int) {
	return file_m2m_finance_report_proto_rawDescGZIP(), []int{3}
}

type GetSupportedCryptoCurrencyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CryptoCurrencyList []string `protobuf:"bytes,1,rep,name=crypto_currency_list,json=cryptoCurrencyList,proto3" json:"crypto_currency_list,omitempty"`
}

func (x *GetSupportedCryptoCurrencyListResponse) Reset() {
	*x = GetSupportedCryptoCurrencyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_finance_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedCryptoCurrencyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedCryptoCurrencyListResponse) ProtoMessage() {}

func (x *GetSupportedCryptoCurrencyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_finance_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedCryptoCurrencyListResponse.ProtoReflect.Descriptor instead.
func (*GetSupportedCryptoCurrencyListResponse) Descriptor() ([]byte, []int) {
	return file_m2m_finance_report_proto_rawDescGZIP(), []int{4}
}

func (x *GetSupportedCryptoCurrencyListResponse) GetCryptoCurrencyList() []string {
	if x != nil {
		return x.CryptoCurrencyList
	}
	return nil
}

type GetMXCMiningReportByDateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId int64                `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Start          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	FiatCurrency   string               `protobuf:"bytes,4,opt,name=fiat_currency,json=fiatCurrency,proto3" json:"fiat_currency,omitempty"`
	Decimals       int32                `protobuf:"varint,5,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (x *GetMXCMiningReportByDateRequest) Reset() {
	*x = GetMXCMiningReportByDateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_finance_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMXCMiningReportByDateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMXCMiningReportByDateRequest) ProtoMessage() {}

func (x *GetMXCMiningReportByDateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_finance_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMXCMiningReportByDateRequest.ProtoReflect.Descriptor instead.
func (*GetMXCMiningReportByDateRequest) Descriptor() ([]byte, []int) {
	return file_m2m_finance_report_proto_rawDescGZIP(), []int{5}
}

func (x *GetMXCMiningReportByDateRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *GetMXCMiningReportByDateRequest) GetStart() *timestamp.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetMXCMiningReportByDateRequest) GetEnd() *timestamp.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *GetMXCMiningReportByDateRequest) GetFiatCurrency() string {
	if x != nil {
		return x.FiatCurrency
	}
	return ""
}

func (x *GetMXCMiningReportByDateRequest) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

type MiningRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateTime           *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`
	MXCMined           string               `protobuf:"bytes,2,opt,name=MXC_mined,json=MXCMined,proto3" json:"MXC_mined,omitempty"`
	MXCSettlementPrice string               `protobuf:"bytes,3,opt,name=MXC_settlement_price,json=MXCSettlementPrice,proto3" json:"MXC_settlement_price,omitempty"`
	// fiat currency here is the same at passed as param in GetMXCMiningReportByDateRequest
	FiatCurrencyMined string `protobuf:"bytes,4,opt,name=fiat_currency_mined,json=fiatCurrencyMined,proto3" json:"fiat_currency_mined,omitempty"`
	OnlineSeconds     int64  `protobuf:"varint,5,opt,name=online_seconds,json=onlineSeconds,proto3" json:"online_seconds,omitempty"`
}

func (x *MiningRecord) Reset() {
	*x = MiningRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_finance_report_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiningRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiningRecord) ProtoMessage() {}

func (x *MiningRecord) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_finance_report_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiningRecord.ProtoReflect.Descriptor instead.
func (*MiningRecord) Descriptor() ([]byte, []int) {
	return file_m2m_finance_report_proto_rawDescGZIP(), []int{6}
}

func (x *MiningRecord) GetDateTime() *timestamp.Timestamp {
	if x != nil {
		return x.DateTime
	}
	return nil
}

func (x *MiningRecord) GetMXCMined() string {
	if x != nil {
		return x.MXCMined
	}
	return ""
}

func (x *MiningRecord) GetMXCSettlementPrice() string {
	if x != nil {
		return x.MXCSettlementPrice
	}
	return ""
}

func (x *MiningRecord) GetFiatCurrencyMined() string {
	if x != nil {
		return x.FiatCurrencyMined
	}
	return ""
}

func (x *MiningRecord) GetOnlineSeconds() int64 {
	if x != nil {
		return x.OnlineSeconds
	}
	return 0
}

type GetMXCMiningReportByDateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiningRecordList []*MiningRecord `protobuf:"bytes,1,rep,name=mining_record_list,json=miningRecordList,proto3" json:"mining_record_list,omitempty"`
}

func (x *GetMXCMiningReportByDateResponse) Reset() {
	*x = GetMXCMiningReportByDateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m2m_finance_report_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMXCMiningReportByDateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMXCMiningReportByDateResponse) ProtoMessage() {}

func (x *GetMXCMiningReportByDateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_m2m_finance_report_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMXCMiningReportByDateResponse.ProtoReflect.Descriptor instead.
func (*GetMXCMiningReportByDateResponse) Descriptor() ([]byte, []int) {
	return file_m2m_finance_report_proto_rawDescGZIP(), []int{7}
}

func (x *GetMXCMiningReportByDateResponse) GetMiningRecordList() []*MiningRecord {
	if x != nil {
		return x.MiningRecordList
	}
	return nil
}

var File_m2m_finance_report_proto protoreflect.FileDescriptor

var file_m2m_finance_report_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x32, 0x6d, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x32, 0x6d, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x46, 0x69, 0x61, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6f, 0x0a, 0x24, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x66, 0x69, 0x61, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x32, 0x6d, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0c, 0x66,
	0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x27, 0x0a, 0x25, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x5a, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x14, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0xeb, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4d, 0x58, 0x43, 0x4d, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x69, 0x61, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x22, 0xed,
	0x01, 0x0a, 0x0c, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x37, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x4d, 0x58, 0x43, 0x5f,
	0x6d, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x58, 0x43,
	0x4d, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x4d, 0x58, 0x43, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x4d, 0x58, 0x43, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x61, 0x74, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4d, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x74,
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x4d, 0x58, 0x43, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x50, 0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6d, 0x32, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x32, 0xd8, 0x03, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x58, 0x43, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x6d, 0x32, 0x6d,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x58, 0x43, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x36, 0x2e, 0x6d, 0x32, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x61,
	0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x58, 0x43, 0x4d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x95, 0x01, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x2e, 0x6d, 0x32, 0x6d,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x46, 0x69,
	0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6d, 0x32, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x9b, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x2e, 0x6d, 0x32, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x73, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3c, 0x2e, 0x6d, 0x32, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x61,
	0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78,
	0x63, 0x2d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x78, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x32, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_m2m_finance_report_proto_rawDescOnce sync.Once
	file_m2m_finance_report_proto_rawDescData = file_m2m_finance_report_proto_rawDesc
)

func file_m2m_finance_report_proto_rawDescGZIP() []byte {
	file_m2m_finance_report_proto_rawDescOnce.Do(func() {
		file_m2m_finance_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_m2m_finance_report_proto_rawDescData)
	})
	return file_m2m_finance_report_proto_rawDescData
}

var file_m2m_finance_report_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_m2m_finance_report_proto_goTypes = []interface{}{
	(*GetSupportedFiatCurrencyListRequest)(nil),    // 0: m2m_serves_appserver.GetSupportedFiatCurrencyListRequest
	(*FiatCurrency)(nil),                           // 1: m2m_serves_appserver.FiatCurrency
	(*GetSupportedFiatCurrencyListResponse)(nil),   // 2: m2m_serves_appserver.GetSupportedFiatCurrencyListResponse
	(*GetSupportedCryptoCurrencyListRequest)(nil),  // 3: m2m_serves_appserver.GetSupportedCryptoCurrencyListRequest
	(*GetSupportedCryptoCurrencyListResponse)(nil), // 4: m2m_serves_appserver.GetSupportedCryptoCurrencyListResponse
	(*GetMXCMiningReportByDateRequest)(nil),        // 5: m2m_serves_appserver.GetMXCMiningReportByDateRequest
	(*MiningRecord)(nil),                           // 6: m2m_serves_appserver.MiningRecord
	(*GetMXCMiningReportByDateResponse)(nil),       // 7: m2m_serves_appserver.GetMXCMiningReportByDateResponse
	(*timestamp.Timestamp)(nil),                    // 8: google.protobuf.Timestamp
}
var file_m2m_finance_report_proto_depIdxs = []int32{
	1, // 0: m2m_serves_appserver.GetSupportedFiatCurrencyListResponse.fiat_currency:type_name -> m2m_serves_appserver.FiatCurrency
	8, // 1: m2m_serves_appserver.GetMXCMiningReportByDateRequest.start:type_name -> google.protobuf.Timestamp
	8, // 2: m2m_serves_appserver.GetMXCMiningReportByDateRequest.end:type_name -> google.protobuf.Timestamp
	8, // 3: m2m_serves_appserver.MiningRecord.date_time:type_name -> google.protobuf.Timestamp
	6, // 4: m2m_serves_appserver.GetMXCMiningReportByDateResponse.mining_record_list:type_name -> m2m_serves_appserver.MiningRecord
	5, // 5: m2m_serves_appserver.FinanceReportService.GetMXCMiningReportByDate:input_type -> m2m_serves_appserver.GetMXCMiningReportByDateRequest
	0, // 6: m2m_serves_appserver.FinanceReportService.GetSupportedFiatCurrencyList:input_type -> m2m_serves_appserver.GetSupportedFiatCurrencyListRequest
	3, // 7: m2m_serves_appserver.FinanceReportService.GetSupportedCryptoCurrencyList:input_type -> m2m_serves_appserver.GetSupportedCryptoCurrencyListRequest
	7, // 8: m2m_serves_appserver.FinanceReportService.GetMXCMiningReportByDate:output_type -> m2m_serves_appserver.GetMXCMiningReportByDateResponse
	2, // 9: m2m_serves_appserver.FinanceReportService.GetSupportedFiatCurrencyList:output_type -> m2m_serves_appserver.GetSupportedFiatCurrencyListResponse
	4, // 10: m2m_serves_appserver.FinanceReportService.GetSupportedCryptoCurrencyList:output_type -> m2m_serves_appserver.GetSupportedCryptoCurrencyListResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_m2m_finance_report_proto_init() }
func file_m2m_finance_report_proto_init() {
	if File_m2m_finance_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_m2m_finance_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupportedFiatCurrencyListRequest); i {
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
		file_m2m_finance_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiatCurrency); i {
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
		file_m2m_finance_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupportedFiatCurrencyListResponse); i {
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
		file_m2m_finance_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupportedCryptoCurrencyListRequest); i {
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
		file_m2m_finance_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupportedCryptoCurrencyListResponse); i {
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
		file_m2m_finance_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMXCMiningReportByDateRequest); i {
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
		file_m2m_finance_report_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiningRecord); i {
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
		file_m2m_finance_report_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMXCMiningReportByDateResponse); i {
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
			RawDescriptor: file_m2m_finance_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_m2m_finance_report_proto_goTypes,
		DependencyIndexes: file_m2m_finance_report_proto_depIdxs,
		MessageInfos:      file_m2m_finance_report_proto_msgTypes,
	}.Build()
	File_m2m_finance_report_proto = out.File
	file_m2m_finance_report_proto_rawDesc = nil
	file_m2m_finance_report_proto_goTypes = nil
	file_m2m_finance_report_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FinanceReportServiceClient is the client API for FinanceReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinanceReportServiceClient interface {
	// GetMXCMiningReportByDate returns list of mining records filtered by start and end date
	GetMXCMiningReportByDate(ctx context.Context, in *GetMXCMiningReportByDateRequest, opts ...grpc.CallOption) (*GetMXCMiningReportByDateResponse, error)
	// GetSupportedFiatCurrencyList returns fiat currency list that is supported by mxp service and verified with coingecko server
	GetSupportedFiatCurrencyList(ctx context.Context, in *GetSupportedFiatCurrencyListRequest, opts ...grpc.CallOption) (*GetSupportedFiatCurrencyListResponse, error)
	GetSupportedCryptoCurrencyList(ctx context.Context, in *GetSupportedCryptoCurrencyListRequest, opts ...grpc.CallOption) (*GetSupportedCryptoCurrencyListResponse, error)
}

type financeReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinanceReportServiceClient(cc grpc.ClientConnInterface) FinanceReportServiceClient {
	return &financeReportServiceClient{cc}
}

func (c *financeReportServiceClient) GetMXCMiningReportByDate(ctx context.Context, in *GetMXCMiningReportByDateRequest, opts ...grpc.CallOption) (*GetMXCMiningReportByDateResponse, error) {
	out := new(GetMXCMiningReportByDateResponse)
	err := c.cc.Invoke(ctx, "/m2m_serves_appserver.FinanceReportService/GetMXCMiningReportByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeReportServiceClient) GetSupportedFiatCurrencyList(ctx context.Context, in *GetSupportedFiatCurrencyListRequest, opts ...grpc.CallOption) (*GetSupportedFiatCurrencyListResponse, error) {
	out := new(GetSupportedFiatCurrencyListResponse)
	err := c.cc.Invoke(ctx, "/m2m_serves_appserver.FinanceReportService/GetSupportedFiatCurrencyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeReportServiceClient) GetSupportedCryptoCurrencyList(ctx context.Context, in *GetSupportedCryptoCurrencyListRequest, opts ...grpc.CallOption) (*GetSupportedCryptoCurrencyListResponse, error) {
	out := new(GetSupportedCryptoCurrencyListResponse)
	err := c.cc.Invoke(ctx, "/m2m_serves_appserver.FinanceReportService/GetSupportedCryptoCurrencyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceReportServiceServer is the server API for FinanceReportService service.
type FinanceReportServiceServer interface {
	// GetMXCMiningReportByDate returns list of mining records filtered by start and end date
	GetMXCMiningReportByDate(context.Context, *GetMXCMiningReportByDateRequest) (*GetMXCMiningReportByDateResponse, error)
	// GetSupportedFiatCurrencyList returns fiat currency list that is supported by mxp service and verified with coingecko server
	GetSupportedFiatCurrencyList(context.Context, *GetSupportedFiatCurrencyListRequest) (*GetSupportedFiatCurrencyListResponse, error)
	GetSupportedCryptoCurrencyList(context.Context, *GetSupportedCryptoCurrencyListRequest) (*GetSupportedCryptoCurrencyListResponse, error)
}

// UnimplementedFinanceReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFinanceReportServiceServer struct {
}

func (*UnimplementedFinanceReportServiceServer) GetMXCMiningReportByDate(context.Context, *GetMXCMiningReportByDateRequest) (*GetMXCMiningReportByDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMXCMiningReportByDate not implemented")
}
func (*UnimplementedFinanceReportServiceServer) GetSupportedFiatCurrencyList(context.Context, *GetSupportedFiatCurrencyListRequest) (*GetSupportedFiatCurrencyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportedFiatCurrencyList not implemented")
}
func (*UnimplementedFinanceReportServiceServer) GetSupportedCryptoCurrencyList(context.Context, *GetSupportedCryptoCurrencyListRequest) (*GetSupportedCryptoCurrencyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportedCryptoCurrencyList not implemented")
}

func RegisterFinanceReportServiceServer(s *grpc.Server, srv FinanceReportServiceServer) {
	s.RegisterService(&_FinanceReportService_serviceDesc, srv)
}

func _FinanceReportService_GetMXCMiningReportByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMXCMiningReportByDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceReportServiceServer).GetMXCMiningReportByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m2m_serves_appserver.FinanceReportService/GetMXCMiningReportByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceReportServiceServer).GetMXCMiningReportByDate(ctx, req.(*GetMXCMiningReportByDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceReportService_GetSupportedFiatCurrencyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupportedFiatCurrencyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceReportServiceServer).GetSupportedFiatCurrencyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m2m_serves_appserver.FinanceReportService/GetSupportedFiatCurrencyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceReportServiceServer).GetSupportedFiatCurrencyList(ctx, req.(*GetSupportedFiatCurrencyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceReportService_GetSupportedCryptoCurrencyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupportedCryptoCurrencyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceReportServiceServer).GetSupportedCryptoCurrencyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m2m_serves_appserver.FinanceReportService/GetSupportedCryptoCurrencyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceReportServiceServer).GetSupportedCryptoCurrencyList(ctx, req.(*GetSupportedCryptoCurrencyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FinanceReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "m2m_serves_appserver.FinanceReportService",
	HandlerType: (*FinanceReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMXCMiningReportByDate",
			Handler:    _FinanceReportService_GetMXCMiningReportByDate_Handler,
		},
		{
			MethodName: "GetSupportedFiatCurrencyList",
			Handler:    _FinanceReportService_GetSupportedFiatCurrencyList_Handler,
		},
		{
			MethodName: "GetSupportedCryptoCurrencyList",
			Handler:    _FinanceReportService_GetSupportedCryptoCurrencyList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "m2m_finance_report.proto",
}
