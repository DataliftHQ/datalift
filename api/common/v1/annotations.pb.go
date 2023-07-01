// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: common/v1/annotations.proto

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_common_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Reference)(nil),
		Field:         58901,
		Name:          "datalift.common.v1.reference",
		Tag:           "bytes,58901,opt,name=reference",
		Filename:      "common/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Identifier)(nil),
		Field:         58902,
		Name:          "datalift.common.v1.id",
		Tag:           "bytes,58902,opt,name=id",
		Filename:      "common/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         58903,
		Name:          "datalift.common.v1.redacted",
		Tag:           "varint,58903,opt,name=redacted",
		Filename:      "common/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         58901,
		Name:          "datalift.common.v1.log",
		Tag:           "varint,58901,opt,name=log",
		Filename:      "common/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Action)(nil),
		Field:         58901,
		Name:          "datalift.common.v1.action",
		Tag:           "bytes,58901,opt,name=action",
		Filename:      "common/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         58902,
		Name:          "datalift.common.v1.disable_audit",
		Tag:           "varint,58902,opt,name=disable_audit",
		Filename:      "common/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Use a random high number that won't conflict with annotations from other
	// libraries.
	//
	// optional datalift.common.v1.Reference reference = 58901;
	E_Reference = &file_common_v1_annotations_proto_extTypes[0]
	// optional datalift.common.v1.Identifier id = 58902;
	E_Id = &file_common_v1_annotations_proto_extTypes[1]
	// optional bool redacted = 58903;
	E_Redacted = &file_common_v1_annotations_proto_extTypes[2]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// Fields with the log option set to false will be cleared during auditing.
	// Defaults to true.
	//
	// optional bool log = 58901;
	E_Log = &file_common_v1_annotations_proto_extTypes[3]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// Use a random high number that won't conflict with annotations from other
	// libraries.
	//
	// optional datalift.common.v1.Action action = 58901;
	E_Action = &file_common_v1_annotations_proto_extTypes[4]
	// optional bool disable_audit = 58902;
	E_DisableAudit = &file_common_v1_annotations_proto_extTypes[5]
)

var File_common_v1_annotations_proto protoreflect.FileDescriptor

var file_common_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x5e, 0x0a, 0x09, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x95, 0xcc, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x3a, 0x51, 0x0a, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x96, 0xcc, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x3d,
	0x0a, 0x08, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x97, 0xcc, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x65, 0x64, 0x3a, 0x31, 0x0a,
	0x03, 0x6c, 0x6f, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x95, 0xcc, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6c, 0x6f, 0x67,
	0x3a, 0x54, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x95, 0xcc, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x45, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x96, 0xcc, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x42, 0xc4, 0x01,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x6f,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x6c, 0x69, 0x66, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44,
	0x43, 0x58, 0xaa, 0x02, 0x12, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69,
	0x66, 0x74, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x44,
	0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14,
	0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_common_v1_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil),  // 2: google.protobuf.MethodOptions
	(*Reference)(nil),                   // 3: datalift.common.v1.Reference
	(*Identifier)(nil),                  // 4: datalift.common.v1.Identifier
	(*Action)(nil),                      // 5: datalift.common.v1.Action
}
var file_common_v1_annotations_proto_depIdxs = []int32{
	0, // 0: datalift.common.v1.reference:extendee -> google.protobuf.MessageOptions
	0, // 1: datalift.common.v1.id:extendee -> google.protobuf.MessageOptions
	0, // 2: datalift.common.v1.redacted:extendee -> google.protobuf.MessageOptions
	1, // 3: datalift.common.v1.log:extendee -> google.protobuf.FieldOptions
	2, // 4: datalift.common.v1.action:extendee -> google.protobuf.MethodOptions
	2, // 5: datalift.common.v1.disable_audit:extendee -> google.protobuf.MethodOptions
	3, // 6: datalift.common.v1.reference:type_name -> datalift.common.v1.Reference
	4, // 7: datalift.common.v1.id:type_name -> datalift.common.v1.Identifier
	5, // 8: datalift.common.v1.action:type_name -> datalift.common.v1.Action
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	6, // [6:9] is the sub-list for extension type_name
	0, // [0:6] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_v1_annotations_proto_init() }
func file_common_v1_annotations_proto_init() {
	if File_common_v1_annotations_proto != nil {
		return
	}
	file_common_v1_schema_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_annotations_proto_goTypes,
		DependencyIndexes: file_common_v1_annotations_proto_depIdxs,
		ExtensionInfos:    file_common_v1_annotations_proto_extTypes,
	}.Build()
	File_common_v1_annotations_proto = out.File
	file_common_v1_annotations_proto_rawDesc = nil
	file_common_v1_annotations_proto_goTypes = nil
	file_common_v1_annotations_proto_depIdxs = nil
}
