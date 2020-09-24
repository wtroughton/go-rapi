// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: request_update_stop_bracket_level.proto

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

type RequestUpdateStopBracketLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"` // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg    []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`           // PB_OFFSET + MNM_USER_MSG
	FcmId      *string  `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`                 // PB_OFFSET + MNM_FCM_ID
	IbId       *string  `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`                    // PB_OFFSET + MNM_IB_ID
	AccountId  *string  `protobuf:"bytes,154008,opt,name=account_id,json=accountId" json:"account_id,omitempty"`     // PB_OFFSET + MNM_ACCOUNT_ID
	BasketId   *string  `protobuf:"bytes,110300,opt,name=basket_id,json=basketId" json:"basket_id,omitempty"`        // PB_OFFSET + MNM_BASKET_ID
	Level      *int32   `protobuf:"varint,154244,opt,name=level" json:"level,omitempty"`                             // PB_OFFSET + MNM_LEVEL
	StopTicks  *int32   `protobuf:"varint,154458,opt,name=stop_ticks,json=stopTicks" json:"stop_ticks,omitempty"`    // PB_OFFSET + MNM_STOP_TICKS
}

func (x *RequestUpdateStopBracketLevel) Reset() {
	*x = RequestUpdateStopBracketLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_update_stop_bracket_level_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUpdateStopBracketLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUpdateStopBracketLevel) ProtoMessage() {}

func (x *RequestUpdateStopBracketLevel) ProtoReflect() protoreflect.Message {
	mi := &file_request_update_stop_bracket_level_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUpdateStopBracketLevel.ProtoReflect.Descriptor instead.
func (*RequestUpdateStopBracketLevel) Descriptor() ([]byte, []int) {
	return file_request_update_stop_bracket_level_proto_rawDescGZIP(), []int{0}
}

func (x *RequestUpdateStopBracketLevel) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestUpdateStopBracketLevel) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *RequestUpdateStopBracketLevel) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *RequestUpdateStopBracketLevel) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *RequestUpdateStopBracketLevel) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *RequestUpdateStopBracketLevel) GetBasketId() string {
	if x != nil && x.BasketId != nil {
		return *x.BasketId
	}
	return ""
}

func (x *RequestUpdateStopBracketLevel) GetLevel() int32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *RequestUpdateStopBracketLevel) GetStopTicks() int32 {
	if x != nil && x.StopTicks != nil {
		return *x.StopTicks
	}
	return 0
}

var File_request_update_stop_bracket_level_proto protoreflect.FileDescriptor

var file_request_update_stop_bracket_level_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x62, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0x88,
	0x02, 0x0a, 0x1d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x70, 0x42, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67,
	0x12, 0x17, 0x0a, 0x06, 0x66, 0x63, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x9d, 0xb3, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x69, 0x62, 0x5f,
	0x69, 0x64, 0x18, 0x9e, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x62, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x98,
	0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0xdc,
	0xdd, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x84, 0xb5, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x70,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x18, 0xda, 0xb6, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x73, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x63, 0x6b, 0x73,
}

var (
	file_request_update_stop_bracket_level_proto_rawDescOnce sync.Once
	file_request_update_stop_bracket_level_proto_rawDescData = file_request_update_stop_bracket_level_proto_rawDesc
)

func file_request_update_stop_bracket_level_proto_rawDescGZIP() []byte {
	file_request_update_stop_bracket_level_proto_rawDescOnce.Do(func() {
		file_request_update_stop_bracket_level_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_update_stop_bracket_level_proto_rawDescData)
	})
	return file_request_update_stop_bracket_level_proto_rawDescData
}

var file_request_update_stop_bracket_level_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_update_stop_bracket_level_proto_goTypes = []interface{}{
	(*RequestUpdateStopBracketLevel)(nil), // 0: rti.RequestUpdateStopBracketLevel
}
var file_request_update_stop_bracket_level_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_update_stop_bracket_level_proto_init() }
func file_request_update_stop_bracket_level_proto_init() {
	if File_request_update_stop_bracket_level_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_update_stop_bracket_level_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUpdateStopBracketLevel); i {
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
			RawDescriptor: file_request_update_stop_bracket_level_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_update_stop_bracket_level_proto_goTypes,
		DependencyIndexes: file_request_update_stop_bracket_level_proto_depIdxs,
		MessageInfos:      file_request_update_stop_bracket_level_proto_msgTypes,
	}.Build()
	File_request_update_stop_bracket_level_proto = out.File
	file_request_update_stop_bracket_level_proto_rawDesc = nil
	file_request_update_stop_bracket_level_proto_goTypes = nil
	file_request_update_stop_bracket_level_proto_depIdxs = nil
}
