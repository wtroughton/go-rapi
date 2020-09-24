// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: request_time_bar_update.proto

package rti

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

type RequestTimeBarUpdate_BarType int32

const (
	RequestTimeBarUpdate_SECOND_BAR RequestTimeBarUpdate_BarType = 1
	RequestTimeBarUpdate_MINUTE_BAR RequestTimeBarUpdate_BarType = 2
	RequestTimeBarUpdate_DAILY_BAR  RequestTimeBarUpdate_BarType = 3
	RequestTimeBarUpdate_WEEKLY_BAR RequestTimeBarUpdate_BarType = 4
)

// Enum value maps for RequestTimeBarUpdate_BarType.
var (
	RequestTimeBarUpdate_BarType_name = map[int32]string{
		1: "SECOND_BAR",
		2: "MINUTE_BAR",
		3: "DAILY_BAR",
		4: "WEEKLY_BAR",
	}
	RequestTimeBarUpdate_BarType_value = map[string]int32{
		"SECOND_BAR": 1,
		"MINUTE_BAR": 2,
		"DAILY_BAR":  3,
		"WEEKLY_BAR": 4,
	}
)

func (x RequestTimeBarUpdate_BarType) Enum() *RequestTimeBarUpdate_BarType {
	p := new(RequestTimeBarUpdate_BarType)
	*p = x
	return p
}

func (x RequestTimeBarUpdate_BarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestTimeBarUpdate_BarType) Descriptor() protoreflect.EnumDescriptor {
	return file_request_time_bar_update_proto_enumTypes[0].Descriptor()
}

func (RequestTimeBarUpdate_BarType) Type() protoreflect.EnumType {
	return &file_request_time_bar_update_proto_enumTypes[0]
}

func (x RequestTimeBarUpdate_BarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestTimeBarUpdate_BarType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestTimeBarUpdate_BarType(num)
	return nil
}

// Deprecated: Use RequestTimeBarUpdate_BarType.Descriptor instead.
func (RequestTimeBarUpdate_BarType) EnumDescriptor() ([]byte, []int) {
	return file_request_time_bar_update_proto_rawDescGZIP(), []int{0, 0}
}

type RequestTimeBarUpdate_Request int32

const (
	RequestTimeBarUpdate_SUBSCRIBE   RequestTimeBarUpdate_Request = 1
	RequestTimeBarUpdate_UNSUBSCRIBE RequestTimeBarUpdate_Request = 2
)

// Enum value maps for RequestTimeBarUpdate_Request.
var (
	RequestTimeBarUpdate_Request_name = map[int32]string{
		1: "SUBSCRIBE",
		2: "UNSUBSCRIBE",
	}
	RequestTimeBarUpdate_Request_value = map[string]int32{
		"SUBSCRIBE":   1,
		"UNSUBSCRIBE": 2,
	}
)

func (x RequestTimeBarUpdate_Request) Enum() *RequestTimeBarUpdate_Request {
	p := new(RequestTimeBarUpdate_Request)
	*p = x
	return p
}

func (x RequestTimeBarUpdate_Request) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestTimeBarUpdate_Request) Descriptor() protoreflect.EnumDescriptor {
	return file_request_time_bar_update_proto_enumTypes[1].Descriptor()
}

func (RequestTimeBarUpdate_Request) Type() protoreflect.EnumType {
	return &file_request_time_bar_update_proto_enumTypes[1]
}

func (x RequestTimeBarUpdate_Request) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestTimeBarUpdate_Request) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestTimeBarUpdate_Request(num)
	return nil
}

// Deprecated: Use RequestTimeBarUpdate_Request.Descriptor instead.
func (RequestTimeBarUpdate_Request) EnumDescriptor() ([]byte, []int) {
	return file_request_time_bar_update_proto_rawDescGZIP(), []int{0, 1}
}

type RequestTimeBarUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId    *int32                        `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                              // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg       []string                      `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                        // PB_OFFSET + MNM_USER_MSG
	Symbol        *string                       `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                                         // PB_OFFSET + MNM_SYMBOL
	Exchange      *string                       `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                                     // PB_OFFSET + MNM_EXCHANGE
	Request       *RequestTimeBarUpdate_Request `protobuf:"varint,100000,opt,name=request,enum=rti.RequestTimeBarUpdate_Request" json:"request,omitempty"`                // PB_OFFSET + MNM_REQUEST
	BarType       *RequestTimeBarUpdate_BarType `protobuf:"varint,119200,opt,name=bar_type,json=barType,enum=rti.RequestTimeBarUpdate_BarType" json:"bar_type,omitempty"` // PB_OFFSET + MNM_DATA_BAR_TYPE
	BarTypePeriod *int32                        `protobuf:"varint,119112,opt,name=bar_type_period,json=barTypePeriod" json:"bar_type_period,omitempty"`                   // PB_OFFSET + MNM_TIME_BAR_PERIOD
}

