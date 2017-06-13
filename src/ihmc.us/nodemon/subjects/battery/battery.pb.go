// Code generated by protoc-gen-go.
// source: battery.proto
// DO NOT EDIT!

/*
Package battery is a generated protocol buffer package.

It is generated from these files:
	battery.proto

It has these top-level messages:
*/
package battery

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// all keys are lowercase, words separated by _
type Str int32

const (
	Str_sensor_ip Str = 0
)

var Str_name = map[int32]string{
	0: "sensor_ip",
}
var Str_value = map[string]int32{
	"sensor_ip": 0,
}

func (x Str) String() string {
	return proto.EnumName(Str_name, int32(x))
}
func (Str) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Double int32

const (
	Double_charge_percent Double = 0
)

var Double_name = map[int32]string{
	0: "charge_percent",
}
var Double_value = map[string]int32{
	"charge_percent": 0,
}

func (x Double) String() string {
	return proto.EnumName(Double_name, int32(x))
}
func (Double) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterEnum("geolocation.Str", Str_name, Str_value)
	proto.RegisterEnum("geolocation.Double", Double_name, Double_value)
}

func init() { proto.RegisterFile("battery.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0x2c, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f, 0xcd, 0xcf, 0xc9,
	0x4f, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x12, 0xe1, 0x62, 0x0e, 0x2e, 0x29, 0x12, 0xe2, 0xe5,
	0xe2, 0x2c, 0x4e, 0xcd, 0x2b, 0xce, 0x2f, 0x8a, 0xcf, 0x2c, 0x10, 0x60, 0xd0, 0x92, 0xe1, 0x62,
	0x73, 0xc9, 0x2f, 0x4d, 0xca, 0x49, 0x15, 0x12, 0xe2, 0xe2, 0x4b, 0xce, 0x48, 0x2c, 0x4a, 0x4f,
	0x8d, 0x2f, 0x48, 0x2d, 0x4a, 0x4e, 0xcd, 0x2b, 0x11, 0x60, 0x70, 0x32, 0xe7, 0x52, 0x28, 0x2d,
	0xd6, 0xcb, 0xcc, 0xc8, 0x4d, 0xd6, 0xcb, 0xcb, 0x4f, 0x49, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x2e,
	0x4d, 0xca, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x83, 0x5a, 0xe5, 0xc4, 0xee, 0x04, 0x61, 0x44, 0xb1,
	0x43, 0x45, 0x16, 0x30, 0x32, 0x26, 0xb1, 0x81, 0x1d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x70, 0x05, 0x08, 0xb6, 0x91, 0x00, 0x00, 0x00,
}
