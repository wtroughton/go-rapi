// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: response_list_exchange_permissions.proto

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

type ResponseListExchangePermissions_EntitlementFlag int32

const (
	ResponseListExchangePermissions_ENABLED  ResponseListExchangePermissions_EntitlementFlag = 1
	ResponseListExchangePermissions_DISABLED ResponseListExchangePermissions_EntitlementFlag = 2
)

// Enum value maps for ResponseListExchangePermissions_EntitlementFlag.
var (
	ResponseListExchangePermissions_EntitlementFlag_name = map[int32]string{
		1: "ENABLED",
		2: "DISABLED",
	}
	ResponseListExchangePermissions_EntitlementFlag_value = map[string]int32{
		"ENABLED":  1,
		"DISABLED": 2,
	}
)

func (x ResponseListExchangePermissions_EntitlementFlag) Enum() *ResponseListExchangePermissions_EntitlementFlag {
	p := new(ResponseListExchangePermissions_EntitlementFlag)
	*p = x
	return p
}

func (x ResponseListExchangePermissions_EntitlementFlag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseListExchangePermissions_EntitlementFlag) Descriptor() protoreflect.EnumDescriptor {
	return file_response_list_exchange_permissions_proto_enumTypes[0].Descriptor()
}

func (ResponseListExchangePermissions_EntitlementFlag) Type() protoreflect.EnumType {
	return &file_response_list_exchange_permissions_proto_enumTypes[0]
}

func (x ResponseListExchangePermissions_EntitlementFlag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ResponseListExchangePermissions_EntitlementFlag) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ResponseListExchangePermissions_EntitlementFlag(num)
	return nil
}

// Deprecated: Use ResponseListExchangePermissions_EntitlementFlag.Descriptor instead.
func (ResponseListExchangePermissions_EntitlementFlag) EnumDescriptor() ([]byte, []int) {
	return file_response_list_exchange_permissions_proto_rawDescGZIP(), []int{0, 0}
}

type ResponseListExchangePermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId      *int32                                           `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                                                                         // PB_OFFSET + MNM_TEMPLATE_ID
	UserMsg         []string                                         `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`                                                                                   // PB_OFFSET + MNM_USER_MSG
	RqHandlerRpCode []string                                         `protobuf:"bytes,132764,rep,name=rq_handler_rp_code,json=rqHandlerRpCode" json:"rq_handler_rp_code,omitempty"`                                                       // PB_OFFSET + MNM_REQUEST_HANDLER_RESPONSE_CODE
	RpCode          []string                                         `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`                                                                                      // PB_OFFSET + MNM_RESPONSE_CODE
	Exchange        *string                                          `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                                                                                // PB_OFFSET + MNM_EXCHANGE
	EntitlementFlag *ResponseListExchangePermissions_EntitlementFlag `protobuf:"varint,153400,opt,name=entitlement_flag,json=entitlementFlag,enum=rti.ResponseListExchangePermissions_EntitlementFlag" json:"entitlement_flag,omitempty"` // PB_OFFSET + MNM_ENTITLEMENT_FLAG
}

func (x *ResponseListExchangePermissions) Reset() {
	*x = ResponseListExchangePermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_list_exchange_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseListExchangePermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseListExchangePermissions) ProtoMessage() {}

func (x *ResponseListExchangePermissions) ProtoReflect() protoreflect.Message {
	mi := &file_response_list_exchange_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseListExchangePermissions.ProtoReflect.Descriptor instead.
func (*ResponseListExchangePermissions) Descriptor() ([]byte, []int) {
	return file_response_list_exchange_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseListExchangePermissions) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseListExchangePermissions) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseListExchangePermissions) GetRqHandlerRpCode() []string {
	if x != nil {
		return x.RqHandlerRpCode
	}
	return nil
}

func (x *ResponseListExchangePermissions) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseListExchangePermissions) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *ResponseListExchangePermissions) GetEntitlementFlag() ResponseListExchangePermissions_EntitlementFlag {
	if x != nil && x.EntitlementFlag != nil {
		return *x.EntitlementFlag
	}
	return ResponseListExchangePermissions_ENABLED
}

var File_response_list_exchange_permissions_proto protoreflect.FileDescriptor

var file_response_list_exchange_permissions_proto_rawDesc = []byte{
	0x0a, 0x28, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22,
	0xda, 0x02, 0x0a, 0x1f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x71, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x5f, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9c, 0x8d, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0f, 0x72, 0x71, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x70, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9e, 0x8d,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0xb8, 0xae, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x72, 0x74, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x0f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x2c,
	0x0a, 0x0f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x6c, 0x61,
	0x67, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02,
}

var (
	file_response_list_exchange_permissions_proto_rawDescOnce sync.Once
	file_response_list_exchange_permissions_proto_rawDescData = file_response_list_exchange_permissions_proto_rawDesc
)

func file_response_list_exchange_permissions_proto_rawDescGZIP() []byte {
	file_response_list_exchange_permissions_proto_rawDescOnce.Do(func() {
		file_response_list_exchange_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_list_exchange_permissions_proto_rawDescData)
	})
	return file_response_list_exchange_permissions_proto_rawDescData
}

var file_response_list_exchange_permissions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_response_list_exchange_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_list_exchange_permissions_proto_goTypes = []interface{}{
	(ResponseListExchangePermissions_EntitlementFlag)(0), // 0: rti.ResponseListExchangePermissions.EntitlementFlag
	(*ResponseListExchangePermissions)(nil),              // 1: rti.ResponseListExchangePermissions
}
var file_response_list_exchange_permissions_proto_depIdxs = []int32{
	0, // 0: rti.ResponseListExchangePermissions.entitlement_flag:type_name -> rti.ResponseListExchangePermissions.EntitlementFlag
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_response_list_exchange_permissions_proto_init() }
func file_response_list_exchange_permissions_proto_init() {
	if File_response_list_exchange_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_list_exchange_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseListExchangePermissions); i {
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
			RawDescriptor: file_response_list_exchange_permissions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_list_exchange_permissions_proto_goTypes,
		DependencyIndexes: file_response_list_exchange_permissions_proto_depIdxs,
		EnumInfos:         file_response_list_exchange_permissions_proto_enumTypes,
		MessageInfos:      file_response_list_exchange_permissions_proto_msgTypes,
	}.Build()
	File_response_list_exchange_permissions_proto = out.File
	file_response_list_exchange_permissions_proto_rawDesc = nil
	file_response_list_exchange_permissions_proto_goTypes = nil
	file_response_list_exchange_permissions_proto_depIdxs = nil
}
