// Code generated by protoc-gen-go.
// source: subject.proto
// DO NOT EDIT!

package measure

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Subject int32

const (
	Subject_host           Subject = 0
	Subject_network        Subject = 1
	Subject_traffic        Subject = 2
	Subject_cpu            Subject = 3
	Subject_memory         Subject = 4
	Subject_os             Subject = 5
	Subject_network_health Subject = 6
	Subject_disservice     Subject = 7
	Subject_mockets        Subject = 8
)

var Subject_name = map[int32]string{
	0: "host",
	1: "network",
	2: "traffic",
	3: "cpu",
	4: "memory",
	5: "os",
	6: "network_health",
	7: "disservice",
	8: "mockets",
}
var Subject_value = map[string]int32{
	"host":           0,
	"network":        1,
	"traffic":        2,
	"cpu":            3,
	"memory":         4,
	"os":             5,
	"network_health": 6,
	"disservice":     7,
	"mockets":        8,
}

func (x Subject) String() string {
	return proto.EnumName(Subject_name, int32(x))
}
func (Subject) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterEnum("measure.Subject", Subject_name, Subject_value)
}

func init() { proto.RegisterFile("subject.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x34, 0xcf, 0xb1, 0x4e, 0x03, 0x31,
	0x0c, 0x06, 0x60, 0xae, 0x2d, 0x49, 0x65, 0xa0, 0xb2, 0x3c, 0xf3, 0x04, 0x0c, 0x59, 0xd8, 0x18,
	0xfb, 0x04, 0x95, 0xd8, 0x58, 0x50, 0x9a, 0xb8, 0x4a, 0x28, 0x39, 0x9f, 0xe2, 0x04, 0x84, 0x78,
	0x19, 0x1e, 0x15, 0x1d, 0x5c, 0x47, 0xcb, 0xfa, 0x7f, 0x7d, 0x3f, 0xdc, 0x69, 0x3f, 0xbe, 0x71,
	0x68, 0x6e, 0xaa, 0xd2, 0x84, 0x6c, 0x61, 0xaf, 0xbd, 0xf2, 0xc3, 0x37, 0xd8, 0xe7, 0xff, 0x0f,
	0x6d, 0x61, 0x93, 0x44, 0x1b, 0x5e, 0xd1, 0x0d, 0xd8, 0x91, 0xdb, 0xa7, 0xd4, 0x33, 0x0e, 0xf3,
	0xd1, 0xaa, 0x3f, 0x9d, 0x72, 0xc0, 0x15, 0x59, 0x58, 0x87, 0xa9, 0xe3, 0x9a, 0x00, 0x4c, 0xe1,
	0x22, 0xf5, 0x0b, 0x37, 0x64, 0x60, 0x25, 0x8a, 0xd7, 0x44, 0xb0, 0x5b, 0x62, 0xaf, 0x89, 0xfd,
	0x7b, 0x4b, 0x68, 0x68, 0x07, 0x10, 0xb3, 0x2a, 0xd7, 0x8f, 0x1c, 0x18, 0xed, 0xdc, 0x56, 0x24,
	0x9c, 0xb9, 0x29, 0x6e, 0xf7, 0x4f, 0x70, 0xdf, 0xd5, 0xe5, 0x54, 0x82, 0x8b, 0xd1, 0x17, 0x37,
	0x4a, 0xe4, 0x22, 0xa3, 0x5b, 0x70, 0xfb, 0xdb, 0x85, 0x76, 0x98, 0xcd, 0x87, 0xe1, 0xe5, 0xa2,
	0xfe, 0x19, 0x86, 0xa3, 0xf9, 0x1b, 0xf2, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xa5, 0x45,
	0xe6, 0xd9, 0x00, 0x00, 0x00,
}