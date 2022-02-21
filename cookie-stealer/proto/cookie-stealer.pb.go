// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/cookie-stealer.proto

package cookiestealer

import (
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

type StealTeamUnifyCookiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StealTeamUnifyCookiesRequest) Reset() {
	*x = StealTeamUnifyCookiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cookie_stealer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StealTeamUnifyCookiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StealTeamUnifyCookiesRequest) ProtoMessage() {}

func (x *StealTeamUnifyCookiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cookie_stealer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StealTeamUnifyCookiesRequest.ProtoReflect.Descriptor instead.
func (*StealTeamUnifyCookiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_cookie_stealer_proto_rawDescGZIP(), []int{0}
}

type StealTeamUnifyCookiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cookies string `protobuf:"bytes,1,opt,name=cookies,proto3" json:"cookies,omitempty"`
	Unready bool   `protobuf:"varint,2,opt,name=unready,proto3" json:"unready,omitempty"`
}

func (x *StealTeamUnifyCookiesResponse) Reset() {
	*x = StealTeamUnifyCookiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cookie_stealer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StealTeamUnifyCookiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StealTeamUnifyCookiesResponse) ProtoMessage() {}

func (x *StealTeamUnifyCookiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cookie_stealer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StealTeamUnifyCookiesResponse.ProtoReflect.Descriptor instead.
func (*StealTeamUnifyCookiesResponse) Descriptor() ([]byte, []int) {
	return file_proto_cookie_stealer_proto_rawDescGZIP(), []int{1}
}

func (x *StealTeamUnifyCookiesResponse) GetCookies() string {
	if x != nil {
		return x.Cookies
	}
	return ""
}

func (x *StealTeamUnifyCookiesResponse) GetUnready() bool {
	if x != nil {
		return x.Unready
	}
	return false
}

var File_proto_cookie_stealer_proto protoreflect.FileDescriptor

var file_proto_cookie_stealer_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2d, 0x73,
	0x74, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x73, 0x74, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x1c, 0x53,
	0x74, 0x65, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x6e, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x1d, 0x53,
	0x74, 0x65, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x6e, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x32, 0x85, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x53, 0x74, 0x65, 0x61, 0x6c,
	0x65, 0x72, 0x12, 0x74, 0x0a, 0x15, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x55,
	0x6e, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x73, 0x74, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x61,
	0x6c, 0x54, 0x65, 0x61, 0x6d, 0x55, 0x6e, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6f, 0x6b, 0x69,
	0x65, 0x73, 0x74, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x54, 0x65,
	0x61, 0x6d, 0x55, 0x6e, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x74, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cookie_stealer_proto_rawDescOnce sync.Once
	file_proto_cookie_stealer_proto_rawDescData = file_proto_cookie_stealer_proto_rawDesc
)

func file_proto_cookie_stealer_proto_rawDescGZIP() []byte {
	file_proto_cookie_stealer_proto_rawDescOnce.Do(func() {
		file_proto_cookie_stealer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cookie_stealer_proto_rawDescData)
	})
	return file_proto_cookie_stealer_proto_rawDescData
}

var file_proto_cookie_stealer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_cookie_stealer_proto_goTypes = []interface{}{
	(*StealTeamUnifyCookiesRequest)(nil),  // 0: cookiestealer.StealTeamUnifyCookiesRequest
	(*StealTeamUnifyCookiesResponse)(nil), // 1: cookiestealer.StealTeamUnifyCookiesResponse
}
var file_proto_cookie_stealer_proto_depIdxs = []int32{
	0, // 0: cookiestealer.CookieStealer.StealTeamUnifyCookies:input_type -> cookiestealer.StealTeamUnifyCookiesRequest
	1, // 1: cookiestealer.CookieStealer.StealTeamUnifyCookies:output_type -> cookiestealer.StealTeamUnifyCookiesResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_cookie_stealer_proto_init() }
func file_proto_cookie_stealer_proto_init() {
	if File_proto_cookie_stealer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cookie_stealer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StealTeamUnifyCookiesRequest); i {
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
		file_proto_cookie_stealer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StealTeamUnifyCookiesResponse); i {
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
			RawDescriptor: file_proto_cookie_stealer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cookie_stealer_proto_goTypes,
		DependencyIndexes: file_proto_cookie_stealer_proto_depIdxs,
		MessageInfos:      file_proto_cookie_stealer_proto_msgTypes,
	}.Build()
	File_proto_cookie_stealer_proto = out.File
	file_proto_cookie_stealer_proto_rawDesc = nil
	file_proto_cookie_stealer_proto_goTypes = nil
	file_proto_cookie_stealer_proto_depIdxs = nil
}
