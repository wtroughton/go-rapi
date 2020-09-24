// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: response_login_info.proto

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

type ResponseLoginInfo_UserType int32

const (
	ResponseLoginInfo_USER_TYPE_ADMIN  ResponseLoginInfo_UserType = 0
	ResponseLoginInfo_USER_TYPE_FCM    ResponseLoginInfo_UserType = 1
	ResponseLoginInfo_USER_TYPE_IB     ResponseLoginInfo_UserType = 2
	ResponseLoginInfo_USER_TYPE_TRADER ResponseLoginInfo_UserType = 3
)

// Enum value maps for ResponseLoginInfo_UserType.
var (
	ResponseLoginInfo_UserType_name = map[int32]string{
		0: "USER_TYPE_ADMIN",
		1: "USER_TYPE_FCM",
		2: "USER_TYPE_IB",
		3: "USER_TYPE_TRADER",
	}
	ResponseLoginInfo_UserType_value = map[string]int32{
		"USER_TYPE_ADMIN":  0,
		"USER_TYPE_FCM":    1,
		"USER_TYPE_IB":     2,
		"USER_TYPE_TRADER": 3,
	}
)

func (x ResponseLoginInfo_UserType) Enum() *ResponseLoginInfo_UserType {
	p := new(ResponseLoginInfo_UserType)
	*p = x
	return p
}

func (x ResponseLoginInfo_UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseLoginInfo_UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_response_login_info_proto_enumTypes[0].Descriptor()
}

func (ResponseLoginInfo_UserType) Type() protoreflect.EnumType {
	return &file_response_login_info_proto_enumTypes[0]
}

func (x ResponseLoginInfo_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ResponseLoginInfo_UserType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ResponseLoginInfo_UserType(num)
	return nil
}

// Deprecated: Use ResponseLoginInfo_UserType.Descriptor instead.
func (ResponseLoginInfo_UserType) EnumDescriptor() ([]byte, []int) {
	return file_response_login_info_proto_rawDescGZIP(), []int{0, 0}
}

type ResponseLoginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId *int32                      `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                               // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg    []string                    `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                         // PB_OFFSET + MNM_USER_MSG
	RpCode     []string                    `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`                                            // PB_OFFSET + MNM_RESPONSE_CODE
	FcmId      *string                     `protobuf:"bytes,154013,opt,name=fcm_id,json=fcmId" json:"fcm_id,omitempty"`                                               // PB_OFFSET + MNM_FCM_ID
	IbId       *string                     `protobuf:"bytes,154014,opt,name=ib_id,json=ibId" json:"ib_id,omitempty"`                                                  // PB_OFFSET + MNM_IB_ID
	FirstName  *string                     `protobuf:"bytes,154216,opt,name=first_name,json=firstName" json:"first_name,omitempty"`                                   // PB_OFFSET + MNM_FIRST_NAME
	LastName   *string                     `protobuf:"bytes,154217,opt,name=last_name,json=lastName" json:"last_name,omitempty"`                                      // PB_OFFSET + MNM_LAST_NAME
	UserType   *ResponseLoginInfo_UserType `protobuf:"varint,154036,opt,name=user_type,json=userType,enum=rti.ResponseLoginInfo_UserType" json:"user_type,omitempty"` // PB_OFFSET + MNM_USER_TYPE
}

func (x *ResponseLoginInfo) Reset() {
	*x = ResponseLoginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_login_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLoginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLoginInfo) ProtoMessage() {}

func (x *ResponseLoginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_response_login_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLoginInfo.ProtoReflect.Descriptor instead.
func (*ResponseLoginInfo) Descriptor() ([]byte, []int) {
	return file_response_login_info_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseLoginInfo) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseLoginInfo) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseLoginInfo) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseLoginInfo) GetFcmId() string {
	if x != nil && x.FcmId != nil {
		return *x.FcmId
	}
	return ""
}

func (x *ResponseLoginInfo) GetIbId() string {
	if x != nil && x.IbId != nil {
		return *x.IbId
	}
	return ""
}

func (x *ResponseLoginInfo) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *ResponseLoginInfo) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *ResponseLoginInfo) GetUserType() ResponseLoginInfo_UserType {
	if x != nil && x.UserType != nil {
		return *x.UserType
	}
	return ResponseLoginInfo_USER_TYPE_ADMIN
}

var File_response_login_info_proto protoreflect.FileDescriptor

var file_response_login_info_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69,
	0x22, 0xfa, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x9e, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x17, 0x0a, 0x06, 0x66, 0x63, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x9d, 0xb3, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x69, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x9e, 0xb3, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x62, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0xe8, 0xb4, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0xe9, 0xb4, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x3e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xb4,
	0xb3, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x5a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x46, 0x43, 0x4d, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x42, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x45, 0x52, 0x10, 0x03,
}

var (
	file_response_login_info_proto_rawDescOnce sync.Once
	file_response_login_info_proto_rawDescData = file_response_login_info_proto_rawDesc
)

func file_response_login_info_proto_rawDescGZIP() []byte {
	file_response_login_info_proto_rawDescOnce.Do(func() {
		file_response_login_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_login_info_proto_rawDescData)
	})
	return file_response_login_info_proto_rawDescData
}

var file_response_login_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_response_login_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_login_info_proto_goTypes = []interface{}{
	(ResponseLoginInfo_UserType)(0), // 0: rti.ResponseLoginInfo.UserType
	(*ResponseLoginInfo)(nil),       // 1: rti.ResponseLoginInfo
}
var file_response_login_info_proto_depIdxs = []int32{
	0, // 0: rti.ResponseLoginInfo.user_type:type_name -> rti.ResponseLoginInfo.UserType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_response_login_info_proto_init() }
func file_response_login_info_proto_init() {
	if File_response_login_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_login_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLoginInfo); i {
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
			RawDescriptor: file_response_login_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_login_info_proto_goTypes,
		DependencyIndexes: file_response_login_info_proto_depIdxs,
		EnumInfos:         file_response_login_info_proto_enumTypes,
		MessageInfos:      file_response_login_info_proto_msgTypes,
	}.Build()
	File_response_login_info_proto = out.File
	file_response_login_info_proto_rawDesc = nil
	file_response_login_info_proto_goTypes = nil
	file_response_login_info_proto_depIdxs = nil
}
