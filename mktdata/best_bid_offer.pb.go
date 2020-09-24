// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: best_bid_offer.proto

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
type BestBidOffer_PresenceBits int32

const (
	BestBidOffer_BID        BestBidOffer_PresenceBits = 1
	BestBidOffer_ASK        BestBidOffer_PresenceBits = 2
	BestBidOffer_LEAN_PRICE BestBidOffer_PresenceBits = 3
)

// Enum value maps for BestBidOffer_PresenceBits.
var (
	BestBidOffer_PresenceBits_name = map[int32]string{
		1: "BID",
		2: "ASK",
		3: "LEAN_PRICE",
	}
	BestBidOffer_PresenceBits_value = map[string]int32{
		"BID":        1,
		"ASK":        2,
		"LEAN_PRICE": 3,
	}
)

func (x BestBidOffer_PresenceBits) Enum() *BestBidOffer_PresenceBits {
	p := new(BestBidOffer_PresenceBits)
	*p = x
	return p
}

func (x BestBidOffer_PresenceBits) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BestBidOffer_PresenceBits) Descriptor() protoreflect.EnumDescriptor {
	return file_best_bid_offer_proto_enumTypes[0].Descriptor()
}

func (BestBidOffer_PresenceBits) Type() protoreflect.EnumType {
	return &file_best_bid_offer_proto_enumTypes[0]
}

func (x BestBidOffer_PresenceBits) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BestBidOffer_PresenceBits) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BestBidOffer_PresenceBits(num)
	return nil
}

// Deprecated: Use BestBidOffer_PresenceBits.Descriptor instead.
func (BestBidOffer_PresenceBits) EnumDescriptor() ([]byte, []int) {
	return file_best_bid_offer_proto_rawDescGZIP(), []int{0, 0}
}

type BestBidOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId      *int32   `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`                  // PB_OFFSET + MNM_TEMPLATE_ID
	Symbol          *string  `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`                                             // PB_OFFSET + MNM_SYMBOL
	Exchange        *string  `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`                                         // PB_OFFSET + MNM_EXCHANGE
	PresenceBits    *uint32  `protobuf:"varint,149138,opt,name=presence_bits,json=presenceBits" json:"presence_bits,omitempty"`            // PB_OFFSET + MNM_PRICING_INDICATOR
	ClearBits       *uint32  `protobuf:"varint,154571,opt,name=clear_bits,json=clearBits" json:"clear_bits,omitempty"`                     // PB_OFFSET + MNM_DISPLAY_INDICATOR
	IsSnapshot      *bool    `protobuf:"varint,110121,opt,name=is_snapshot,json=isSnapshot" json:"is_snapshot,omitempty"`                  // PB_OFFSET + MNM_UPDATE_TYPE
	BidPrice        *float64 `protobuf:"fixed64,100022,opt,name=bid_price,json=bidPrice" json:"bid_price,omitempty"`                       // PB_OFFSET + MNM_BID_PRICE
	BidSize         *int32   `protobuf:"varint,100030,opt,name=bid_size,json=bidSize" json:"bid_size,omitempty"`                           // PB_OFFSET + MNM_BID_SIZE
	BidOrders       *int32   `protobuf:"varint,154403,opt,name=bid_orders,json=bidOrders" json:"bid_orders,omitempty"`                     // PB_OFFSET + MNM_BID_NO_OF_ORDERS
	BidImplicitSize *int32   `protobuf:"varint,154867,opt,name=bid_implicit_size,json=bidImplicitSize" json:"bid_implicit_size,omitempty"` // PB_OFFSET + MNM_BID_IMPLICIT_SIZE
	AskPrice        *float64 `protobuf:"fixed64,100025,opt,name=ask_price,json=askPrice" json:"ask_price,omitempty"`                       // PB_OFFSET + MNM_ASK_PRICE
	AskSize         *int32   `protobuf:"varint,100031,opt,name=ask_size,json=askSize" json:"ask_size,omitempty"`                           // PB_OFFSET + MNM_ASK_SIZE
	AskOrders       *int32   `protobuf:"varint,154404,opt,name=ask_orders,json=askOrders" json:"ask_orders,omitempty"`                     // PB_OFFSET + MNM_ASK_NO_OF_ORDERS
	AskImplicitSize *int32   `protobuf:"varint,154868,opt,name=ask_implicit_size,json=askImplicitSize" json:"ask_implicit_size,omitempty"` // PB_OFFSET + MNM_ASK_IMPLICIT_SIZE
	LeanPrice       *float64 `protobuf:"fixed64,154909,opt,name=lean_price,json=leanPrice" json:"lean_price,omitempty"`                    // PB_OFFSET + MNM_LEAN_PRICE
	Ssboe           *int32   `protobuf:"varint,150100,opt,name=ssboe" json:"ssboe,omitempty"`                                              // PB_OFFSET + MNM_SECONDS_SINCE_BOE
	Usecs           *int32   `protobuf:"varint,150101,opt,name=usecs" json:"usecs,omitempty"`                                              // PB_OFFSET + MNM_USECS
}

func (x *BestBidOffer) Reset() {
	*x = BestBidOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_best_bid_offer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BestBidOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BestBidOffer) ProtoMessage() {}

func (x *BestBidOffer) ProtoReflect() protoreflect.Message {
	mi := &file_best_bid_offer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BestBidOffer.ProtoReflect.Descriptor instead.
func (*BestBidOffer) Descriptor() ([]byte, []int) {
	return file_best_bid_offer_proto_rawDescGZIP(), []int{0}
}

func (x *BestBidOffer) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *BestBidOffer) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *BestBidOffer) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *BestBidOffer) GetPresenceBits() uint32 {
	if x != nil && x.PresenceBits != nil {
		return *x.PresenceBits
	}
	return 0
}

