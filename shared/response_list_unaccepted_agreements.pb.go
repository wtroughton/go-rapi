// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: response_list_unaccepted_agreements.proto

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

type ResponseListUnacceptedAgreements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId      *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                   // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg         []string `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                             // PB_OFFSET + MNM_USER_MSG
	RqHandlerRpCode []string `protobuf:"bytes,132764,rep,name=rq_handler_rp_code,json=rqHandlerRpCode" json:"rq_handler_rp_code,omitempty"` // PB_OFFSET + MNM_REQUEST_HANDLER_RESPONSE_CODE
	RpCode          []string `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`                                // PB_OFFSET + MNM_RESPONSE_CODE
	FcmId           *string  `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`                                   // PB_OFFSET + MNM_FCM_ID
	IbId            *string  `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`                                      // PB_OFFSET + MNM_IB_ID
	User            *string  `protobuf:"bytes,131003,opt,name=user" json:"user,omitempty"`                                                  // PB_OFFSET + MNM_USER
	AgreementTitle  *string  `protobuf:"bytes,153406,opt,name=agreement_title,json=agreementTitle" json:"agreement_title,omitempty"`        // PB_OFFSET + MNM_AGREEMENT_TITLE
	AgreementId     *string  `protobuf:"bytes,153407,opt,name=agreement_id,json=agreementId" json:"agreement_id,omitempty"`                 // PB_OFFSET + MNM_AGREEMENT_ID
}

func (x *ResponseListUnacceptedAgreements) Reset() {
	*x = ResponseListUnacceptedAgreements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_list_unaccepted_agreements_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseListUnacceptedAgreements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseListUnacceptedAgreements) ProtoMessage() {}

func (x *ResponseListUnacceptedAgreements) ProtoReflect() protoreflect.Message {
	mi := &file_response_list_unaccepted_agreements_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseListUnacceptedAgreements.ProtoReflect.Descriptor instead.
func (*ResponseListUnacceptedAgreements) Descriptor() ([]byte, []int) {
	return file_response_list_unaccepted_agreements_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseListUnacceptedAgreements) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseListUnacceptedAgreements) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseListUnacceptedAgreements) GetRqHandlerRpCode() []string {
	if x != nil {
		return x.RqHandlerRpCode
	}
	return nil
}

func (x *ResponseListUnacceptedAgreements) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseListUnacceptedAgreements) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *ResponseListUnacceptedAgreements) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *ResponseListUnacceptedAgreements) GetUser() string {
	if x != nil && x.User != nil {
		return *x.User
	}
	return ""
}

func (x *ResponseListUnacceptedAgreements) GetAgreementTitle() string {
	if x != nil && x.AgreementTitle != nil {
		return *x.AgreementTitle
	}
	return ""
}

func (x *ResponseListUnacceptedAgreements) GetAgreementId() string {
	if x != nil && x.AgreementId != nil {
		return *x.AgreementId
	}
	return ""
}

var File_response_list_unaccepted_agreements_proto protoreflect.FileDescriptor

var file_response_list_unaccepted_agreements_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x75, 0x6e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x67, 0x72, 0x65, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69,
	0x22, 0xc2, 0x02, 0x0a, 0x20, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x6e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x41, 0x67, 0x72, 0x65, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x71, 0x5f, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9c, 0x8d, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x71, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x9e, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x17, 0x0a, 0x06, 0x66, 0x63, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x9d, 0xb3, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x63, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x69, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x9e, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x62, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0xbb, 0xff, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0f, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0xbe, 0xae, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x23, 0x0a, 0x0c, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0xbf, 0xae, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64,
}

var (
	file_response_list_unaccepted_agreements_proto_rawDescOnce sync.Once
	file_response_list_unaccepted_agreements_proto_rawDescData = file_response_list_unaccepted_agreements_proto_rawDesc
)

func file_response_list_unaccepted_agreements_proto_rawDescGZIP() []byte {
	file_response_list_unaccepted_agreements_proto_rawDescOnce.Do(func() {
		file_response_list_unaccepted_agreements_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_list_unaccepted_agreements_proto_rawDescData)
	})
	return file_response_list_unaccepted_agreements_proto_rawDescData
}

var file_response_list_unaccepted_agreements_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_list_unaccepted_agreements_proto_goTypes = []interface{}{
	(*ResponseListUnacceptedAgreements)(nil), // 0: rti.ResponseListUnacceptedAgreements
}
var file_response_list_unaccepted_agreements_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_list_unaccepted_agreements_proto_init() }
func file_response_list_unaccepted_agreements_proto_init() {
	if File_response_list_unaccepted_agreements_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_list_unaccepted_agreements_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseListUnacceptedAgreements); i {
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
			RawDescriptor: file_response_list_unaccepted_agreements_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_list_unaccepted_agreements_proto_goTypes,
		DependencyIndexes: file_response_list_unaccepted_agreements_proto_depIdxs,
		MessageInfos:      file_response_list_unaccepted_agreements_proto_msgTypes,
	}.Build()
	File_response_list_unaccepted_agreements_proto = out.File
	file_response_list_unaccepted_agreements_proto_rawDesc = nil
	file_response_list_unaccepted_agreements_proto_goTypes = nil
	file_response_list_unaccepted_agreements_proto_depIdxs = nil
}
