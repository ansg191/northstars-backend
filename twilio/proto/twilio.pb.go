// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/twilio.proto

package twilio

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message_Status int32

const (
	Message_ACCEPTED    Message_Status = 0
	Message_SCHEDULED   Message_Status = 1
	Message_QUEUED      Message_Status = 2
	Message_SENDING     Message_Status = 3
	Message_SENT        Message_Status = 4
	Message_RECEIVING   Message_Status = 5
	Message_RECEIVED    Message_Status = 6
	Message_DELIVERED   Message_Status = 7
	Message_UNDELIVERED Message_Status = 8
	Message_FAILED      Message_Status = 9
	Message_READ        Message_Status = 10
	Message_CANCELED    Message_Status = 11
)

// Enum value maps for Message_Status.
var (
	Message_Status_name = map[int32]string{
		0:  "ACCEPTED",
		1:  "SCHEDULED",
		2:  "QUEUED",
		3:  "SENDING",
		4:  "SENT",
		5:  "RECEIVING",
		6:  "RECEIVED",
		7:  "DELIVERED",
		8:  "UNDELIVERED",
		9:  "FAILED",
		10: "READ",
		11: "CANCELED",
	}
	Message_Status_value = map[string]int32{
		"ACCEPTED":    0,
		"SCHEDULED":   1,
		"QUEUED":      2,
		"SENDING":     3,
		"SENT":        4,
		"RECEIVING":   5,
		"RECEIVED":    6,
		"DELIVERED":   7,
		"UNDELIVERED": 8,
		"FAILED":      9,
		"READ":        10,
		"CANCELED":    11,
	}
)

func (x Message_Status) Enum() *Message_Status {
	p := new(Message_Status)
	*p = x
	return p
}

func (x Message_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_twilio_proto_enumTypes[0].Descriptor()
}

func (Message_Status) Type() protoreflect.EnumType {
	return &file_proto_twilio_proto_enumTypes[0]
}

func (x Message_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_Status.Descriptor instead.
func (Message_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_twilio_proto_rawDescGZIP(), []int{0, 0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body        string                 `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=dateCreated,proto3" json:"dateCreated,omitempty"`
	DateSent    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=dateSent,proto3" json:"dateSent,omitempty"`
	DateUpdated *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=dateUpdated,proto3" json:"dateUpdated,omitempty"`
	Direction   string                 `protobuf:"bytes,5,opt,name=direction,proto3" json:"direction,omitempty"`
	From        string                 `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	NumMedia    string                 `protobuf:"bytes,7,opt,name=numMedia,proto3" json:"numMedia,omitempty"`
	NumSegments string                 `protobuf:"bytes,8,opt,name=numSegments,proto3" json:"numSegments,omitempty"`
	Price       string                 `protobuf:"bytes,9,opt,name=price,proto3" json:"price,omitempty"`
	PriceUnit   string                 `protobuf:"bytes,10,opt,name=priceUnit,proto3" json:"priceUnit,omitempty"`
	Sid         string                 `protobuf:"bytes,11,opt,name=sid,proto3" json:"sid,omitempty"`
	Status      Message_Status         `protobuf:"varint,12,opt,name=status,proto3,enum=twilio.Message_Status" json:"status,omitempty"`
	To          string                 `protobuf:"bytes,14,opt,name=to,proto3" json:"to,omitempty"`
	Uri         string                 `protobuf:"bytes,15,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twilio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twilio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_twilio_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Message) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

func (x *Message) GetDateSent() *timestamppb.Timestamp {
	if x != nil {
		return x.DateSent
	}
	return nil
}

func (x *Message) GetDateUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateUpdated
	}
	return nil
}

func (x *Message) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Message) GetNumMedia() string {
	if x != nil {
		return x.NumMedia
	}
	return ""
}

func (x *Message) GetNumSegments() string {
	if x != nil {
		return x.NumSegments
	}
	return ""
}

func (x *Message) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Message) GetPriceUnit() string {
	if x != nil {
		return x.PriceUnit
	}
	return ""
}

