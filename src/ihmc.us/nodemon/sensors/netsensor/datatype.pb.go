// Code generated by protoc-gen-go.
// source: datatype.proto
// DO NOT EDIT!

package netsensor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DataType int32

const (
	DataType_TRAFFIC  DataType = 0
	DataType_TOPOLOGY DataType = 1
	DataType_NETPROXY DataType = 2
)

var DataType_name = map[int32]string{
	0: "TRAFFIC",
	1: "TOPOLOGY",
	2: "NETPROXY",
}
var DataType_value = map[string]int32{
	"TRAFFIC":  0,
	"TOPOLOGY": 1,
	"NETPROXY": 2,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type StatType int32

const (
	StatType_TRAFFIC_AVERAGE StatType = 0
	StatType_PACKETS         StatType = 1
)

var StatType_name = map[int32]string{
	0: "TRAFFIC_AVERAGE",
	1: "PACKETS",
}
var StatType_value = map[string]int32{
	"TRAFFIC_AVERAGE": 0,
	"PACKETS":         1,
}

func (x StatType) String() string {
	return proto.EnumName(StatType_name, int32(x))
}
func (StatType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterEnum("netsensor.DataType", DataType_name, DataType_value)
	proto.RegisterEnum("netsensor.StatType", StatType_name, StatType_value)
}

func init() { proto.RegisterFile("datatype.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x49, 0x2c, 0x49,
	0x2c, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x4b, 0x2d, 0x29,
	0x4e, 0xcd, 0x2b, 0xce, 0x2f, 0xd2, 0x32, 0xe6, 0xe2, 0x70, 0x49, 0x2c, 0x49, 0x0c, 0xa9, 0x2c,
	0x48, 0x15, 0xe2, 0xe6, 0x62, 0x0f, 0x09, 0x72, 0x74, 0x73, 0xf3, 0x74, 0x16, 0x60, 0x10, 0xe2,
	0xe1, 0xe2, 0x08, 0xf1, 0x0f, 0xf0, 0xf7, 0xf1, 0x77, 0x8f, 0x14, 0x60, 0x04, 0xf1, 0xfc, 0x5c,
	0x43, 0x02, 0x82, 0xfc, 0x23, 0x22, 0x05, 0x98, 0xb4, 0x74, 0xb8, 0x38, 0x82, 0x4b, 0x12, 0x4b,
	0xc0, 0x9a, 0x84, 0xb9, 0xf8, 0xa1, 0x9a, 0xe2, 0x1d, 0xc3, 0x5c, 0x83, 0x1c, 0xdd, 0x5d, 0x05,
	0x18, 0x40, 0x26, 0x05, 0x38, 0x3a, 0x7b, 0xbb, 0x86, 0x04, 0x0b, 0x30, 0x3a, 0x39, 0x71, 0x29,
	0x96, 0x16, 0xeb, 0x65, 0x66, 0xe4, 0x26, 0xeb, 0xe5, 0xe5, 0xa7, 0xa4, 0xe6, 0xe6, 0xe7, 0xe9,
	0x41, 0x2c, 0x2f, 0xd6, 0x83, 0xbb, 0xc3, 0x89, 0x17, 0xe6, 0x8a, 0x00, 0x90, 0x0b, 0x03, 0x18,
	0xa3, 0x10, 0x6e, 0x5c, 0xc0, 0xc8, 0x98, 0xc4, 0x06, 0x76, 0xb8, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x09, 0x87, 0x8c, 0x58, 0xca, 0x00, 0x00, 0x00,
}
