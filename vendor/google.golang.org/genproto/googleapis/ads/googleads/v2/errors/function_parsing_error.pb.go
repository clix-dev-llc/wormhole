// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.2
// source: google/ads/googleads/v2/errors/function_parsing_error.proto

package errors

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Enum describing possible function parsing errors.
type FunctionParsingErrorEnum_FunctionParsingError int32

const (
	// Enum unspecified.
	FunctionParsingErrorEnum_UNSPECIFIED FunctionParsingErrorEnum_FunctionParsingError = 0
	// The received error code is not known in this version.
	FunctionParsingErrorEnum_UNKNOWN FunctionParsingErrorEnum_FunctionParsingError = 1
	// Unexpected end of function string.
	FunctionParsingErrorEnum_NO_MORE_INPUT FunctionParsingErrorEnum_FunctionParsingError = 2
	// Could not find an expected character.
	FunctionParsingErrorEnum_EXPECTED_CHARACTER FunctionParsingErrorEnum_FunctionParsingError = 3
	// Unexpected separator character.
	FunctionParsingErrorEnum_UNEXPECTED_SEPARATOR FunctionParsingErrorEnum_FunctionParsingError = 4
	// Unmatched left bracket or parenthesis.
	FunctionParsingErrorEnum_UNMATCHED_LEFT_BRACKET FunctionParsingErrorEnum_FunctionParsingError = 5
	// Unmatched right bracket or parenthesis.
	FunctionParsingErrorEnum_UNMATCHED_RIGHT_BRACKET FunctionParsingErrorEnum_FunctionParsingError = 6
	// Functions are nested too deeply.
	FunctionParsingErrorEnum_TOO_MANY_NESTED_FUNCTIONS FunctionParsingErrorEnum_FunctionParsingError = 7
	// Missing right-hand-side operand.
	FunctionParsingErrorEnum_MISSING_RIGHT_HAND_OPERAND FunctionParsingErrorEnum_FunctionParsingError = 8
	// Invalid operator/function name.
	FunctionParsingErrorEnum_INVALID_OPERATOR_NAME FunctionParsingErrorEnum_FunctionParsingError = 9
	// Feed attribute operand's argument is not an integer.
	FunctionParsingErrorEnum_FEED_ATTRIBUTE_OPERAND_ARGUMENT_NOT_INTEGER FunctionParsingErrorEnum_FunctionParsingError = 10
	// Missing function operands.
	FunctionParsingErrorEnum_NO_OPERANDS FunctionParsingErrorEnum_FunctionParsingError = 11
	// Function had too many operands.
	FunctionParsingErrorEnum_TOO_MANY_OPERANDS FunctionParsingErrorEnum_FunctionParsingError = 12
)

// Enum value maps for FunctionParsingErrorEnum_FunctionParsingError.
var (
	FunctionParsingErrorEnum_FunctionParsingError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "NO_MORE_INPUT",
		3:  "EXPECTED_CHARACTER",
		4:  "UNEXPECTED_SEPARATOR",
		5:  "UNMATCHED_LEFT_BRACKET",
		6:  "UNMATCHED_RIGHT_BRACKET",
		7:  "TOO_MANY_NESTED_FUNCTIONS",
		8:  "MISSING_RIGHT_HAND_OPERAND",
		9:  "INVALID_OPERATOR_NAME",
		10: "FEED_ATTRIBUTE_OPERAND_ARGUMENT_NOT_INTEGER",
		11: "NO_OPERANDS",
		12: "TOO_MANY_OPERANDS",
	}
	FunctionParsingErrorEnum_FunctionParsingError_value = map[string]int32{
		"UNSPECIFIED":                                 0,
		"UNKNOWN":                                     1,
		"NO_MORE_INPUT":                               2,
		"EXPECTED_CHARACTER":                          3,
		"UNEXPECTED_SEPARATOR":                        4,
		"UNMATCHED_LEFT_BRACKET":                      5,
		"UNMATCHED_RIGHT_BRACKET":                     6,
		"TOO_MANY_NESTED_FUNCTIONS":                   7,
		"MISSING_RIGHT_HAND_OPERAND":                  8,
		"INVALID_OPERATOR_NAME":                       9,
		"FEED_ATTRIBUTE_OPERAND_ARGUMENT_NOT_INTEGER": 10,
		"NO_OPERANDS":                                 11,
		"TOO_MANY_OPERANDS":                           12,
	}
)

