// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: request_list_unaccepted_agreements.proto

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

type RequestListUnacceptedAgreements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"` // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg    []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`           // PB_OFFSET + MNM_USER_MSG
}

func (x *RequestListUnacceptedAgreements) Reset() {
	*x = RequestListUnacceptedAgreements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_list_unaccepted_agreements_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestListUnacceptedAgreements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestListUnacceptedAgreements) ProtoMessage() {}

func (x *RequestListUnacceptedAgreements) ProtoReflect() protoreflect.Message {
	mi := &file_request_list_unaccepted_agreements_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestListUnacceptedAgreements.ProtoReflect.Descriptor instead.
func (*RequestListUnacceptedAgreements) Descriptor() ([]byte, []int) {
	return file_request_list_unaccepted_agreements_proto_rawDescGZIP(), []int{0}
}

func (x *RequestListUnacceptedAgreements) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *RequestListUnacceptedAgreements) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

var File_request_list_unaccepted_agreements_proto protoreflect.FileDescriptor

var file_request_list_unaccepted_agreements_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x75,
	0x6e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22,
	0x61, 0x0a, 0x1f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4d,
	0x73, 0x67,
}

var (
	file_request_list_unaccepted_agreements_proto_rawDescOnce sync.Once
	file_request_list_unaccepted_agreements_proto_rawDescData = file_request_list_unaccepted_agreements_proto_rawDesc
)

func file_request_list_unaccepted_agreements_proto_rawDescGZIP() []byte {
	file_request_list_unaccepted_agreements_proto_rawDescOnce.Do(func() {
		file_request_list_unaccepted_agreements_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_list_unaccepted_agreements_proto_rawDescData)
	})
	return file_request_list_unaccepted_agreements_proto_rawDescData
}

var file_request_list_unaccepted_agreements_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_list_unaccepted_agreements_proto_goTypes = []interface{}{
	(*RequestListUnacceptedAgreements)(nil), // 0: rti.RequestListUnacceptedAgreements
}
var file_request_list_unaccepted_agreements_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_list_unaccepted_agreements_proto_init() }
func file_request_list_unaccepted_agreements_proto_init() {
	if File_request_list_unaccepted_agreements_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_list_unaccepted_agreements_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestListUnacceptedAgreements); i {
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
			RawDescriptor: file_request_list_unaccepted_agreements_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_list_unaccepted_agreements_proto_goTypes,
		DependencyIndexes: file_request_list_unaccepted_agreements_proto_depIdxs,
		MessageInfos:      file_request_list_unaccepted_agreements_proto_msgTypes,
	}.Build()
	File_request_list_unaccepted_agreements_proto = out.File
	file_request_list_unaccepted_agreements_proto_rawDesc = nil
	file_request_list_unaccepted_agreements_proto_goTypes = nil
	file_request_list_unaccepted_agreements_proto_depIdxs = nil
}
