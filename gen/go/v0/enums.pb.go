// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: enums.proto

package sdk

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Header int32

const (
	Header_HEADER_NONE                           Header = 0
	Header_HEADER_AUTHORIZATION                  Header = 1
	Header_HEADER_CONTENT_TYPE                   Header = 2
	Header_HEADER_RETRY_AFTER                    Header = 3
	Header_HEADER_SERVER_TIMING                  Header = 4
	Header_HEADER_USER_AGENT                     Header = 5
	Header_HEADER_X_INNGEST_ENV                  Header = 6
	Header_HEADER_X_INNGEST_EXPECTED_SERVER_KIND Header = 7
	Header_HEADER_X_INNGEST_FRAMEWORK            Header = 8
	Header_HEADER_X_INNGEST_NO_RETRY             Header = 9
	Header_HEADER_X_INNGEST_REQ_VERSION          Header = 10
	Header_HEADER_X_INNGEST_SDK                  Header = 11
	Header_HEADER_X_INNGEST_SERVER_KIND          Header = 12
	Header_HEADER_X_INNGEST_SIGNATURE            Header = 13
	Header_HEADER_X_INNGEST_SYNC_KIND            Header = 14
)

// Enum value maps for Header.
var (
	Header_name = map[int32]string{
		0:  "HEADER_NONE",
		1:  "HEADER_AUTHORIZATION",
		2:  "HEADER_CONTENT_TYPE",
		3:  "HEADER_RETRY_AFTER",
		4:  "HEADER_SERVER_TIMING",
		5:  "HEADER_USER_AGENT",
		6:  "HEADER_X_INNGEST_ENV",
		7:  "HEADER_X_INNGEST_EXPECTED_SERVER_KIND",
		8:  "HEADER_X_INNGEST_FRAMEWORK",
		9:  "HEADER_X_INNGEST_NO_RETRY",
		10: "HEADER_X_INNGEST_REQ_VERSION",
		11: "HEADER_X_INNGEST_SDK",
		12: "HEADER_X_INNGEST_SERVER_KIND",
		13: "HEADER_X_INNGEST_SIGNATURE",
		14: "HEADER_X_INNGEST_SYNC_KIND",
	}
	Header_value = map[string]int32{
		"HEADER_NONE":                           0,
		"HEADER_AUTHORIZATION":                  1,
		"HEADER_CONTENT_TYPE":                   2,
		"HEADER_RETRY_AFTER":                    3,
		"HEADER_SERVER_TIMING":                  4,
		"HEADER_USER_AGENT":                     5,
		"HEADER_X_INNGEST_ENV":                  6,
		"HEADER_X_INNGEST_EXPECTED_SERVER_KIND": 7,
		"HEADER_X_INNGEST_FRAMEWORK":            8,
		"HEADER_X_INNGEST_NO_RETRY":             9,
		"HEADER_X_INNGEST_REQ_VERSION":          10,
		"HEADER_X_INNGEST_SDK":                  11,
		"HEADER_X_INNGEST_SERVER_KIND":          12,
		"HEADER_X_INNGEST_SIGNATURE":            13,
		"HEADER_X_INNGEST_SYNC_KIND":            14,
	}
)

func (x Header) Enum() *Header {
	p := new(Header)
	*p = x
	return p
}

func (x Header) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Header) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[0].Descriptor()
}

func (Header) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[0]
}

func (x Header) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Header.Descriptor instead.
func (Header) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

type Mode int32

const (
	Mode_MODE_NONE  Mode = 0
	Mode_MODE_CLOUD Mode = 1
	Mode_MODE_DEV   Mode = 2
)

// Enum value maps for Mode.
var (
	Mode_name = map[int32]string{
		0: "MODE_NONE",
		1: "MODE_CLOUD",
		2: "MODE_DEV",
	}
	Mode_value = map[string]int32{
		"MODE_NONE":  0,
		"MODE_CLOUD": 1,
		"MODE_DEV":   2,
	}
)

func (x Mode) Enum() *Mode {
	p := new(Mode)
	*p = x
	return p
}

func (x Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[1].Descriptor()
}

func (Mode) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[1]
}

func (x Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mode.Descriptor instead.
func (Mode) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{1}
}

var file_enums_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         10000,
		Name:          "sdk.shared.str",
		Tag:           "bytes,10000,opt,name=str",
		Filename:      "enums.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string str = 10000;
	E_Str = &file_enums_proto_extTypes[0]
)

var File_enums_proto protoreflect.FileDescriptor

