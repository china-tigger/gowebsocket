// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/policy_violation_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible policy violation errors.
type PolicyViolationErrorEnum_PolicyViolationError int32

const (
	// Enum unspecified.
	PolicyViolationErrorEnum_UNSPECIFIED PolicyViolationErrorEnum_PolicyViolationError = 0
	// The received error code is not known in this version.
	PolicyViolationErrorEnum_UNKNOWN PolicyViolationErrorEnum_PolicyViolationError = 1
	// A policy was violated. See PolicyViolationDetails for more detail.
	PolicyViolationErrorEnum_POLICY_ERROR PolicyViolationErrorEnum_PolicyViolationError = 2
)

var PolicyViolationErrorEnum_PolicyViolationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "POLICY_ERROR",
}

var PolicyViolationErrorEnum_PolicyViolationError_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"POLICY_ERROR": 2,
}

func (x PolicyViolationErrorEnum_PolicyViolationError) String() string {
	return proto.EnumName(PolicyViolationErrorEnum_PolicyViolationError_name, int32(x))
}

func (PolicyViolationErrorEnum_PolicyViolationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_82022b324da9a65a, []int{0, 0}
}

// Container for enum describing possible policy violation errors.
type PolicyViolationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyViolationErrorEnum) Reset()         { *m = PolicyViolationErrorEnum{} }
func (m *PolicyViolationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyViolationErrorEnum) ProtoMessage()    {}
func (*PolicyViolationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_82022b324da9a65a, []int{0}
}

func (m *PolicyViolationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyViolationErrorEnum.Unmarshal(m, b)
}
func (m *PolicyViolationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyViolationErrorEnum.Marshal(b, m, deterministic)
}
func (m *PolicyViolationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyViolationErrorEnum.Merge(m, src)
}
func (m *PolicyViolationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyViolationErrorEnum.Size(m)
}
func (m *PolicyViolationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyViolationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyViolationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.PolicyViolationErrorEnum_PolicyViolationError", PolicyViolationErrorEnum_PolicyViolationError_name, PolicyViolationErrorEnum_PolicyViolationError_value)
	proto.RegisterType((*PolicyViolationErrorEnum)(nil), "google.ads.googleads.v1.errors.PolicyViolationErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/policy_violation_error.proto", fileDescriptor_82022b324da9a65a)
}

var fileDescriptor_82022b324da9a65a = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0x84, 0x40,
	0x14, 0x86, 0xd3, 0xa0, 0x60, 0x36, 0x48, 0xa4, 0x8b, 0x8a, 0xd8, 0x0b, 0x1f, 0x60, 0x06, 0xe9,
	0x6e, 0xf6, 0xca, 0xdd, 0x75, 0x17, 0x29, 0x54, 0x8c, 0x35, 0x0a, 0x41, 0xdc, 0x55, 0x06, 0xc1,
	0x9d, 0x23, 0x8e, 0x09, 0xbd, 0x4e, 0x97, 0x3d, 0x4a, 0x8f, 0xd2, 0x75, 0x0f, 0x10, 0xce, 0xa4,
	0x57, 0xdb, 0x5e, 0xf9, 0x73, 0xfc, 0xbf, 0xff, 0xfc, 0x73, 0xd0, 0x8c, 0x01, 0xb0, 0xaa, 0x20,
	0x59, 0x2e, 0x88, 0x92, 0xbd, 0xea, 0x6c, 0x52, 0x34, 0x0d, 0x34, 0x82, 0xd4, 0x50, 0x95, 0xbb,
	0xf7, 0xb4, 0x2b, 0xa1, 0xca, 0xda, 0x12, 0x78, 0x2a, 0xe7, 0xb8, 0x6e, 0xa0, 0x05, 0x73, 0xaa,
	0x08, 0x9c, 0xe5, 0x02, 0x8f, 0x30, 0xee, 0x6c, 0xac, 0xe0, 0xdb, 0xbb, 0x21, 0xbc, 0x2e, 0x49,
	0xc6, 0x39, 0xb4, 0x32, 0x42, 0x28, 0xda, 0xda, 0xa2, 0xeb, 0x50, 0xa6, 0xc7, 0x43, 0xb8, 0xdb,
	0x63, 0x2e, 0x7f, 0xdb, 0x5b, 0x2b, 0x74, 0x75, 0xe8, 0x9f, 0x79, 0x89, 0x26, 0x1b, 0xff, 0x29,
	0x74, 0x17, 0xde, 0xca, 0x73, 0x97, 0xc6, 0x89, 0x39, 0x41, 0xe7, 0x1b, 0xff, 0xc1, 0x0f, 0x9e,
	0x7d, 0x43, 0x33, 0x0d, 0x74, 0x11, 0x06, 0x8f, 0xde, 0xe2, 0x25, 0x75, 0xa3, 0x28, 0x88, 0x0c,
	0x7d, 0xfe, 0xa3, 0x21, 0x6b, 0x07, 0x7b, 0x7c, 0xbc, 0xe8, 0xfc, 0xe6, 0xd0, 0xb2, 0xb0, 0x6f,
	0x19, 0x6a, 0xaf, 0xcb, 0x3f, 0x98, 0x41, 0x95, 0x71, 0x86, 0xa1, 0x61, 0x84, 0x15, 0x5c, 0xbe,
	0x61, 0x38, 0x59, 0x5d, 0x8a, 0xff, 0x2e, 0x38, 0x53, 0x9f, 0x0f, 0xfd, 0x74, 0xed, 0x38, 0x9f,
	0xfa, 0x74, 0xad, 0xc2, 0x9c, 0x5c, 0x60, 0x25, 0x7b, 0x15, 0xdb, 0x58, 0xae, 0x14, 0x5f, 0x83,
	0x21, 0x71, 0x72, 0x91, 0x8c, 0x86, 0x24, 0xb6, 0x13, 0x65, 0xf8, 0xd6, 0x2d, 0x35, 0xa5, 0xd4,
	0xc9, 0x05, 0xa5, 0xa3, 0x85, 0xd2, 0xd8, 0xa6, 0x54, 0x99, 0xb6, 0x67, 0xb2, 0xdd, 0xfd, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x0b, 0x2a, 0xb5, 0xde, 0x01, 0x00, 0x00,
}
