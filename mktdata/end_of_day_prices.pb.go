// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: end_of_day_prices.proto

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
type EndOfDayPrices_PresenceBits int32

const (
	EndOfDayPrices_CLOSE                EndOfDayPrices_PresenceBits = 1
	EndOfDayPrices_SETTLEMENT           EndOfDayPrices_PresenceBits = 2
	EndOfDayPrices_PROJECTED_SETTLEMENT EndOfDayPrices_PresenceBits = 4
)

// Enum value maps for EndOfDayPrices_PresenceBits.
var (
	EndOfDayPrices_PresenceBits_name = map[int32]string{
		1: "CLOSE",
		2: "SETTLEMENT",
		4: "PROJECTED_SETTLEMENT",
	}
	EndOfDayPrices_PresenceBits_value = map[string]int32{
		"CLOSE":                1,
		"SETTLEMENT":           2,
		"PROJECTED_SETTLEMENT": 4,
	}
)

func (x EndOfDayPrices_PresenceBits) Enum() *EndOfDayPrices_PresenceBits {
	p := new(EndOfDayPrices_PresenceBits)
	*p = x
	return p
}

func (x EndOfDayPrices_PresenceBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndOfDayPrices_PresenceBits) Descriptor() protoreflect.EnumDescriptor {
	return file_end_of_day_prices_proto_enumTypes[0].Descriptor()
}

func (EndOfDayPrices_PresenceBits) Type() protoreflect.EnumType {
	return &file_end_of_day_prices_proto_enumTypes[0]
}

func (x EndOfDayPrices_PresenceBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EndOfDayPrices_PresenceBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EndOfDayPrices_PresenceBits(num)
	return nil
}

// Deprecated: Use EndOfDayPrices_PresenceBits.Descriptor instead.
func (EndOfDayPrices_PresenceBits) EnumDescriptor() ([]byte, []int) {
	return file_end_of_day_prices_proto_rawDescGZIP(), []int{0, 0}
}

type EndOfDayPrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId               *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                                              // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol                   *string  `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                                                         // PB_OFFSET + MNM_SYMBOL
	Exchange                 *string  `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                                                     // PB_OFFSET + MNM_EXCHANGE
	PresenceBits             *uint32  `protobuf:"varint,149138,opt,name=presence_bits,json=presenceBits" json:"presence_bits,omitempty"`                                        // PB_OFFSET + MNM_PRICING_INDICATOR
	ClearBits                *uint32  `protobuf:"varint,154571,opt,name=clear_bits,json=clearBits" json:"clear_bits,omitempty"`                                                 // PB_OFFSET + MNM_DISPLAY_INDICATOR
	IsSnapshot               *bool    `protobuf:"varint,110121,opt,name=is_snapshot,json=isSnapshot" json:"is_snapshot,omitempty"`                                              // PB_OFFSET + MNM_UPDATE_TYPE
	ClosePrice               *float64 `protobuf:"fixed64,100021,opt,name=close_price,json=closePrice" json:"close_price,omitempty"`                                             // PB_OFFSET + MNM_CLOSE_PRICE
	CloseDate                *string  `protobuf:"bytes,100079,opt,name=close_date,json=closeDate" json:"close_date,omitempty"`                                                  // PB_OFFSET + MNM_CLOSE_DATE
	SettlementPrice          *float64 `protobuf:"fixed64,100070,opt,name=settlement_price,json=settlementPrice" json:"settlement_price,omitempty"`                              // PB_OFFSET + MNM_SETTLEMENT_PRICE
	SettlementDate           *string  `protobuf:"bytes,154132,opt,name=settlement_date,json=settlementDate" json:"settlement_date,omitempty"`                                   // PB_OFFSET + MNM_SETTLEMENT_DATE
	ProjectedSettlementPrice *float64 `protobuf:"fixed64,155005,opt,name=projected_settlement_price,json=projectedSettlementPrice" json:"projected_settlement_price,omitempty"` // PB_OFFSET + MNM_PROJECTED_SETTLEMENT_PRICE
	Ssboe                    *int32   `protobuf:"varint,150100,opt,name=ssboe" json:"ssboe,omitempty"`                                                                          // PB_OFFSET + MNM_SECONDS_SINCE_BOE
	Usecs                    *int32   `protobuf:"varint,150101,opt,name=usecs" json:"usecs,omitempty"`                                                                          // PB_OFFSET + MNM_USECS
}

func (x *EndOfDayPrices) Reset() {
	*x = EndOfDayPrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_end_of_day_prices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndOfDayPrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndOfDayPrices) ProtoMessage() {}

func (x *EndOfDayPrices) ProtoReflect() protoreflect.Message {
	mi := &file_end_of_day_prices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndOfDayPrices.ProtoReflect.Descriptor instead.
func (*EndOfDayPrices) Descriptor() ([]byte, []int) {
	return file_end_of_day_prices_proto_rawDescGZIP(), []int{0}
}

func (x *EndOfDayPrices) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *EndOfDayPrices) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *EndOfDayPrices) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *EndOfDayPrices) GetPresenceBits() uint32 {
	if x != nil && x.PresenceBits != nil {
		return *x.PresenceBits
	}
	return 0
}

func (x *EndOfDayPrices) GetClearBits() uint32 {
	if x != nil && x.ClearBits != nil {
		return *x.ClearBits
	}
	return 0
}

func (x *EndOfDayPrices) GetIsSnapshot() bool {
	if x != nil && x.IsSnapshot != nil {
		return *x.IsSnapshot
	}
	return false
}

func (x *EndOfDayPrices) GetClosePrice() float64 {
	if x != nil && x.ClosePrice != nil {
		return *x.ClosePrice
	}
	return 0
}

func (x *EndOfDayPrices) GetCloseDate() string {
	if x != nil && x.CloseDate != nil {
		return *x.CloseDate
	}
	return ""
}

func (x *EndOfDayPrices) GetSettlementPrice() float64 {
	if x != nil && x.SettlementPrice != nil {
		return *x.SettlementPrice
	}
	return 0
}

func (x *EndOfDayPrices) GetSettlementDate() string {
	if x != nil && x.SettlementDate != nil {
		return *x.SettlementDate
	}
	return ""
}

func (x *EndOfDayPrices) GetProjectedSettlementPrice() float64 {
	if x != nil && x.ProjectedSettlementPrice != nil {
		return *x.ProjectedSettlementPrice
	}
	return 0
}

func (x *EndOfDayPrices) GetSsboe() int32 {
	if x != nil && x.Ssboe != nil {
		return *x.Ssboe
	}
	return 0
}

func (x *EndOfDayPrices) GetUsecs() int32 {
	if x != nil && x.Usecs != nil {
		return *x.Usecs
	}
	return 0
}

var File_end_of_day_prices_proto protoreflect.FileDescriptor

var file_end_of_day_prices_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x64, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xa7,
	0x04, 0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65,
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
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xb5, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xef, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x10, 0x73,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0xe6, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x73, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x94, 0xb4, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0xfd, 0xba, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x18, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x18, 0xd4, 0x94, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x73, 0x62, 0x6f, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x63, 0x73, 0x18, 0xd5, 0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73,
	0x65, 0x63, 0x73, 0x22, 0x43, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42,
	0x69, 0x74, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x18,
	0x0a, 0x14, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x54, 0x54,
	0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04,
}

var (
	file_end_of_day_prices_proto_rawDescOnce sync.Once
	file_end_of_day_prices_proto_rawDescData = file_end_of_day_prices_proto_rawDesc
)

func file_end_of_day_prices_proto_rawDescGZIP() []byte {
	file_end_of_day_prices_proto_rawDescOnce.Do(func() {
		file_end_of_day_prices_proto_rawDescData = protoimpl.X.CompressGZIP(file_end_of_day_prices_proto_rawDescData)
	})
	return file_end_of_day_prices_proto_rawDescData
}

var file_end_of_day_prices_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_end_of_day_prices_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_end_of_day_prices_proto_goTypes = []interface{}{
	(EndOfDayPrices_PresenceBits)(0), // 0: rti.EndOfDayPrices.PresenceBits
	(*EndOfDayPrices)(nil),           // 1: rti.EndOfDayPrices
}
var file_end_of_day_prices_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_end_of_day_prices_proto_init() }
func file_end_of_day_prices_proto_init() {
	if File_end_of_day_prices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_end_of_day_prices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndOfDayPrices); i {
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
			RawDescriptor: file_end_of_day_prices_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_end_of_day_prices_proto_goTypes,
		DependencyIndexes: file_end_of_day_prices_proto_depIdxs,
		EnumInfos:         file_end_of_day_prices_proto_enumTypes,
		MessageInfos:      file_end_of_day_prices_proto_msgTypes,
	}.Build()
	File_end_of_day_prices_proto = out.File
	file_end_of_day_prices_proto_rawDesc = nil
	file_end_of_day_prices_proto_goTypes = nil
	file_end_of_day_prices_proto_depIdxs = nil
}