var file_enums_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73,
	0x64, 0x6b, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xfd, 0x05, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x0b, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x1a, 0x08, 0x82, 0xf1, 0x04, 0x04, 0x6e, 0x6f, 0x6e,
	0x65, 0x12, 0x2b, 0x0a, 0x14, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x55, 0x54, 0x48,
	0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x1a, 0x11, 0x82, 0xf1, 0x04,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29,
	0x0a, 0x13, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x02, 0x1a, 0x10, 0x82, 0xf1, 0x04, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x12, 0x48, 0x45, 0x41,
	0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x10,
	0x03, 0x1a, 0x0f, 0x82, 0xf1, 0x04, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x2d, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x14, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x1a, 0x11, 0x82, 0xf1,
	0x04, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x12,
	0x25, 0x0a, 0x11, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41,
	0x47, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x1a, 0x0e, 0x82, 0xf1, 0x04, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x14, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52,
	0x5f, 0x58, 0x5f, 0x49, 0x4e, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x56, 0x10, 0x06,
	0x1a, 0x11, 0x82, 0xf1, 0x04, 0x0d, 0x78, 0x2d, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2d,
	0x65, 0x6e, 0x76, 0x12, 0x4d, 0x0a, 0x25, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x58, 0x5f,
	0x49, 0x4e, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0x07, 0x1a, 0x22,
	0x82, 0xf1, 0x04, 0x1e, 0x78, 0x2d, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2d, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x37, 0x0a, 0x1a, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x58, 0x5f, 0x49,
	0x4e, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x57, 0x4f, 0x52, 0x4b,
	0x10, 0x08, 0x1a, 0x17, 0x82, 0xf1, 0x04, 0x13, 0x78, 0x2d, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x35, 0x0a, 0x19, 0x48,
	0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x58, 0x5f, 0x49, 0x4e, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x5f,
	0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x10, 0x09, 0x1a, 0x16, 0x82, 0xf1, 0x04, 0x12,
	0x78, 0x2d, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2d, 0x6e, 0x6f, 0x2d, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x12, 0x3b, 0x0a, 0x1c, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x58, 0x5f, 0x49,
	0x4e, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x10, 0x0a, 0x1a, 0x19, 0x82, 0xf1, 0x04, 0x15, 0x78, 0x2d, 0x69, 0x6e, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x2d, 0x72, 0x65, 0x71, 0x2d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2b, 0x0a, 0x14, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x58, 0x5f, 0x49, 0x4e, 0x4e, 0x47,
	0x45, 0x53, 0x54, 0x5f, 0x53, 0x44, 0x4b, 0x10, 0x0b, 0x1a, 0x11, 0x82, 0xf1, 0x04, 0x0d, 0x78,
	0x2d, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x12, 0x3b, 0x0a, 0x1c,
	0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x58, 0x5f, 0x49, 0x4e, 0x4e, 0x47, 0x45, 0x53, 0x54,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0x0c, 0x1a, 0x19,
	0x82, 0xf1, 0x04, 0x15, 0x78, 0x2d, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2d, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x37, 0x0a, 0x1a, 0x48, 0x45, 0x41,
	0x44, 0x45, 0x52, 0x5f, 0x58, 0x5f, 0x49, 0x4e, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x49,
	0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x0d, 0x1a, 0x17, 0x82, 0xf1, 0x04, 0x13, 0x78,
	0x2d, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x37, 0x0a, 0x1a, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x58, 0x5f, 0x49,
	0x4e, 0x4e, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x10, 0x0e, 0x1a, 0x17, 0x82, 0xf1, 0x04, 0x13, 0x78, 0x2d, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x6b, 0x69, 0x6e, 0x64, 0x2a, 0x51, 0x0a, 0x04, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x09, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x1a, 0x08, 0x82, 0xf1, 0x04, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x0a,
	0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x10, 0x01, 0x1a, 0x09, 0x82, 0xf1,
	0x04, 0x05, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x15, 0x0a, 0x08, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x44, 0x45, 0x56, 0x10, 0x02, 0x1a, 0x07, 0x82, 0xf1, 0x04, 0x03, 0x64, 0x65, 0x76, 0x3a, 0x37,
	0x0a, 0x03, 0x73, 0x74, 0x72, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x90, 0x4e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x74, 0x72, 0x88, 0x01, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x64,
	0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x30, 0x2f,
	0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_proto_rawDescOnce sync.Once
	file_enums_proto_rawDescData = file_enums_proto_rawDesc
)

func file_enums_proto_rawDescGZIP() []byte {
	file_enums_proto_rawDescOnce.Do(func() {
		file_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_proto_rawDescData)
	})
	return file_enums_proto_rawDescData
}

var file_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_enums_proto_goTypes = []any{
	(Header)(0),                           // 0: sdk.shared.Header
	(Mode)(0),                             // 1: sdk.shared.Mode
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
}
var file_enums_proto_depIdxs = []int32{
	2, // 0: sdk.shared.str:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_proto_init() }
func file_enums_proto_init() {
	if File_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_enums_proto_goTypes,
		DependencyIndexes: file_enums_proto_depIdxs,
		EnumInfos:         file_enums_proto_enumTypes,
		ExtensionInfos:    file_enums_proto_extTypes,
	}.Build()
	File_enums_proto = out.File
	file_enums_proto_rawDesc = nil
	file_enums_proto_goTypes = nil
	file_enums_proto_depIdxs = nil
}
