// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: depth_by_order_end_event.proto

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

type DepthByOrderEndEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId     *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`             // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol         []string `protobuf:"bytes,110100,rep,name=symbol" json:"symbol,omitempty"`                                        // PB_OFFSET + MNM_SYMBOL
	Exchange       []string `protobuf:"bytes,110101,rep,name=exchange" json:"exchange,omitempty"`                                    // PB_OFFSET + MNM_EXCHANGE
	SequenceNumber *uint64  `protobuf:"varint,112002,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"` // PB_OFFSET + MNM_SEQUENCE_NUMBER
	Ssboe          *int32   `protobuf:"varint,150100,opt,name=ssboe" json:"ssboe,omitempty"`                                         // PB_OFFSET + MNM_SECONDS_SINCE_BOE
	Usecs          *int32   `protobuf:"varint,150101,opt,name=usecs" json:"usecs,omitempty"`                                         // PB_OFFSET + MNM_USECS
}

func (x *DepthByOrderEndEvent) Reset() {
	*x = DepthByOrderEndEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depth_by_order_end_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepthByOrderEndEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepthByOrderEndEvent) ProtoMessage() {}

func (x *DepthByOrderEndEvent) ProtoReflect() protoreflect.Message {
	mi := &file_depth_by_order_end_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepthByOrderEndEvent.ProtoReflect.Descriptor instead.
func (*DepthByOrderEndEvent) Descriptor() ([]byte, []int) {
	return file_depth_by_order_end_event_proto_rawDescGZIP(), []int{0}
}

func (x *DepthByOrderEndEvent) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *DepthByOrderEndEvent) GetSymbol() []string {
	if x != nil {
		return x.Symbol
	}
	return nil
}

func (x *DepthByOrderEndEvent) GetExchange() []string {
	if x != nil {
		return x.Exchange
	}
	return nil
}

func (x *DepthByOrderEndEvent) GetSequenceNumber() uint64 {
	if x != nil && x.SequenceNumber != nil {
		return *x.SequenceNumber
	}
	return 0
}

func (x *DepthByOrderEndEvent) GetSsboe() int32 {
	if x != nil && x.Ssboe != nil {
		return *x.Ssboe
	}
	return 0
}

func (x *DepthByOrderEndEvent) GetUsecs() int32 {
	if x != nil && x.Usecs != nil {
		return *x.Usecs
	}
	return 0
}

var File_depth_by_order_end_event_proto protoreflect.FileDescriptor

var file_depth_by_order_end_event_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x62, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xcc, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x74, 0x68, 0x42,
	0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6,
	0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x82, 0xeb, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0xd4, 0x94,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x16, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0xd5, 0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x63, 0x73,
}

var (
	file_depth_by_order_end_event_proto_rawDescOnce sync.Once
	file_depth_by_order_end_event_proto_rawDescData = file_depth_by_order_end_event_proto_rawDesc
)

func file_depth_by_order_end_event_proto_rawDescGZIP() []byte {
	file_depth_by_order_end_event_proto_rawDescOnce.Do(func() {
		file_depth_by_order_end_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_depth_by_order_end_event_proto_rawDescData)
	})
	return file_depth_by_order_end_event_proto_rawDescData
}

var file_depth_by_order_end_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_depth_by_order_end_event_proto_goTypes = []interface{}{
	(*DepthByOrderEndEvent)(nil), // 0: rti.DepthByOrderEndEvent
}
var file_depth_by_order_end_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_depth_by_order_end_event_proto_init() }
func file_depth_by_order_end_event_proto_init() {
	if File_depth_by_order_end_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_depth_by_order_end_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepthByOrderEndEvent); i {
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
			RawDescriptor: file_depth_by_order_end_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_depth_by_order_end_event_proto_goTypes,
		DependencyIndexes: file_depth_by_order_end_event_proto_depIdxs,
		MessageInfos:      file_depth_by_order_end_event_proto_msgTypes,
	}.Build()
	File_depth_by_order_end_event_proto = out.File
	file_depth_by_order_end_event_proto_rawDesc = nil
	file_depth_by_order_end_event_proto_goTypes = nil
	file_depth_by_order_end_event_proto_depIdxs = nil
}