func (x *BestBidOffer) GetClearBits() uint32 {
	if x != nil && x.ClearBits != nil {
		return *x.ClearBits
	}
	return 0
}

func (x *BestBidOffer) GetIsSnapshot() bool {
	if x != nil && x.IsSnapshot != nil {
		return *x.IsSnapshot
	}
	return false
}

func (x *BestBidOffer) GetBidPrice() float64 {
	if x != nil && x.BidPrice != nil {
		return *x.BidPrice
	}
	return 0
}

func (x *BestBidOffer) GetBidSize() int32 {
	if x != nil && x.BidSize != nil {
		return *x.BidSize
	}
	return 0
}

func (x *BestBidOffer) GetBidOrders() int32 {
	if x != nil && x.BidOrders != nil {
		return *x.BidOrders
	}
	return 0
}

func (x *BestBidOffer) GetBidImplicitSize() int32 {
	if x != nil && x.BidImplicitSize != nil {
		return *x.BidImplicitSize
	}
	return 0
}

func (x *BestBidOffer) GetAskPrice() float64 {
	if x != nil && x.AskPrice != nil {
		return *x.AskPrice
	}
	return 0
}

func (x *BestBidOffer) GetAskSize() int32 {
	if x != nil && x.AskSize != nil {
		return *x.AskSize
	}
	return 0
}

func (x *BestBidOffer) GetAskOrders() int32 {
	if x != nil && x.AskOrders != nil {
		return *x.AskOrders
	}
	return 0
}

func (x *BestBidOffer) GetAskImplicitSize() int32 {
	if x != nil && x.AskImplicitSize != nil {
		return *x.AskImplicitSize
	}
	return 0
}

func (x *BestBidOffer) GetLeanPrice() float64 {
	if x != nil && x.LeanPrice != nil {
		return *x.LeanPrice
	}
	return 0
}

func (x *BestBidOffer) GetSsboe() int32 {
	if x != nil && x.Ssboe != nil {
		return *x.Ssboe
	}
	return 0
}

func (x *BestBidOffer) GetUsecs() int32 {
	if x != nil && x.Usecs != nil {
		return *x.Usecs
	}
	return 0
}

var File_best_bid_offer_proto protoreflect.FileDescriptor

var file_best_bid_offer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69, 0x22, 0xed, 0x04, 0x0a, 0x0c,
	0x42, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x92, 0x8d, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x69, 0x74, 0x73, 0x12, 0x1f,
	0x0a, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0xcb, 0xb7, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x42, 0x69, 0x74, 0x73, 0x12,
	0x21, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0xa9,
	0xdc, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0xb6, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0xbe, 0x8d,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x69, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f,
	0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0xa3, 0xb6, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x2c, 0x0a, 0x11, 0x62, 0x69, 0x64, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0xf3, 0xb9, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x62, 0x69,
	0x64, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xb9, 0x8d, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x08,
	0x61, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0xbf, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x61, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x73, 0x6b,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0xa4, 0xb6, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x61, 0x73, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x73,
	0x6b, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0xf4, 0xb9, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x73, 0x6b, 0x49, 0x6d, 0x70, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x6e,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x9d, 0xba, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x6c, 0x65, 0x61, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x73, 0x73, 0x62,
	0x6f, 0x65, 0x18, 0xd4, 0x94, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x73, 0x62, 0x6f,
	0x65, 0x12, 0x16, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x63, 0x73, 0x18, 0xd5, 0x94, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x65, 0x63, 0x73, 0x22, 0x30, 0x0a, 0x0c, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x69, 0x74, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49, 0x44,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x4b, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4c,
	0x45, 0x41, 0x4e, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x03,
}

var (
	file_best_bid_offer_proto_rawDescOnce sync.Once
	file_best_bid_offer_proto_rawDescData = file_best_bid_offer_proto_rawDesc
)

func file_best_bid_offer_proto_rawDescGZIP() []byte {
	file_best_bid_offer_proto_rawDescOnce.Do(func() {
		file_best_bid_offer_proto_rawDescData = protoimpl.X.CompressGZIP(file_best_bid_offer_proto_rawDescData)
	})
	return file_best_bid_offer_proto_rawDescData
}

var file_best_bid_offer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_best_bid_offer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_best_bid_offer_proto_goTypes = []interface{}{
	(BestBidOffer_PresenceBits)(0), // 0: rti.BestBidOffer.PresenceBits
	(*BestBidOffer)(nil),           // 1: rti.BestBidOffer
}
var file_best_bid_offer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_best_bid_offer_proto_init() }
func file_best_bid_offer_proto_init() {
	if File_best_bid_offer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_best_bid_offer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BestBidOffer); i {
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
			RawDescriptor: file_best_bid_offer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_best_bid_offer_proto_goTypes,
		DependencyIndexes: file_best_bid_offer_proto_depIdxs,
		EnumInfos:         file_best_bid_offer_proto_enumTypes,
		MessageInfos:      file_best_bid_offer_proto_msgTypes,
	}.Build()
	File_best_bid_offer_proto = out.File
	file_best_bid_offer_proto_rawDesc = nil
	file_best_bid_offer_proto_goTypes = nil
	file_best_bid_offer_proto_depIdxs = nil
}
