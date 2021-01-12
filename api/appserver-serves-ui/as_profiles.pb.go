// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: as_profiles.proto

package appserver_serves_ui

import (
	proto "github.com/golang/protobuf/proto"
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

type RatePolicy int32

const (
	// Drop
	RatePolicy_DROP RatePolicy = 0
	// Mark
	RatePolicy_MARK RatePolicy = 1
)

// Enum value maps for RatePolicy.
var (
	RatePolicy_name = map[int32]string{
		0: "DROP",
		1: "MARK",
	}
	RatePolicy_value = map[string]int32{
		"DROP": 0,
		"MARK": 1,
	}
)

func (x RatePolicy) Enum() *RatePolicy {
	p := new(RatePolicy)
	*p = x
	return p
}

func (x RatePolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RatePolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_as_profiles_proto_enumTypes[0].Descriptor()
}

func (RatePolicy) Type() protoreflect.EnumType {
	return &file_as_profiles_proto_enumTypes[0]
}

func (x RatePolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RatePolicy.Descriptor instead.
func (RatePolicy) EnumDescriptor() ([]byte, []int) {
	return file_as_profiles_proto_rawDescGZIP(), []int{0}
}

type ServiceProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service-profile ID (UUID string).
	// This will be automatically set on create.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Service-profile name.
	Name string `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID to which the service-profile is assigned.
	OrganizationId int64 `protobuf:"varint,22,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID on which the service-profile is provisioned.
	NetworkServerId int64 `protobuf:"varint,23,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Token bucket filling rate, including ACKs (packet/h).
	UlRate uint32 `protobuf:"varint,2,opt,name=ul_rate,json=ulRate,proto3" json:"ul_rate,omitempty"`
	// Token bucket burst size.
	UlBucketSize uint32 `protobuf:"varint,3,opt,name=ul_bucket_size,json=ulBucketSize,proto3" json:"ul_bucket_size,omitempty"`
	// Drop or mark when exceeding ULRate.
	UlRatePolicy RatePolicy `protobuf:"varint,4,opt,name=ul_rate_policy,json=ulRatePolicy,proto3,enum=appserver_serves_ui.RatePolicy" json:"ul_rate_policy,omitempty"`
	// Token bucket filling rate, including ACKs (packet/h).
	DlRate uint32 `protobuf:"varint,5,opt,name=dl_rate,json=dlRate,proto3" json:"dl_rate,omitempty"`
	// Token bucket burst size.
	DlBucketSize uint32 `protobuf:"varint,6,opt,name=dl_bucket_size,json=dlBucketSize,proto3" json:"dl_bucket_size,omitempty"`
	// Drop or mark when exceeding DLRate.
	DlRatePolicy RatePolicy `protobuf:"varint,7,opt,name=dl_rate_policy,json=dlRatePolicy,proto3,enum=appserver_serves_ui.RatePolicy" json:"dl_rate_policy,omitempty"`
	// GW metadata (RSSI, SNR, GW geoloc., etc.) are added to the packet sent to AS.
	AddGwMetadata bool `protobuf:"varint,8,opt,name=add_gw_metadata,json=addGWMetaData,proto3" json:"add_gw_metadata,omitempty"`
	// Frequency to initiate an End-Device status request (request/day).
	DevStatusReqFreq uint32 `protobuf:"varint,9,opt,name=dev_status_req_freq,json=devStatusReqFreq,proto3" json:"dev_status_req_freq,omitempty"`
	// Report End-Device battery level to AS.
	ReportDevStatusBattery bool `protobuf:"varint,10,opt,name=report_dev_status_battery,json=reportDevStatusBattery,proto3" json:"report_dev_status_battery,omitempty"`
	// Report End-Device margin to AS.
	ReportDevStatusMargin bool `protobuf:"varint,11,opt,name=report_dev_status_margin,json=reportDevStatusMargin,proto3" json:"report_dev_status_margin,omitempty"`
	// Minimum allowed data rate. Used for ADR.
	DrMin uint32 `protobuf:"varint,12,opt,name=dr_min,json=drMin,proto3" json:"dr_min,omitempty"`
	// Maximum allowed data rate. Used for ADR.
	DrMax uint32 `protobuf:"varint,13,opt,name=dr_max,json=drMax,proto3" json:"dr_max,omitempty"`
	// Channel mask. sNS does not have to obey (i.e., informative).
	ChannelMask []byte `protobuf:"bytes,14,opt,name=channel_mask,json=channelMask,proto3" json:"channel_mask,omitempty"`
	// Passive Roaming allowed.
	PrAllowed bool `protobuf:"varint,15,opt,name=pr_allowed,json=prAllowed,proto3" json:"pr_allowed,omitempty"`
	// Handover Roaming allowed.
	HrAllowed bool `protobuf:"varint,16,opt,name=hr_allowed,json=hrAllowed,proto3" json:"hr_allowed,omitempty"`
	// Roaming Activation allowed.
	RaAllowed bool `protobuf:"varint,17,opt,name=ra_allowed,json=raAllowed,proto3" json:"ra_allowed,omitempty"`
	// Enable network geolocation service.
	NwkGeoLoc bool `protobuf:"varint,18,opt,name=nwk_geo_loc,json=nwkGeoLoc,proto3" json:"nwk_geo_loc,omitempty"`
	// Target Packet Error Rate.
	TargetPer uint32 `protobuf:"varint,19,opt,name=target_per,json=targetPER,proto3" json:"target_per,omitempty"`
	// Minimum number of receiving GWs (informative).
	MinGwDiversity uint32 `protobuf:"varint,20,opt,name=min_gw_diversity,json=minGWDiversity,proto3" json:"min_gw_diversity,omitempty"`
}

func (x *ServiceProfile) Reset() {
	*x = ServiceProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_profiles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceProfile) ProtoMessage() {}

func (x *ServiceProfile) ProtoReflect() protoreflect.Message {
	mi := &file_as_profiles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceProfile.ProtoReflect.Descriptor instead.
func (*ServiceProfile) Descriptor() ([]byte, []int) {
	return file_as_profiles_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceProfile) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *ServiceProfile) GetNetworkServerId() int64 {
	if x != nil {
		return x.NetworkServerId
	}
	return 0
}

func (x *ServiceProfile) GetUlRate() uint32 {
	if x != nil {
		return x.UlRate
	}
	return 0
}

func (x *ServiceProfile) GetUlBucketSize() uint32 {
	if x != nil {
		return x.UlBucketSize
	}
	return 0
}

func (x *ServiceProfile) GetUlRatePolicy() RatePolicy {
	if x != nil {
		return x.UlRatePolicy
	}
	return RatePolicy_DROP
}

func (x *ServiceProfile) GetDlRate() uint32 {
	if x != nil {
		return x.DlRate
	}
	return 0
}

func (x *ServiceProfile) GetDlBucketSize() uint32 {
	if x != nil {
		return x.DlBucketSize
	}
	return 0
}

func (x *ServiceProfile) GetDlRatePolicy() RatePolicy {
	if x != nil {
		return x.DlRatePolicy
	}
	return RatePolicy_DROP
}

func (x *ServiceProfile) GetAddGwMetadata() bool {
	if x != nil {
		return x.AddGwMetadata
	}
	return false
}

func (x *ServiceProfile) GetDevStatusReqFreq() uint32 {
	if x != nil {
		return x.DevStatusReqFreq
	}
	return 0
}

func (x *ServiceProfile) GetReportDevStatusBattery() bool {
	if x != nil {
		return x.ReportDevStatusBattery
	}
	return false
}

func (x *ServiceProfile) GetReportDevStatusMargin() bool {
	if x != nil {
		return x.ReportDevStatusMargin
	}
	return false
}

func (x *ServiceProfile) GetDrMin() uint32 {
	if x != nil {
		return x.DrMin
	}
	return 0
}

func (x *ServiceProfile) GetDrMax() uint32 {
	if x != nil {
		return x.DrMax
	}
	return 0
}

func (x *ServiceProfile) GetChannelMask() []byte {
	if x != nil {
		return x.ChannelMask
	}
	return nil
}

func (x *ServiceProfile) GetPrAllowed() bool {
	if x != nil {
		return x.PrAllowed
	}
	return false
}

func (x *ServiceProfile) GetHrAllowed() bool {
	if x != nil {
		return x.HrAllowed
	}
	return false
}

func (x *ServiceProfile) GetRaAllowed() bool {
	if x != nil {
		return x.RaAllowed
	}
	return false
}

func (x *ServiceProfile) GetNwkGeoLoc() bool {
	if x != nil {
		return x.NwkGeoLoc
	}
	return false
}

func (x *ServiceProfile) GetTargetPer() uint32 {
	if x != nil {
		return x.TargetPer
	}
	return 0
}

func (x *ServiceProfile) GetMinGwDiversity() uint32 {
	if x != nil {
		return x.MinGwDiversity
	}
	return 0
}

type DeviceProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Device-profile name.
	Name string `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID to which the service-profile is assigned.
	OrganizationId int64 `protobuf:"varint,22,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID on which the service-profile is provisioned.
	NetworkServerId int64 `protobuf:"varint,23,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// End-Device supports Class B.
	SupportsClassB bool `protobuf:"varint,2,opt,name=supports_class_b,json=supportsClassB,proto3" json:"supports_class_b,omitempty"`
	// Maximum delay for the End-Device to answer a MAC request or a confirmed DL frame (mandatory if class B mode supported).
	ClassBTimeout uint32 `protobuf:"varint,3,opt,name=class_b_timeout,json=classBTimeout,proto3" json:"class_b_timeout,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotPeriod uint32 `protobuf:"varint,4,opt,name=ping_slot_period,json=pingSlotPeriod,proto3" json:"ping_slot_period,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotDr uint32 `protobuf:"varint,5,opt,name=ping_slot_dr,json=pingSlotDR,proto3" json:"ping_slot_dr,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotFreq uint32 `protobuf:"varint,6,opt,name=ping_slot_freq,json=pingSlotFreq,proto3" json:"ping_slot_freq,omitempty"`
	// End-Device supports Class C.
	SupportsClassC bool `protobuf:"varint,7,opt,name=supports_class_c,json=supportsClassC,proto3" json:"supports_class_c,omitempty"`
	// Maximum delay for the End-Device to answer a MAC request or a confirmed DL frame (mandatory if class C mode supported).
	ClassCTimeout uint32 `protobuf:"varint,8,opt,name=class_c_timeout,json=classCTimeout,proto3" json:"class_c_timeout,omitempty"`
	// Version of the LoRaWAN supported by the End-Device.
	MacVersion string `protobuf:"bytes,9,opt,name=mac_version,json=macVersion,proto3" json:"mac_version,omitempty"`
	// Revision of the Regional Parameters document supported by the End-Device.
	RegParamsRevision string `protobuf:"bytes,10,opt,name=reg_params_revision,json=regParamsRevision,proto3" json:"reg_params_revision,omitempty"`
	// Class A RX1 delay (mandatory for ABP).
	RxDelay_1 uint32 `protobuf:"varint,11,opt,name=rx_delay_1,json=rxDelay1,proto3" json:"rx_delay_1,omitempty"`
	// RX1 data rate offset (mandatory for ABP).
	RxDrOffset_1 uint32 `protobuf:"varint,12,opt,name=rx_dr_offset_1,json=rxDROffset1,proto3" json:"rx_dr_offset_1,omitempty"`
	// RX2 data rate (mandatory for ABP).
	RxDatarate_2 uint32 `protobuf:"varint,13,opt,name=rx_datarate_2,json=rxDataRate2,proto3" json:"rx_datarate_2,omitempty"`
	// RX2 channel frequency (mandatory for ABP).
	RxFreq_2 uint32 `protobuf:"varint,14,opt,name=rx_freq_2,json=rxFreq2,proto3" json:"rx_freq_2,omitempty"`
	// List of factory-preset frequencies (mandatory for ABP).
	FactoryPresetFreqs []uint32 `protobuf:"varint,15,rep,packed,name=factory_preset_freqs,json=factoryPresetFreqs,proto3" json:"factory_preset_freqs,omitempty"`
	// Maximum EIRP supported by the End-Device.
	MaxEirp uint32 `protobuf:"varint,16,opt,name=max_eirp,json=maxEIRP,proto3" json:"max_eirp,omitempty"`
	// Maximum duty cycle supported by the End-Device.
	MaxDutyCycle uint32 `protobuf:"varint,17,opt,name=max_duty_cycle,json=maxDutyCycle,proto3" json:"max_duty_cycle,omitempty"`
	// End-Device supports Join (OTAA) or not (ABP).
	SupportsJoin bool `protobuf:"varint,18,opt,name=supports_join,json=supportsJoin,proto3" json:"supports_join,omitempty"`
	// RF region name.
	RfRegion string `protobuf:"bytes,19,opt,name=rf_region,json=rfRegion,proto3" json:"rf_region,omitempty"`
	// End-Device uses 32bit FCnt (mandatory for LoRaWAN 1.0 End-Device).
	Supports_32BitFCnt bool `protobuf:"varint,20,opt,name=supports_32bit_f_cnt,json=supports32BitFCnt,proto3" json:"supports_32bit_f_cnt,omitempty"`
	// Payload codec.
	// Leave blank to disable the codec feature.
	PayloadCodec string `protobuf:"bytes,24,opt,name=payload_codec,json=payloadCodec,proto3" json:"payload_codec,omitempty"`
	// Payload encoder script.
	// Depending the codec, it is possible to provide a script which implements
	// the encoder function.
	PayloadEncoderScript string `protobuf:"bytes,25,opt,name=payload_encoder_script,json=payloadEncoderScript,proto3" json:"payload_encoder_script,omitempty"`
	// Payload decoder script.
	// Depending the codec, it is possible to provide a script which implements
	// the decoder function.
	PayloadDecoderScript string `protobuf:"bytes,26,opt,name=payload_decoder_script,json=payloadDecoderScript,proto3" json:"payload_decoder_script,omitempty"`
	// Geolocation buffer TTL (in seconds).
	// When > 0, uplink RX meta-data will be stored in a buffer so that
	// the meta-data of multiple uplinks can be used for geolocation.
	GeolocBufferTtl uint32 `protobuf:"varint,27,opt,name=geoloc_buffer_ttl,json=geolocBufferTTL,proto3" json:"geoloc_buffer_ttl,omitempty"`
	// Geolocation minimum buffer size.
	// When > 0, geolocation will only be performed when the buffer has
	// at least the given size.
	GeolocMinBufferSize uint32 `protobuf:"varint,28,opt,name=geoloc_min_buffer_size,json=geolocMinBufferSize,proto3" json:"geoloc_min_buffer_size,omitempty"`
	// User defined tags.
	Tags map[string]string `protobuf:"bytes,29,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeviceProfile) Reset() {
	*x = DeviceProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_profiles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceProfile) ProtoMessage() {}

func (x *DeviceProfile) ProtoReflect() protoreflect.Message {
	mi := &file_as_profiles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceProfile.ProtoReflect.Descriptor instead.
func (*DeviceProfile) Descriptor() ([]byte, []int) {
	return file_as_profiles_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceProfile) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *DeviceProfile) GetNetworkServerId() int64 {
	if x != nil {
		return x.NetworkServerId
	}
	return 0
}

func (x *DeviceProfile) GetSupportsClassB() bool {
	if x != nil {
		return x.SupportsClassB
	}
	return false
}

func (x *DeviceProfile) GetClassBTimeout() uint32 {
	if x != nil {
		return x.ClassBTimeout
	}
	return 0
}

func (x *DeviceProfile) GetPingSlotPeriod() uint32 {
	if x != nil {
		return x.PingSlotPeriod
	}
	return 0
}

func (x *DeviceProfile) GetPingSlotDr() uint32 {
	if x != nil {
		return x.PingSlotDr
	}
	return 0
}

func (x *DeviceProfile) GetPingSlotFreq() uint32 {
	if x != nil {
		return x.PingSlotFreq
	}
	return 0
}

func (x *DeviceProfile) GetSupportsClassC() bool {
	if x != nil {
		return x.SupportsClassC
	}
	return false
}

func (x *DeviceProfile) GetClassCTimeout() uint32 {
	if x != nil {
		return x.ClassCTimeout
	}
	return 0
}

func (x *DeviceProfile) GetMacVersion() string {
	if x != nil {
		return x.MacVersion
	}
	return ""
}

func (x *DeviceProfile) GetRegParamsRevision() string {
	if x != nil {
		return x.RegParamsRevision
	}
	return ""
}

func (x *DeviceProfile) GetRxDelay_1() uint32 {
	if x != nil {
		return x.RxDelay_1
	}
	return 0
}

func (x *DeviceProfile) GetRxDrOffset_1() uint32 {
	if x != nil {
		return x.RxDrOffset_1
	}
	return 0
}

func (x *DeviceProfile) GetRxDatarate_2() uint32 {
	if x != nil {
		return x.RxDatarate_2
	}
	return 0
}

func (x *DeviceProfile) GetRxFreq_2() uint32 {
	if x != nil {
		return x.RxFreq_2
	}
	return 0
}

func (x *DeviceProfile) GetFactoryPresetFreqs() []uint32 {
	if x != nil {
		return x.FactoryPresetFreqs
	}
	return nil
}

func (x *DeviceProfile) GetMaxEirp() uint32 {
	if x != nil {
		return x.MaxEirp
	}
	return 0
}

func (x *DeviceProfile) GetMaxDutyCycle() uint32 {
	if x != nil {
		return x.MaxDutyCycle
	}
	return 0
}

func (x *DeviceProfile) GetSupportsJoin() bool {
	if x != nil {
		return x.SupportsJoin
	}
	return false
}

func (x *DeviceProfile) GetRfRegion() string {
	if x != nil {
		return x.RfRegion
	}
	return ""
}

func (x *DeviceProfile) GetSupports_32BitFCnt() bool {
	if x != nil {
		return x.Supports_32BitFCnt
	}
	return false
}

func (x *DeviceProfile) GetPayloadCodec() string {
	if x != nil {
		return x.PayloadCodec
	}
	return ""
}

func (x *DeviceProfile) GetPayloadEncoderScript() string {
	if x != nil {
		return x.PayloadEncoderScript
	}
	return ""
}

func (x *DeviceProfile) GetPayloadDecoderScript() string {
	if x != nil {
		return x.PayloadDecoderScript
	}
	return ""
}

func (x *DeviceProfile) GetGeolocBufferTtl() uint32 {
	if x != nil {
		return x.GeolocBufferTtl
	}
	return 0
}

func (x *DeviceProfile) GetGeolocMinBufferSize() uint32 {
	if x != nil {
		return x.GeolocMinBufferSize
	}
	return 0
}

func (x *DeviceProfile) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_as_profiles_proto protoreflect.FileDescriptor

var file_as_profiles_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x22, 0xf7, 0x06, 0x0a, 0x0e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x75, 0x6c, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x75, 0x6c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x75, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x75,
	0x69, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0c, 0x75, 0x6c,
	0x52, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x6c,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x64, 0x6c, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x6c, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x6c, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x64, 0x6c, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x0c, 0x64, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x26, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x77, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x47, 0x57,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x13, 0x64, 0x65, 0x76, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x64, 0x65, 0x76, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x46, 0x72, 0x65, 0x71, 0x12, 0x39, 0x0a, 0x19, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x62, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x65, 0x76, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x79, 0x12, 0x37, 0x0a, 0x18, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x76,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x76, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x64,
	0x72, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x64, 0x72, 0x4d,
	0x69, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x72, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x64, 0x72, 0x4d, 0x61, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x70, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x68,
	0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x68, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x61,
	0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x72, 0x61, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x77, 0x6b,
	0x5f, 0x67, 0x65, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x6e, 0x77, 0x6b, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x45, 0x52, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x5f,
	0x67, 0x77, 0x5f, 0x64, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x47, 0x57, 0x44, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x74, 0x79, 0x22, 0xc5, 0x09, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a,
	0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x62, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x28, 0x0a, 0x10, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x53,
	0x6c, 0x6f, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x52, 0x12, 0x24, 0x0a, 0x0e, 0x70,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x6f, 0x74, 0x46, 0x72, 0x65,
	0x71, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x72, 0x65, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x72, 0x78, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x5f, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x78, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x31, 0x12, 0x23, 0x0a, 0x0e, 0x72, 0x78, 0x5f, 0x64, 0x72, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x5f, 0x31, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x78, 0x44, 0x52,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x31, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x78, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x72, 0x78, 0x44, 0x61, 0x74, 0x61, 0x52, 0x61, 0x74, 0x65, 0x32, 0x12, 0x1a, 0x0a, 0x09, 0x72,
	0x78, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x5f, 0x32, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x78, 0x46, 0x72, 0x65, 0x71, 0x32, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x46, 0x72, 0x65, 0x71, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78,
	0x5f, 0x65, 0x69, 0x72, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x78,
	0x45, 0x49, 0x52, 0x50, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x75, 0x74, 0x79,
	0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x44, 0x75, 0x74, 0x79, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x66, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x66, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x14,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x33, 0x32, 0x62, 0x69, 0x74, 0x5f, 0x66,
	0x5f, 0x63, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x33, 0x32, 0x42, 0x69, 0x74, 0x46, 0x43, 0x6e, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x63, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x67, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f,
	0x74, 0x74, 0x6c, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x67, 0x65, 0x6f, 0x6c, 0x6f,
	0x63, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x54, 0x54, 0x4c, 0x12, 0x33, 0x0a, 0x16, 0x67, 0x65,
	0x6f, 0x6c, 0x6f, 0x63, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x67, 0x65, 0x6f, 0x6c,
	0x6f, 0x63, 0x4d, 0x69, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x40, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73,
	0x5f, 0x75, 0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x20, 0x0a, 0x0a, 0x52, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x4f, 0x50,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x01, 0x42, 0x58, 0x5a, 0x56,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x63, 0x2d, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x70, 0x77, 0x61, 0x6e, 0x2d,
	0x61, 0x70, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x2d,
	0x75, 0x69, 0x3b, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_as_profiles_proto_rawDescOnce sync.Once
	file_as_profiles_proto_rawDescData = file_as_profiles_proto_rawDesc
)

func file_as_profiles_proto_rawDescGZIP() []byte {
	file_as_profiles_proto_rawDescOnce.Do(func() {
		file_as_profiles_proto_rawDescData = protoimpl.X.CompressGZIP(file_as_profiles_proto_rawDescData)
	})
	return file_as_profiles_proto_rawDescData
}

var file_as_profiles_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_as_profiles_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_as_profiles_proto_goTypes = []interface{}{
	(RatePolicy)(0),        // 0: appserver_serves_ui.RatePolicy
	(*ServiceProfile)(nil), // 1: appserver_serves_ui.ServiceProfile
	(*DeviceProfile)(nil),  // 2: appserver_serves_ui.DeviceProfile
	nil,                    // 3: appserver_serves_ui.DeviceProfile.TagsEntry
}
var file_as_profiles_proto_depIdxs = []int32{
	0, // 0: appserver_serves_ui.ServiceProfile.ul_rate_policy:type_name -> appserver_serves_ui.RatePolicy
	0, // 1: appserver_serves_ui.ServiceProfile.dl_rate_policy:type_name -> appserver_serves_ui.RatePolicy
	3, // 2: appserver_serves_ui.DeviceProfile.tags:type_name -> appserver_serves_ui.DeviceProfile.TagsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_as_profiles_proto_init() }
func file_as_profiles_proto_init() {
	if File_as_profiles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_as_profiles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceProfile); i {
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
		file_as_profiles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceProfile); i {
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
			RawDescriptor: file_as_profiles_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_as_profiles_proto_goTypes,
		DependencyIndexes: file_as_profiles_proto_depIdxs,
		EnumInfos:         file_as_profiles_proto_enumTypes,
		MessageInfos:      file_as_profiles_proto_msgTypes,
	}.Build()
	File_as_profiles_proto = out.File
	file_as_profiles_proto_rawDesc = nil
	file_as_profiles_proto_goTypes = nil
	file_as_profiles_proto_depIdxs = nil
}
