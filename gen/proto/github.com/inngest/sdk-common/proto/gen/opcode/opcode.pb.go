// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: opcode.proto

package opcode

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

type Opcode int32

const (
	Opcode_OPCODE_NONE            Opcode = 0
	Opcode_OPCODE_STEP            Opcode = 1
	Opcode_OPCODE_STEP_RUN        Opcode = 2
	Opcode_OPCODE_STEP_ERROR      Opcode = 3
	Opcode_OPCODE_STEP_PLANNED    Opcode = 4
	Opcode_OPCODE_SLEEP           Opcode = 5
	Opcode_OPCODE_WAIT_FOR_EVENT  Opcode = 6
	Opcode_OPCODE_INVOKE_FUNCTION Opcode = 7
)

// Enum value maps for Opcode.
var (
	Opcode_name = map[int32]string{
		0: "OPCODE_NONE",
		1: "OPCODE_STEP",
		2: "OPCODE_STEP_RUN",
		3: "OPCODE_STEP_ERROR",
		4: "OPCODE_STEP_PLANNED",
		5: "OPCODE_SLEEP",
		6: "OPCODE_WAIT_FOR_EVENT",
		7: "OPCODE_INVOKE_FUNCTION",
	}
	Opcode_value = map[string]int32{
		"OPCODE_NONE":            0,
		"OPCODE_STEP":            1,
		"OPCODE_STEP_RUN":        2,
		"OPCODE_STEP_ERROR":      3,
		"OPCODE_STEP_PLANNED":    4,
		"OPCODE_SLEEP":           5,
		"OPCODE_WAIT_FOR_EVENT":  6,
		"OPCODE_INVOKE_FUNCTION": 7,
	}
)

func (x Opcode) Enum() *Opcode {
	p := new(Opcode)
	*p = x
	return p
}

func (x Opcode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Opcode) Descriptor() protoreflect.EnumDescriptor {
	return file_opcode_proto_enumTypes[0].Descriptor()
}

func (Opcode) Type() protoreflect.EnumType {
	return &file_opcode_proto_enumTypes[0]
}

func (x Opcode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Opcode.Descriptor instead.
func (Opcode) EnumDescriptor() ([]byte, []int) {
	return file_opcode_proto_rawDescGZIP(), []int{0}
}

var File_opcode_proto protoreflect.FileDescriptor

var file_opcode_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x73, 0x64, 0x6b, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x30, 0x2a, 0xb8, 0x01,
	0x0a, 0x06, 0x4f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x50,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x52, 0x55, 0x4e, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x4c, 0x45, 0x45, 0x50, 0x10,
	0x05, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x57, 0x41, 0x49, 0x54,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16,
	0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x46, 0x55,
	0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x73,
	0x64, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_opcode_proto_rawDescOnce sync.Once
	file_opcode_proto_rawDescData = file_opcode_proto_rawDesc
)

func file_opcode_proto_rawDescGZIP() []byte {
	file_opcode_proto_rawDescOnce.Do(func() {
		file_opcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_opcode_proto_rawDescData)
	})
	return file_opcode_proto_rawDescData
}

var file_opcode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_opcode_proto_goTypes = []any{
	(Opcode)(0), // 0: sdk.shared.v0.Opcode
}
var file_opcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_opcode_proto_init() }
func file_opcode_proto_init() {
	if File_opcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opcode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opcode_proto_goTypes,
		DependencyIndexes: file_opcode_proto_depIdxs,
		EnumInfos:         file_opcode_proto_enumTypes,
	}.Build()
	File_opcode_proto = out.File
	file_opcode_proto_rawDesc = nil
	file_opcode_proto_goTypes = nil
	file_opcode_proto_depIdxs = nil
}
