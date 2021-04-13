// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: deviceQueue.proto

package appserver_serves_ui

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type DeviceQueueItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Set this to true when an acknowledgement from the device is required.
	// Please note that this must not be used to guarantee a delivery.
	Confirmed bool `protobuf:"varint,2,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	// Downlink frame-counter.
	// This will be automatically set on enquue.
	FCnt uint32 `protobuf:"varint,6,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// FPort used (must be > 0)
	FPort uint32 `protobuf:"varint,3,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// Base64 encoded data.
	// Or use the json_object field when an application codec has been configured.
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// JSON object (string).
	// Only use this when an application codec has been configured that can convert
	// this object into binary form.
	JsonObject string `protobuf:"bytes,5,opt,name=json_object,json=jsonObject,proto3" json:"json_object,omitempty"`
}

func (x *DeviceQueueItem) Reset() {
	*x = DeviceQueueItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceQueue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceQueueItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceQueueItem) ProtoMessage() {}

func (x *DeviceQueueItem) ProtoReflect() protoreflect.Message {
	mi := &file_deviceQueue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceQueueItem.ProtoReflect.Descriptor instead.
func (*DeviceQueueItem) Descriptor() ([]byte, []int) {
	return file_deviceQueue_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceQueueItem) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

func (x *DeviceQueueItem) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

func (x *DeviceQueueItem) GetFCnt() uint32 {
	if x != nil {
		return x.FCnt
	}
	return 0
}

func (x *DeviceQueueItem) GetFPort() uint32 {
	if x != nil {
		return x.FPort
	}
	return 0
}

func (x *DeviceQueueItem) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DeviceQueueItem) GetJsonObject() string {
	if x != nil {
		return x.JsonObject
	}
	return ""
}

type EnqueueDeviceQueueItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Queue-item object to enqueue.
	DeviceQueueItem *DeviceQueueItem `protobuf:"bytes,1,opt,name=device_queue_item,json=deviceQueueItem,proto3" json:"device_queue_item,omitempty"`
}

func (x *EnqueueDeviceQueueItemRequest) Reset() {
	*x = EnqueueDeviceQueueItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceQueue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueDeviceQueueItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueDeviceQueueItemRequest) ProtoMessage() {}

func (x *EnqueueDeviceQueueItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deviceQueue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueDeviceQueueItemRequest.ProtoReflect.Descriptor instead.
func (*EnqueueDeviceQueueItemRequest) Descriptor() ([]byte, []int) {
	return file_deviceQueue_proto_rawDescGZIP(), []int{1}
}

func (x *EnqueueDeviceQueueItemRequest) GetDeviceQueueItem() *DeviceQueueItem {
	if x != nil {
		return x.DeviceQueueItem
	}
	return nil
}

type EnqueueDeviceQueueItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Frame-counter for the enqueued payload.
	FCnt uint32 `protobuf:"varint,1,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
}

func (x *EnqueueDeviceQueueItemResponse) Reset() {
	*x = EnqueueDeviceQueueItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceQueue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueDeviceQueueItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueDeviceQueueItemResponse) ProtoMessage() {}

func (x *EnqueueDeviceQueueItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deviceQueue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueDeviceQueueItemResponse.ProtoReflect.Descriptor instead.
func (*EnqueueDeviceQueueItemResponse) Descriptor() ([]byte, []int) {
	return file_deviceQueue_proto_rawDescGZIP(), []int{2}
}

func (x *EnqueueDeviceQueueItemResponse) GetFCnt() uint32 {
	if x != nil {
		return x.FCnt
	}
	return 0
}

type FlushDeviceQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
}

func (x *FlushDeviceQueueRequest) Reset() {
	*x = FlushDeviceQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceQueue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushDeviceQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushDeviceQueueRequest) ProtoMessage() {}

func (x *FlushDeviceQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deviceQueue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushDeviceQueueRequest.ProtoReflect.Descriptor instead.
func (*FlushDeviceQueueRequest) Descriptor() ([]byte, []int) {
	return file_deviceQueue_proto_rawDescGZIP(), []int{3}
}

func (x *FlushDeviceQueueRequest) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

type ListDeviceQueueItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
}

func (x *ListDeviceQueueItemsRequest) Reset() {
	*x = ListDeviceQueueItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceQueue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceQueueItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceQueueItemsRequest) ProtoMessage() {}

