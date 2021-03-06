// Code generated by protoc-gen-go.
// source: nerdz.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	nerdz.proto

It has these top-level messages:
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Language int32

const (
	Language_INVALID    Language = 0
	Language_ENGLISH    Language = 1
	Language_ITALIAN    Language = 2
	Language_CROATIAN   Language = 3
	Language_GERMAN     Language = 4
	Language_PORTUGUESE Language = 5
	Language_ROMANIAN   Language = 6
)

var Language_name = map[int32]string{
	0: "INVALID",
	1: "ENGLISH",
	2: "ITALIAN",
	3: "CROATIAN",
	4: "GERMAN",
	5: "PORTUGUESE",
	6: "ROMANIAN",
}
var Language_value = map[string]int32{
	"INVALID":    0,
	"ENGLISH":    1,
	"ITALIAN":    2,
	"CROATIAN":   3,
	"GERMAN":     4,
	"PORTUGUESE": 5,
	"ROMANIAN":   6,
}

func (x Language) String() string {
	return proto1.EnumName(Language_name, int32(x))
}
func (Language) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto1.RegisterEnum("nerdz.Language", Language_name, Language_value)
}

func init() { proto1.RegisterFile("nerdz.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4b, 0x2d, 0x4a,
	0xa9, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xb4, 0x32, 0xb9, 0x38, 0x7c,
	0x12, 0xf3, 0xd2, 0x4b, 0x13, 0xd3, 0x53, 0x85, 0xb8, 0xb9, 0xd8, 0x3d, 0xfd, 0xc2, 0x1c, 0x7d,
	0x3c, 0x5d, 0x04, 0x18, 0x40, 0x1c, 0x57, 0x3f, 0x77, 0x1f, 0xcf, 0x60, 0x0f, 0x01, 0x46, 0xb0,
	0x4c, 0x88, 0xa3, 0x8f, 0xa7, 0xa3, 0x9f, 0x00, 0x93, 0x10, 0x0f, 0x17, 0x87, 0x73, 0x90, 0xbf,
	0x63, 0x08, 0x88, 0xc7, 0x2c, 0xc4, 0xc5, 0xc5, 0xe6, 0xee, 0x1a, 0xe4, 0xeb, 0xe8, 0x27, 0xc0,
	0x22, 0xc4, 0xc7, 0xc5, 0x15, 0xe0, 0x1f, 0x14, 0x12, 0xea, 0x1e, 0xea, 0x1a, 0xec, 0x2a, 0xc0,
	0x0a, 0x52, 0x19, 0xe4, 0xef, 0xeb, 0xe8, 0x07, 0x52, 0xc9, 0xe6, 0xc4, 0x1e, 0xc5, 0x0a, 0xb6,
	0x3a, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x54, 0x86, 0x91, 0x8a, 0x90,
	0x00, 0x00, 0x00,
}
