// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: query_param_key.proto

package query_param

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

type QueryParam int32

const (
	QueryParam_QUERY_PARAM_NONE      QueryParam = 0
	QueryParam_QUERY_PARAM_DEPLOY_ID QueryParam = 1
	QueryParam_QUERY_PARAM_FN_ID     QueryParam = 2
	QueryParam_QUERY_PARAM_PROBE     QueryParam = 3
	QueryParam_QUERY_PARAM_STEP_ID   QueryParam = 4
)

// Enum value maps for QueryParam.
var (
	QueryParam_name = map[int32]string{
		0: "QUERY_PARAM_NONE",
		1: "QUERY_PARAM_DEPLOY_ID",
		2: "QUERY_PARAM_FN_ID",
		3: "QUERY_PARAM_PROBE",
		4: "QUERY_PARAM_STEP_ID",
	}
	QueryParam_value = map[string]int32{
		"QUERY_PARAM_NONE":      0,
		"QUERY_PARAM_DEPLOY_ID": 1,
		"QUERY_PARAM_FN_ID":     2,
		"QUERY_PARAM_PROBE":     3,
		"QUERY_PARAM_STEP_ID":   4,
	}
)

func (x QueryParam) Enum() *QueryParam {
	p := new(QueryParam)
	*p = x
	return p
}

func (x QueryParam) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryParam) Descriptor() protoreflect.EnumDescriptor {
	return file_query_param_key_proto_enumTypes[0].Descriptor()
}

func (QueryParam) Type() protoreflect.EnumType {
	return &file_query_param_key_proto_enumTypes[0]
}

func (x QueryParam) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryParam.Descriptor instead.
func (QueryParam) EnumDescriptor() ([]byte, []int) {
	return file_query_param_key_proto_rawDescGZIP(), []int{0}
}

var File_query_param_key_proto protoreflect.FileDescriptor

var file_query_param_key_proto_rawDesc = []byte{
	0x0a, 0x15, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6b, 0x65,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x76, 0x30, 0x2a, 0x84, 0x01, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x46, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x15, 0x0a,
	0x11, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x50, 0x52, 0x4f,
	0x42, 0x45, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x44, 0x10, 0x04, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x2f, 0x73, 0x64, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_param_key_proto_rawDescOnce sync.Once
	file_query_param_key_proto_rawDescData = file_query_param_key_proto_rawDesc
)

func file_query_param_key_proto_rawDescGZIP() []byte {
	file_query_param_key_proto_rawDescOnce.Do(func() {
		file_query_param_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_param_key_proto_rawDescData)
	})
	return file_query_param_key_proto_rawDescData
}

var file_query_param_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_query_param_key_proto_goTypes = []any{
	(QueryParam)(0), // 0: sdk.shared.v0.QueryParam
}
var file_query_param_key_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_query_param_key_proto_init() }
func file_query_param_key_proto_init() {
	if File_query_param_key_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_query_param_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_query_param_key_proto_goTypes,
		DependencyIndexes: file_query_param_key_proto_depIdxs,
		EnumInfos:         file_query_param_key_proto_enumTypes,
	}.Build()
	File_query_param_key_proto = out.File
	file_query_param_key_proto_rawDesc = nil
	file_query_param_key_proto_goTypes = nil
	file_query_param_key_proto_depIdxs = nil
}
