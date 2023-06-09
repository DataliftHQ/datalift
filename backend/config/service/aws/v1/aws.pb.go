// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: config/service/aws/v1/aws.proto

package awsv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regions      []string      `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
	ClientConfig *ClientConfig `protobuf:"bytes,2,opt,name=client_config,json=clientConfig,proto3" json:"client_config,omitempty"`
	// The current account alias display name, if this is not set the default will be "default"
	// The account alias display name will be used when resolving resources.
	// EG: if this is set to "production" a resource would be referenced like so
	// "production/us-east-1/my-asg"
	PrimaryAccountAliasDisplayName string `protobuf:"bytes,4,opt,name=primary_account_alias_display_name,json=primaryAccountAliasDisplayName,proto3" json:"primary_account_alias_display_name,omitempty"`
	// If you are using an aws configuration file, this overrides the default profile that is loaded.
	// TODO: This is currently not implemented, but was created to prevent confusion between
	// primary_account_alias_display_name
	AwsConfigProfileName string `protobuf:"bytes,5,opt,name=aws_config_profile_name,json=awsConfigProfileName,proto3" json:"aws_config_profile_name,omitempty"`
	// A list of additional accounts you would like clutch to be able to operate in
	AdditionalAccounts []*AWSAccount `protobuf:"bytes,6,rep,name=additional_accounts,json=additionalAccounts,proto3" json:"additional_accounts,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_aws_v1_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_aws_v1_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_service_aws_v1_aws_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Config) GetClientConfig() *ClientConfig {
	if x != nil {
		return x.ClientConfig
	}
	return nil
}

func (x *Config) GetPrimaryAccountAliasDisplayName() string {
	if x != nil {
		return x.PrimaryAccountAliasDisplayName
	}
	return ""
}

func (x *Config) GetAwsConfigProfileName() string {
	if x != nil {
		return x.AwsConfigProfileName
	}
	return ""
}

func (x *Config) GetAdditionalAccounts() []*AWSAccount {
	if x != nil {
		return x.AdditionalAccounts
	}
	return nil
}

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If not set explicity, retries default to 0
	Retries int32 `protobuf:"varint,1,opt,name=retries,proto3" json:"retries,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_aws_v1_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_aws_v1_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_config_service_aws_v1_aws_proto_rawDescGZIP(), []int{1}
}

func (x *ClientConfig) GetRetries() int32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

type AWSAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account alias for this account
	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	// The account number for the aws account
	AccountNumber string `protobuf:"bytes,2,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	// The IAM role to use when performing operations against this account
	// NOTE: The role that Clutch assumes by default must have proper permissions
	// to assume the role below
	IamRole string `protobuf:"bytes,3,opt,name=iam_role,json=iamRole,proto3" json:"iam_role,omitempty"`
	// The list of regions you would like to operate in
	Regions []string `protobuf:"bytes,4,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *AWSAccount) Reset() {
	*x = AWSAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_aws_v1_aws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSAccount) ProtoMessage() {}

func (x *AWSAccount) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_aws_v1_aws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSAccount.ProtoReflect.Descriptor instead.
func (*AWSAccount) Descriptor() ([]byte, []int) {
	return file_config_service_aws_v1_aws_proto_rawDescGZIP(), []int{2}
}

func (x *AWSAccount) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *AWSAccount) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *AWSAccount) GetIamRole() string {
	if x != nil {
		return x.IamRole
	}
	return ""
}

func (x *AWSAccount) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

var File_config_service_aws_v1_aws_proto protoreflect.FileDescriptor

var file_config_service_aws_v1_aws_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4a, 0x0a, 0x22,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x61, 0x77, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x77, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x5b, 0x0a, 0x13, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57,
	0x53, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x31, 0x0a, 0x0c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0xa3, 0x01, 0x0a, 0x0a, 0x41, 0x57, 0x53, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x2e, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x08, 0x69, 0x61, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x07, 0x69, 0x61, 0x6d, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x8f, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x77,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x48, 0x51, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x77, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x44,
	0x43, 0x53, 0x41, 0xaa, 0x02, 0x1e, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x77,
	0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x5c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x41,
	0x77, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2a, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74,
	0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x41, 0x77, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x22, 0x44, 0x61, 0x74, 0x61, 0x6c, 0x69, 0x66, 0x74, 0x3a, 0x3a, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a,
	0x41, 0x77, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_service_aws_v1_aws_proto_rawDescOnce sync.Once
	file_config_service_aws_v1_aws_proto_rawDescData = file_config_service_aws_v1_aws_proto_rawDesc
)

func file_config_service_aws_v1_aws_proto_rawDescGZIP() []byte {
	file_config_service_aws_v1_aws_proto_rawDescOnce.Do(func() {
		file_config_service_aws_v1_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_aws_v1_aws_proto_rawDescData)
	})
	return file_config_service_aws_v1_aws_proto_rawDescData
}

var file_config_service_aws_v1_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_service_aws_v1_aws_proto_goTypes = []interface{}{
	(*Config)(nil),       // 0: datalift.config.service.aws.v1.Config
	(*ClientConfig)(nil), // 1: datalift.config.service.aws.v1.ClientConfig
	(*AWSAccount)(nil),   // 2: datalift.config.service.aws.v1.AWSAccount
}
var file_config_service_aws_v1_aws_proto_depIdxs = []int32{
	1, // 0: datalift.config.service.aws.v1.Config.client_config:type_name -> datalift.config.service.aws.v1.ClientConfig
	2, // 1: datalift.config.service.aws.v1.Config.additional_accounts:type_name -> datalift.config.service.aws.v1.AWSAccount
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_config_service_aws_v1_aws_proto_init() }
func file_config_service_aws_v1_aws_proto_init() {
	if File_config_service_aws_v1_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_service_aws_v1_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_service_aws_v1_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
		file_config_service_aws_v1_aws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSAccount); i {
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
			RawDescriptor: file_config_service_aws_v1_aws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_service_aws_v1_aws_proto_goTypes,
		DependencyIndexes: file_config_service_aws_v1_aws_proto_depIdxs,
		MessageInfos:      file_config_service_aws_v1_aws_proto_msgTypes,
	}.Build()
	File_config_service_aws_v1_aws_proto = out.File
	file_config_service_aws_v1_aws_proto_rawDesc = nil
	file_config_service_aws_v1_aws_proto_goTypes = nil
	file_config_service_aws_v1_aws_proto_depIdxs = nil
}