func (x FunctionParsingErrorEnum_FunctionParsingError) Enum() *FunctionParsingErrorEnum_FunctionParsingError {
	p := new(FunctionParsingErrorEnum_FunctionParsingError)
	*p = x
	return p
}

func (x FunctionParsingErrorEnum_FunctionParsingError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FunctionParsingErrorEnum_FunctionParsingError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v2_errors_function_parsing_error_proto_enumTypes[0].Descriptor()
}

func (FunctionParsingErrorEnum_FunctionParsingError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v2_errors_function_parsing_error_proto_enumTypes[0]
}

func (x FunctionParsingErrorEnum_FunctionParsingError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FunctionParsingErrorEnum_FunctionParsingError.Descriptor instead.
func (FunctionParsingErrorEnum_FunctionParsingError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible function parsing errors.
type FunctionParsingErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FunctionParsingErrorEnum) Reset() {
	*x = FunctionParsingErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_errors_function_parsing_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionParsingErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionParsingErrorEnum) ProtoMessage() {}

func (x *FunctionParsingErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_errors_function_parsing_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionParsingErrorEnum.ProtoReflect.Descriptor instead.
func (*FunctionParsingErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v2_errors_function_parsing_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x18,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xe5, 0x02, 0x0a, 0x14, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x4d, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x43,
	0x48, 0x41, 0x52, 0x41, 0x43, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e,
	0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x50, 0x41, 0x52, 0x41, 0x54,
	0x4f, 0x52, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x45,
	0x44, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x5f, 0x42, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x05,
	0x12, 0x1b, 0x0a, 0x17, 0x55, 0x4e, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44, 0x5f, 0x52, 0x49,
	0x47, 0x48, 0x54, 0x5f, 0x42, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x06, 0x12, 0x1d, 0x0a,
	0x19, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44,
	0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x48, 0x41,
	0x4e, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x09, 0x12, 0x2f, 0x0a, 0x2b, 0x46, 0x45, 0x45, 0x44, 0x5f,
	0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e,
	0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x53, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x4f, 0x4f,
	0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x53, 0x10, 0x0c,
	0x42, 0xf4, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x19, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a,
	0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescData = file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDesc
)

func file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDescData
}

var file_google_ads_googleads_v2_errors_function_parsing_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v2_errors_function_parsing_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_errors_function_parsing_error_proto_goTypes = []interface{}{
	(FunctionParsingErrorEnum_FunctionParsingError)(0), // 0: google.ads.googleads.v2.errors.FunctionParsingErrorEnum.FunctionParsingError
	(*FunctionParsingErrorEnum)(nil),                   // 1: google.ads.googleads.v2.errors.FunctionParsingErrorEnum
}
var file_google_ads_googleads_v2_errors_function_parsing_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_errors_function_parsing_error_proto_init() }
func file_google_ads_googleads_v2_errors_function_parsing_error_proto_init() {
	if File_google_ads_googleads_v2_errors_function_parsing_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_errors_function_parsing_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionParsingErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_errors_function_parsing_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_errors_function_parsing_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v2_errors_function_parsing_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v2_errors_function_parsing_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_errors_function_parsing_error_proto = out.File
	file_google_ads_googleads_v2_errors_function_parsing_error_proto_rawDesc = nil
	file_google_ads_googleads_v2_errors_function_parsing_error_proto_goTypes = nil
	file_google_ads_googleads_v2_errors_function_parsing_error_proto_depIdxs = nil
}