func (x *Message) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *Message) GetStatus() Message_Status {
	if x != nil {
		return x.Status
	}
	return Message_ACCEPTED
}

func (x *Message) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Message) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To   string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twilio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twilio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_twilio_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendMessageRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twilio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twilio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_twilio_proto_rawDescGZIP(), []int{2}
}

func (x *SendMessageResponse) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid string `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twilio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twilio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_twilio_proto_rawDescGZIP(), []int{3}
}

func (x *GetMessageRequest) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

type GetMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetMessageResponse) Reset() {
	*x = GetMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twilio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse) ProtoMessage() {}

func (x *GetMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twilio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_twilio_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessageResponse) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_proto_twilio_proto protoreflect.FileDescriptor

var file_proto_twilio_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x05,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3c, 0x0a,
	0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0xa9, 0x01, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45,
	0x4e, 0x54, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x49, 0x4e,
	0x47, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10,
	0x06, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x07,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10,
	0x08, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x12, 0x08, 0x0a,
	0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x45, 0x44, 0x10, 0x0b, 0x22, 0x38, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22,
	0x38, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64,
	0x22, 0x37, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x99, 0x01, 0x0a, 0x06, 0x54, 0x77,
	0x69, 0x6c, 0x69, 0x6f, 0x12, 0x48, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x74,
	0x77, 0x69, 0x6c, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_twilio_proto_rawDescOnce sync.Once
	file_proto_twilio_proto_rawDescData = file_proto_twilio_proto_rawDesc
)

func file_proto_twilio_proto_rawDescGZIP() []byte {
	file_proto_twilio_proto_rawDescOnce.Do(func() {
		file_proto_twilio_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_twilio_proto_rawDescData)
	})
	return file_proto_twilio_proto_rawDescData
}

var file_proto_twilio_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_twilio_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_twilio_proto_goTypes = []interface{}{
	(Message_Status)(0),           // 0: twilio.Message.Status
	(*Message)(nil),               // 1: twilio.Message
	(*SendMessageRequest)(nil),    // 2: twilio.SendMessageRequest
	(*SendMessageResponse)(nil),   // 3: twilio.SendMessageResponse
	(*GetMessageRequest)(nil),     // 4: twilio.GetMessageRequest
	(*GetMessageResponse)(nil),    // 5: twilio.GetMessageResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_proto_twilio_proto_depIdxs = []int32{
	6, // 0: twilio.Message.dateCreated:type_name -> google.protobuf.Timestamp
	6, // 1: twilio.Message.dateSent:type_name -> google.protobuf.Timestamp
	6, // 2: twilio.Message.dateUpdated:type_name -> google.protobuf.Timestamp
	0, // 3: twilio.Message.status:type_name -> twilio.Message.Status
	1, // 4: twilio.SendMessageResponse.msg:type_name -> twilio.Message
	1, // 5: twilio.GetMessageResponse.msg:type_name -> twilio.Message
	2, // 6: twilio.Twilio.SendMessage:input_type -> twilio.SendMessageRequest
	4, // 7: twilio.Twilio.GetMessage:input_type -> twilio.GetMessageRequest
	3, // 8: twilio.Twilio.SendMessage:output_type -> twilio.SendMessageResponse
	5, // 9: twilio.Twilio.GetMessage:output_type -> twilio.GetMessageResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_twilio_proto_init() }
func file_proto_twilio_proto_init() {
	if File_proto_twilio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_twilio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_twilio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_proto_twilio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_proto_twilio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
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
		file_proto_twilio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResponse); i {
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
			RawDescriptor: file_proto_twilio_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_twilio_proto_goTypes,
		DependencyIndexes: file_proto_twilio_proto_depIdxs,
		EnumInfos:         file_proto_twilio_proto_enumTypes,
		MessageInfos:      file_proto_twilio_proto_msgTypes,
	}.Build()
	File_proto_twilio_proto = out.File
	file_proto_twilio_proto_rawDesc = nil
	file_proto_twilio_proto_goTypes = nil
	file_proto_twilio_proto_depIdxs = nil
}