func (x *RequestTimeBarUpdate) Reset() {
	*x = RequestTimeBarUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_time_bar_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTimeBarUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTimeBarUpdate) ProtoMessage() {}

func (x *RequestTimeBarUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_request_time_bar_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTimeBarUpdate.ProtoReflect.Descriptor instead.
func (*RequestTimeBarUpdate) Descriptor() ([]byte, []int) {
	return file_request_time_bar_update_proto_rawDescGZIP(), []int{0}
}

func (x *RequestTimeBarUpdate) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestTimeBarUpdate) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestTimeBarUpdate) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *RequestTimeBarUpdate) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *RequestTimeBarUpdate) GetRequest() RequestTimeBarUpdate_Request {
	if x != nil && x.Request != nil {
		return *x.Request
	}
	return RequestTimeBarUpdate_SUBSCRIBE
}

func (x *RequestTimeBarUpdate) GetBarType() RequestTimeBarUpdate_BarType {
	if x != nil && x.BarType != nil {
		return *x.BarType
	}
	return RequestTimeBarUpdate_SECOND_BAR
}

func (x *RequestTimeBarUpdate) GetBarTypePeriod() int32 {
	if x != nil && x.BarTypePeriod != nil {
		return *x.BarTypePeriod
	}
	return 0
}

var File_request_time_bar_update_proto protoreflect.FileDescriptor

var file_request_time_bar_update_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x62,
	0x61, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x74, 0x69, 0x22, 0xac, 0x03, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0xa0, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x61, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x62, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0xa0, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x61, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x62, 0x61, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0xc8, 0xa2, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x62, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x48,
	0x0a, 0x07, 0x42, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x43,
	0x4f, 0x4e, 0x44, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x49, 0x4e,
	0x55, 0x54, 0x45, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x41, 0x49,
	0x4c, 0x59, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x45, 0x45, 0x4b,
	0x4c, 0x59, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x04, 0x22, 0x29, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42,
	0x45, 0x10, 0x02,
}

var (
	file_request_time_bar_update_proto_rawDescOnce sync.Once
	file_request_time_bar_update_proto_rawDescData = file_request_time_bar_update_proto_rawDesc
)

func file_request_time_bar_update_proto_rawDescGZIP() []byte {
	file_request_time_bar_update_proto_rawDescOnce.Do(func() {
		file_request_time_bar_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_time_bar_update_proto_rawDescData)
	})
	return file_request_time_bar_update_proto_rawDescData
}

var file_request_time_bar_update_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_request_time_bar_update_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_time_bar_update_proto_goTypes = []interface{}{
	(RequestTimeBarUpdate_BarType)(0), // 0: rti.RequestTimeBarUpdate.BarType
	(RequestTimeBarUpdate_Request)(0), // 1: rti.RequestTimeBarUpdate.Request
	(*RequestTimeBarUpdate)(nil),      // 2: rti.RequestTimeBarUpdate
}
var file_request_time_bar_update_proto_depIdxs = []int32{
	1, // 0: rti.RequestTimeBarUpdate.request:type_name -> rti.RequestTimeBarUpdate.Request
	0, // 1: rti.RequestTimeBarUpdate.bar_type:type_name -> rti.RequestTimeBarUpdate.BarType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_request_time_bar_update_proto_init() }
func file_request_time_bar_update_proto_init() {
	if File_request_time_bar_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_time_bar_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTimeBarUpdate); i {
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
			RawDescriptor: file_request_time_bar_update_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_time_bar_update_proto_goTypes,
		DependencyIndexes: file_request_time_bar_update_proto_depIdxs,
		EnumInfos:         file_request_time_bar_update_proto_enumTypes,
		MessageInfos:      file_request_time_bar_update_proto_msgTypes,
	}.Build()
	File_request_time_bar_update_proto = out.File
	file_request_time_bar_update_proto_rawDesc = nil
	file_request_time_bar_update_proto_goTypes = nil
	file_request_time_bar_update_proto_depIdxs = nil
}