func (x *ListDeviceQueueItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deviceQueue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceQueueItemsRequest.ProtoReflect.Descriptor instead.
func (*ListDeviceQueueItemsRequest) Descriptor() ([]byte, []int) {
	return file_deviceQueue_proto_rawDescGZIP(), []int{4}
}

func (x *ListDeviceQueueItemsRequest) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

type ListDeviceQueueItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceQueueItems []*DeviceQueueItem `protobuf:"bytes,1,rep,name=device_queue_items,json=deviceQueueItems,proto3" json:"device_queue_items,omitempty"`
}

func (x *ListDeviceQueueItemsResponse) Reset() {
	*x = ListDeviceQueueItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deviceQueue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceQueueItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceQueueItemsResponse) ProtoMessage() {}

func (x *ListDeviceQueueItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deviceQueue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceQueueItemsResponse.ProtoReflect.Descriptor instead.
func (*ListDeviceQueueItemsResponse) Descriptor() ([]byte, []int) {
	return file_deviceQueue_proto_rawDescGZIP(), []int{5}
}

func (x *ListDeviceQueueItemsResponse) GetDeviceQueueItems() []*DeviceQueueItem {
	if x != nil {
		return x.DeviceQueueItems
	}
	return nil
}

var File_deviceQueue_proto protoreflect.FileDescriptor

var file_deviceQueue_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x65,
	0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x55, 0x49,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x13,
	0x0a, 0x05, 0x66, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66,
	0x43, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f,
	0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x71, 0x0a, 0x1d, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x50, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x75,
	0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x22, 0x35, 0x0a, 0x1e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x66, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x43, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x17, 0x46, 0x6c, 0x75,
	0x73, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x55, 0x49, 0x22, 0x36, 0x0a,
	0x1b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x76, 0x45, 0x55, 0x49, 0x22, 0x72, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xcd, 0x03, 0x0a, 0x12, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xad, 0x01, 0x0a, 0x07, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x32, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f,
	0x75, 0x69, 0x2e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x22, 0x2e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x64,
	0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x7d, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x73, 0x0a, 0x05, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e,
	0x46, 0x6c, 0x75, 0x73, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x7d, 0x2f,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x73, 0x5f, 0x75, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x73, 0x5f, 0x75, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x5f, 0x65,
	0x75, 0x69, 0x7d, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x63, 0x2d, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x70, 0x77, 0x61, 0x6e, 0x2d, 0x61, 0x70, 0x70,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x2d, 0x75, 0x69, 0x3b,
	0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73,
	0x5f, 0x75, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deviceQueue_proto_rawDescOnce sync.Once
	file_deviceQueue_proto_rawDescData = file_deviceQueue_proto_rawDesc
)

func file_deviceQueue_proto_rawDescGZIP() []byte {
	file_deviceQueue_proto_rawDescOnce.Do(func() {
		file_deviceQueue_proto_rawDescData = protoimpl.X.CompressGZIP(file_deviceQueue_proto_rawDescData)
	})
	return file_deviceQueue_proto_rawDescData
}

