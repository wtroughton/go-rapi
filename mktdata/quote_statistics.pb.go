// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: quote_statistics.proto

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

// below enum is just for reference only, not used in this message
type QuoteStatistics_PresenceBits int32

const (
	QuoteStatistics_HIGHEST_BID QuoteStatistics_PresenceBits = 1
	QuoteStatistics_LOWEST_ASK  QuoteStatistics_PresenceBits = 2
)

// Enum value maps for QuoteStatistics_PresenceBits.
var (
	QuoteStatistics_PresenceBits_name = map[int32]string{
		1: "HIGHEST_BID",
		2: "LOWEST_ASK",
	}
	QuoteStatistics_PresenceBits_value = map[string]int32{
		"HIGHEST_BID": 1,
		"LOWEST_ASK":  2,
	}
)

func (x QuoteStatistics_PresenceBits) Enum() *QuoteStatistics_PresenceBits {
	p := new(QuoteStatistics_PresenceBits)
	*p = x
	return p
}

func (x QuoteStatistics_PresenceBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuoteStatistics_PresenceBits) Descriptor() protoreflect.EnumDescriptor {
	return file_quote_statistics_proto_enumTypes[0].Descriptor()
}

func (QuoteStatistics_PresenceBits) Type() protoreflect.EnumType {
	return &file_quote_statistics_proto_enumTypes[0]
}

func (x QuoteStatistics_PresenceBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *QuoteStatistics_PresenceBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = QuoteStatistics_PresenceBits(num)
	return nil
}

// Deprecated: Use QuoteStatistics_PresenceBits.Descriptor instead.
func (QuoteStatistics_PresenceBits) EnumDescriptor() ([]byte, []int) {
	return file_quote_statistics_proto_rawDescGZIP(), []int{0, 0}
}

type QuoteStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId      *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                   // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol          *string  `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                              // PB_OFFSET + MNM_SYMBOL
	Exchange        *string  `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                          // PB_OFFSET + MNM_EXCHANGE
	PresenceBits    *uint32  `protobuf:"varint,149138,opt,name=presence_bits,json=presenceBits" json:"presence_bits,omitempty"`             // PB_OFFSET + MNM_PRICING_INDICATOR
	ClearBits       *uint32  `protobuf:"varint,154571,opt,name=clear_bits,json=clearBits" json:"clear_bits,omitempty"`                      // PB_OFFSET + MNM_DISPLAY_INDICATOR
	IsSnapshot      *bool    `protobuf:"varint,110121,opt,name=is_snapshot,json=isSnapshot" json:"is_snapshot,omitempty"`                   // PB_OFFSET + MNM_UPDATE_TYPE
	HighestBidPrice *float64 `protobuf:"fixed64,154195,opt,name=highest_bid_price,json=highestBidPrice" json:"highest_bid_price,omitempty"` // PB_OFFSET + MNM_HIGHEST_BID_PRICE
	LowestAskPrice  *float64 `protobuf:"fixed64,154197,opt,name=lowest_ask_price,json=lowestAskPrice" json:"lowest_ask_price,omitempty"`    // PB_OFFSET + MNM_LOWEST_ASK_PRICE
	Ssboe           *int32   `protobuf:"varint,150100,opt,name=ssboe" json:"ssboe,omitempty"`                                               // PB_OFFSET + MNM_SECONDS_SINCE_BOE
	Usecs           *int32   `protobuf:"varint,150101,opt,name=usecs" json:"usecs,omitempty"`                                               // PB_OFFSET + MNM_USECS
}

func (x *QuoteStatistics) Reset() {
	*x = QuoteStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteStatistics) ProtoMessage() {}

func (x *QuoteStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_quote_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteStatistics.ProtoReflect.Descriptor instead.
func (*QuoteStatistics) Descriptor() ([]byte, []int) {
	return file_quote_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *QuoteStatistics) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *QuoteStatistics) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *QuoteStatistics) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *QuoteStatistics) GetPresenceBits() uint32 {
	if x != nil && x.PresenceBits != nil {
		return *x.PresenceBits
	}
	return 0
}

func (x *QuoteStatistics) GetClearBits() uint32 {
	if x != nil && x.ClearBits != nil {
		return *x.ClearBits
	}
	return 0
}

func (x *QuoteStatistics) GetIsSnapshot() bool {
	if x != nil && x.IsSnapshot != nil {
		return *x.IsSnapshot
	}
	return false
}

func (x *QuoteStatistics) GetHighestBidPrice() float64 {
	if x != nil && x.HighestBidPrice != nil {
		return *x.HighestBidPrice
	}
	return 0
}

func (x *QuoteStatistics) GetLowestAskPrice() float64 {
	if x != nil && x.LowestAskPrice != nil {
		return *x.LowestAskPrice
	}
	return 0
}

func (x *QuoteStatistics) GetSsboe() int32 {
	if x != nil && x.Ssboe != nil {
		return *x.Ssboe
	}
	return 0
}

func (x *QuoteStatistics) GetUsecs() int32 {
	if x != nil && x.Usecs != nil {
		return *x.Usecs
	}
	return 0
}

var File_quote_statistics_proto protoreflect.FileDescriptor

var file_quote_statistics_proto_rawDesc = []byte{
	0x0a, 0x16, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0x92, 0x03,
	0x0a, 0x0f, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94,
	0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c,
	0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0d,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x92, 0x8d,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x62, 0x69, 0x74,
	0x73, 0x18, 0xcb, 0xb7, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x72,
	0x42, 0x69, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x18, 0xa9, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x68, 0x69, 0x67, 0x68, 0x65,
	0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xd3, 0xb4, 0x09,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xd5, 0xb4, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0e, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x41, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0xd4, 0x94, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x63, 0x73, 0x18, 0xd5, 0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x65, 0x63,
	0x73, 0x22, 0x2f, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x69, 0x74,
	0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x49, 0x47, 0x48, 0x45, 0x53, 0x54, 0x5f, 0x42, 0x49, 0x44,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x57, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x53, 0x4b,
	0x10, 0x02,
}

var (
	file_quote_statistics_proto_rawDescOnce sync.Once
	file_quote_statistics_proto_rawDescData = file_quote_statistics_proto_rawDesc
)

func file_quote_statistics_proto_rawDescGZIP() []byte {
	file_quote_statistics_proto_rawDescOnce.Do(func() {
		file_quote_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_quote_statistics_proto_rawDescData)
	})
	return file_quote_statistics_proto_rawDescData
}

var file_quote_statistics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_quote_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_quote_statistics_proto_goTypes = []interface{}{
	(QuoteStatistics_PresenceBits)(0), // 0: rti.QuoteStatistics.PresenceBits
	(*QuoteStatistics)(nil),           // 1: rti.QuoteStatistics
}
var file_quote_statistics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_quote_statistics_proto_init() }
func file_quote_statistics_proto_init() {
	if File_quote_statistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quote_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteStatistics); i {
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
			RawDescriptor: file_quote_statistics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quote_statistics_proto_goTypes,
		DependencyIndexes: file_quote_statistics_proto_depIdxs,
		EnumInfos:         file_quote_statistics_proto_enumTypes,
		MessageInfos:      file_quote_statistics_proto_msgTypes,
	}.Build()
	File_quote_statistics_proto = out.File
	file_quote_statistics_proto_rawDesc = nil
	file_quote_statistics_proto_goTypes = nil
	file_quote_statistics_proto_depIdxs = nil
}
