// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/proto/server.proto

package game

import (
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

type RetrieveBalanceInput_Currency int32

const (
	RetrieveBalanceInput_CURRENCY_UNKNOWN RetrieveBalanceInput_Currency = 0
	RetrieveBalanceInput_CURRENCY_USD     RetrieveBalanceInput_Currency = 1
	RetrieveBalanceInput_CURRENCY_EUR     RetrieveBalanceInput_Currency = 2
)

// Enum value maps for RetrieveBalanceInput_Currency.
var (
	RetrieveBalanceInput_Currency_name = map[int32]string{
		0: "CURRENCY_UNKNOWN",
		1: "CURRENCY_USD",
		2: "CURRENCY_EUR",
	}
	RetrieveBalanceInput_Currency_value = map[string]int32{
		"CURRENCY_UNKNOWN": 0,
		"CURRENCY_USD":     1,
		"CURRENCY_EUR":     2,
	}
)

func (x RetrieveBalanceInput_Currency) Enum() *RetrieveBalanceInput_Currency {
	p := new(RetrieveBalanceInput_Currency)
	*p = x
	return p
}

func (x RetrieveBalanceInput_Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RetrieveBalanceInput_Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_server_proto_enumTypes[0].Descriptor()
}

func (RetrieveBalanceInput_Currency) Type() protoreflect.EnumType {
	return &file_api_proto_server_proto_enumTypes[0]
}

func (x RetrieveBalanceInput_Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RetrieveBalanceInput_Currency.Descriptor instead.
func (RetrieveBalanceInput_Currency) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_server_proto_rawDescGZIP(), []int{1, 0}
}

type RetrieveBalanceOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance float64 `protobuf:"fixed64,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *RetrieveBalanceOutput) Reset() {
	*x = RetrieveBalanceOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveBalanceOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveBalanceOutput) ProtoMessage() {}

func (x *RetrieveBalanceOutput) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveBalanceOutput.ProtoReflect.Descriptor instead.
func (*RetrieveBalanceOutput) Descriptor() ([]byte, []int) {
	return file_api_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *RetrieveBalanceOutput) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type RetrieveBalanceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string                        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Currency RetrieveBalanceInput_Currency `protobuf:"varint,2,opt,name=currency,proto3,enum=server.RetrieveBalanceInput_Currency" json:"currency,omitempty"`
}

func (x *RetrieveBalanceInput) Reset() {
	*x = RetrieveBalanceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveBalanceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveBalanceInput) ProtoMessage() {}

func (x *RetrieveBalanceInput) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveBalanceInput.ProtoReflect.Descriptor instead.
func (*RetrieveBalanceInput) Descriptor() ([]byte, []int) {
	return file_api_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *RetrieveBalanceInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RetrieveBalanceInput) GetCurrency() RetrieveBalanceInput_Currency {
	if x != nil {
		return x.Currency
	}
	return RetrieveBalanceInput_CURRENCY_UNKNOWN
}

var File_api_proto_server_proto protoreflect.FileDescriptor

var file_api_proto_server_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x22, 0x31, 0x0a, 0x15, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x44, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x55,
	0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x53, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x45, 0x55, 0x52, 0x10, 0x02, 0x32, 0x57,
	0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x50, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x67, 0x61,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_server_proto_rawDescOnce sync.Once
	file_api_proto_server_proto_rawDescData = file_api_proto_server_proto_rawDesc
)

func file_api_proto_server_proto_rawDescGZIP() []byte {
	file_api_proto_server_proto_rawDescOnce.Do(func() {
		file_api_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_server_proto_rawDescData)
	})
	return file_api_proto_server_proto_rawDescData
}

var file_api_proto_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_server_proto_goTypes = []interface{}{
	(RetrieveBalanceInput_Currency)(0), // 0: server.RetrieveBalanceInput.Currency
	(*RetrieveBalanceOutput)(nil),      // 1: server.RetrieveBalanceOutput
	(*RetrieveBalanceInput)(nil),       // 2: server.RetrieveBalanceInput
}
var file_api_proto_server_proto_depIdxs = []int32{
	0, // 0: server.RetrieveBalanceInput.currency:type_name -> server.RetrieveBalanceInput.Currency
	2, // 1: server.API.RetrieveBalance:input_type -> server.RetrieveBalanceInput
	1, // 2: server.API.RetrieveBalance:output_type -> server.RetrieveBalanceOutput
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_server_proto_init() }
func file_api_proto_server_proto_init() {
	if File_api_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveBalanceOutput); i {
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
		file_api_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveBalanceInput); i {
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
			RawDescriptor: file_api_proto_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_server_proto_goTypes,
		DependencyIndexes: file_api_proto_server_proto_depIdxs,
		EnumInfos:         file_api_proto_server_proto_enumTypes,
		MessageInfos:      file_api_proto_server_proto_msgTypes,
	}.Build()
	File_api_proto_server_proto = out.File
	file_api_proto_server_proto_rawDesc = nil
	file_api_proto_server_proto_goTypes = nil
	file_api_proto_server_proto_depIdxs = nil
}