var file_deviceQueue_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_deviceQueue_proto_goTypes = []interface{}{
	(*DeviceQueueItem)(nil),                // 0: appserver_serves_ui.DeviceQueueItem
	(*EnqueueDeviceQueueItemRequest)(nil),  // 1: appserver_serves_ui.EnqueueDeviceQueueItemRequest
	(*EnqueueDeviceQueueItemResponse)(nil), // 2: appserver_serves_ui.EnqueueDeviceQueueItemResponse
	(*FlushDeviceQueueRequest)(nil),        // 3: appserver_serves_ui.FlushDeviceQueueRequest
	(*ListDeviceQueueItemsRequest)(nil),    // 4: appserver_serves_ui.ListDeviceQueueItemsRequest
	(*ListDeviceQueueItemsResponse)(nil),   // 5: appserver_serves_ui.ListDeviceQueueItemsResponse
	(*empty.Empty)(nil),                    // 6: google.protobuf.Empty
}
var file_deviceQueue_proto_depIdxs = []int32{
	0, // 0: appserver_serves_ui.EnqueueDeviceQueueItemRequest.device_queue_item:type_name -> appserver_serves_ui.DeviceQueueItem
	0, // 1: appserver_serves_ui.ListDeviceQueueItemsResponse.device_queue_items:type_name -> appserver_serves_ui.DeviceQueueItem
	1, // 2: appserver_serves_ui.DeviceQueueService.Enqueue:input_type -> appserver_serves_ui.EnqueueDeviceQueueItemRequest
	3, // 3: appserver_serves_ui.DeviceQueueService.Flush:input_type -> appserver_serves_ui.FlushDeviceQueueRequest
	4, // 4: appserver_serves_ui.DeviceQueueService.List:input_type -> appserver_serves_ui.ListDeviceQueueItemsRequest
	2, // 5: appserver_serves_ui.DeviceQueueService.Enqueue:output_type -> appserver_serves_ui.EnqueueDeviceQueueItemResponse
	6, // 6: appserver_serves_ui.DeviceQueueService.Flush:output_type -> google.protobuf.Empty
	5, // 7: appserver_serves_ui.DeviceQueueService.List:output_type -> appserver_serves_ui.ListDeviceQueueItemsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_deviceQueue_proto_init() }
func file_deviceQueue_proto_init() {
	if File_deviceQueue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deviceQueue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceQueueItem); i {
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
		file_deviceQueue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueDeviceQueueItemRequest); i {
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
		file_deviceQueue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueDeviceQueueItemResponse); i {
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
		file_deviceQueue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushDeviceQueueRequest); i {
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
		file_deviceQueue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceQueueItemsRequest); i {
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
		file_deviceQueue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceQueueItemsResponse); i {
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
			RawDescriptor: file_deviceQueue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deviceQueue_proto_goTypes,
		DependencyIndexes: file_deviceQueue_proto_depIdxs,
		MessageInfos:      file_deviceQueue_proto_msgTypes,
	}.Build()
	File_deviceQueue_proto = out.File
	file_deviceQueue_proto_rawDesc = nil
	file_deviceQueue_proto_goTypes = nil
	file_deviceQueue_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceQueueServiceClient is the client API for DeviceQueueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceQueueServiceClient interface {
	// Enqueue adds the given item to the device-queue.
	Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error)
	// Flush flushes the downlink device-queue.
	Flush(ctx context.Context, in *FlushDeviceQueueRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the items in the device-queue.
	List(ctx context.Context, in *ListDeviceQueueItemsRequest, opts ...grpc.CallOption) (*ListDeviceQueueItemsResponse, error)
}

type deviceQueueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceQueueServiceClient(cc grpc.ClientConnInterface) DeviceQueueServiceClient {
	return &deviceQueueServiceClient{cc}
}

func (c *deviceQueueServiceClient) Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error) {
	out := new(EnqueueDeviceQueueItemResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.DeviceQueueService/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceQueueServiceClient) Flush(ctx context.Context, in *FlushDeviceQueueRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.DeviceQueueService/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceQueueServiceClient) List(ctx context.Context, in *ListDeviceQueueItemsRequest, opts ...grpc.CallOption) (*ListDeviceQueueItemsResponse, error) {
	out := new(ListDeviceQueueItemsResponse)
	err := c.cc.Invoke(ctx, "/appserver_serves_ui.DeviceQueueService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceQueueServiceServer is the server API for DeviceQueueService service.
type DeviceQueueServiceServer interface {
	// Enqueue adds the given item to the device-queue.
	Enqueue(context.Context, *EnqueueDeviceQueueItemRequest) (*EnqueueDeviceQueueItemResponse, error)
	// Flush flushes the downlink device-queue.
	Flush(context.Context, *FlushDeviceQueueRequest) (*empty.Empty, error)
	// List lists the items in the device-queue.
	List(context.Context, *ListDeviceQueueItemsRequest) (*ListDeviceQueueItemsResponse, error)
}

// UnimplementedDeviceQueueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceQueueServiceServer struct {
}

func (*UnimplementedDeviceQueueServiceServer) Enqueue(context.Context, *EnqueueDeviceQueueItemRequest) (*EnqueueDeviceQueueItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (*UnimplementedDeviceQueueServiceServer) Flush(context.Context, *FlushDeviceQueueRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (*UnimplementedDeviceQueueServiceServer) List(context.Context, *ListDeviceQueueItemsRequest) (*ListDeviceQueueItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterDeviceQueueServiceServer(s *grpc.Server, srv DeviceQueueServiceServer) {
	s.RegisterService(&_DeviceQueueService_serviceDesc, srv)
}

func _DeviceQueueService_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDeviceQueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServiceServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.DeviceQueueService/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServiceServer).Enqueue(ctx, req.(*EnqueueDeviceQueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceQueueService_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushDeviceQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServiceServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.DeviceQueueService/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServiceServer).Flush(ctx, req.(*FlushDeviceQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceQueueService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceQueueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appserver_serves_ui.DeviceQueueService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServiceServer).List(ctx, req.(*ListDeviceQueueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceQueueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appserver_serves_ui.DeviceQueueService",
	HandlerType: (*DeviceQueueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _DeviceQueueService_Enqueue_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _DeviceQueueService_Flush_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceQueueService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deviceQueue.proto",
}
