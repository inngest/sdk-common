// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: probe.proto

package probe

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

type Probe int32

const (
	Probe_PROBE_NONE  Probe = 0
	Probe_PROBE_TRUST Probe = 1
)

// Enum value maps for Probe.
var (
	Probe_name = map[int32]string{
		0: "PROBE_NONE",
		1: "PROBE_TRUST",
	}
	Probe_value = map[string]int32{
		"PROBE_NONE":  0,
		"PROBE_TRUST": 1,
	}
)

func (x Probe) Enum() *Probe {
	p := new(Probe)
	*p = x
	return p
}

func (x Probe) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Probe) Descriptor() protoreflect.EnumDescriptor {
	return file_probe_proto_enumTypes[0].Descriptor()
}

func (Probe) Type() protoreflect.EnumType {
	return &file_probe_proto_enumTypes[0]
}

func (x Probe) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Probe.Descriptor instead.
func (Probe) EnumDescriptor() ([]byte, []int) {
	return file_probe_proto_rawDescGZIP(), []int{0}
}

var File_probe_proto protoreflect.FileDescriptor

var file_probe_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x28, 0x0a,
	0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x42, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x42, 0x45, 0x5f,
	0x54, 0x52, 0x55, 0x53, 0x54, 0x10, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x64,
	0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_probe_proto_rawDescOnce sync.Once
	file_probe_proto_rawDescData = file_probe_proto_rawDesc
)

func file_probe_proto_rawDescGZIP() []byte {
	file_probe_proto_rawDescOnce.Do(func() {
		file_probe_proto_rawDescData = protoimpl.X.CompressGZIP(file_probe_proto_rawDescData)
	})
	return file_probe_proto_rawDescData
}

var file_probe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_probe_proto_goTypes = []any{
	(Probe)(0), // 0: Probe
}
var file_probe_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_probe_proto_init() }
func file_probe_proto_init() {
	if File_probe_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_probe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_probe_proto_goTypes,
		DependencyIndexes: file_probe_proto_depIdxs,
		EnumInfos:         file_probe_proto_enumTypes,
	}.Build()
	File_probe_proto = out.File
	file_probe_proto_rawDesc = nil
	file_probe_proto_goTypes = nil
	file_probe_proto_depIdxs = nil
}
