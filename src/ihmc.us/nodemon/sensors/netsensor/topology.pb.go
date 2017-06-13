// Code generated by protoc-gen-go.
// source: topology.proto
// DO NOT EDIT!

package netsensor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Topology struct {
	NetworkInfo *NetworkInfo `protobuf:"bytes,1,opt,name=networkInfo" json:"networkInfo,omitempty"`
	Internals   []*Host      `protobuf:"bytes,2,rep,name=internals" json:"internals,omitempty"`
	LocalGws    []*Host      `protobuf:"bytes,3,rep,name=localGws" json:"localGws,omitempty"`
}

func (m *Topology) Reset()                    { *m = Topology{} }
func (m *Topology) String() string            { return proto.CompactTextString(m) }
func (*Topology) ProtoMessage()               {}
func (*Topology) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Topology) GetNetworkInfo() *NetworkInfo {
	if m != nil {
		return m.NetworkInfo
	}
	return nil
}

func (m *Topology) GetInternals() []*Host {
	if m != nil {
		return m.Internals
	}
	return nil
}

func (m *Topology) GetLocalGws() []*Host {
	if m != nil {
		return m.LocalGws
	}
	return nil
}

type Host struct {
	Ip  uint32 `protobuf:"varint,1,opt,name=ip" json:"ip,omitempty"`
	Mac string `protobuf:"bytes,2,opt,name=mac" json:"mac,omitempty"`
}

func (m *Host) Reset()                    { *m = Host{} }
func (m *Host) String() string            { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()               {}
func (*Host) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Host) GetIp() uint32 {
	if m != nil {
		return m.Ip
	}
	return 0
}

func (m *Host) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type NetworkInfo struct {
	NetworkName    string `protobuf:"bytes,1,opt,name=networkName" json:"networkName,omitempty"`
	NetworkNetmask string `protobuf:"bytes,2,opt,name=networkNetmask" json:"networkNetmask,omitempty"`
	InterfaceIp    uint32 `protobuf:"varint,3,opt,name=interfaceIp" json:"interfaceIp,omitempty"`
}

func (m *NetworkInfo) Reset()                    { *m = NetworkInfo{} }
func (m *NetworkInfo) String() string            { return proto.CompactTextString(m) }
func (*NetworkInfo) ProtoMessage()               {}
func (*NetworkInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *NetworkInfo) GetNetworkName() string {
	if m != nil {
		return m.NetworkName
	}
	return ""
}

func (m *NetworkInfo) GetNetworkNetmask() string {
	if m != nil {
		return m.NetworkNetmask
	}
	return ""
}

func (m *NetworkInfo) GetInterfaceIp() uint32 {
	if m != nil {
		return m.InterfaceIp
	}
	return 0
}

func init() {
	proto.RegisterType((*Topology)(nil), "netsensor.Topology")
	proto.RegisterType((*Host)(nil), "netsensor.Host")
	proto.RegisterType((*NetworkInfo)(nil), "netsensor.NetworkInfo")
}

func init() { proto.RegisterFile("topology.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xe5, 0x04, 0xa1, 0xe6, 0xa2, 0x16, 0xe4, 0x01, 0x79, 0x0c, 0x19, 0x50, 0x24, 0x84,
	0x87, 0xb2, 0x30, 0x67, 0x81, 0x2e, 0x55, 0x65, 0x31, 0xb1, 0x99, 0xe0, 0x42, 0xd4, 0xd8, 0x67,
	0xc5, 0x46, 0x55, 0xff, 0x4d, 0x7f, 0x2a, 0x8a, 0x69, 0x13, 0x0b, 0xb1, 0x25, 0xef, 0xbe, 0x3b,
	0xbd, 0xe7, 0x07, 0x0b, 0x8f, 0x16, 0x3b, 0xfc, 0x3c, 0x70, 0xdb, 0xa3, 0x47, 0x9a, 0x19, 0xe5,
	0x9d, 0x32, 0x0e, 0xfb, 0xf2, 0x48, 0x60, 0xf6, 0x7a, 0x9a, 0xd2, 0x27, 0xc8, 0x8d, 0xf2, 0x7b,
	0xec, 0x77, 0x2b, 0xb3, 0x45, 0x46, 0x0a, 0x52, 0xe5, 0xcb, 0x1b, 0x3e, 0xd2, 0x7c, 0x3d, 0x4d,
	0x45, 0x8c, 0xd2, 0x07, 0xc8, 0x5a, 0xe3, 0x55, 0x6f, 0x64, 0xe7, 0x58, 0x52, 0xa4, 0x55, 0xbe,
	0xbc, 0x8a, 0xf6, 0x5e, 0xd0, 0x79, 0x31, 0x11, 0xf4, 0x1e, 0x66, 0x1d, 0x36, 0xb2, 0x7b, 0xde,
	0x3b, 0x96, 0xfe, 0x4f, 0x8f, 0x40, 0x59, 0xc1, 0xc5, 0xa0, 0xd0, 0x05, 0x24, 0xad, 0x0d, 0xa6,
	0xe6, 0x22, 0x69, 0x2d, 0xbd, 0x86, 0x54, 0xcb, 0x86, 0x25, 0x05, 0xa9, 0x32, 0x31, 0x7c, 0x96,
	0x07, 0xc8, 0x23, 0x87, 0xb4, 0x18, 0xe3, 0xac, 0xa5, 0x56, 0x61, 0x33, 0x13, 0xb1, 0x44, 0xef,
	0x60, 0x71, 0xfe, 0x55, 0x5e, 0x4b, 0xb7, 0x3b, 0x5d, 0xfb, 0xa3, 0x0e, 0x97, 0x82, 0xf9, 0xad,
	0x6c, 0xd4, 0xca, 0xb2, 0x34, 0x78, 0x88, 0xa5, 0xba, 0x86, 0xdb, 0x6f, 0xc7, 0xdb, 0x2f, 0xdd,
	0x70, 0x83, 0x1f, 0x4a, 0xa3, 0xe1, 0xbf, 0x69, 0xdc, 0x14, 0xac, 0x9e, 0x9f, 0x5f, 0x7a, 0x33,
	0xd4, 0xb0, 0x21, 0x6f, 0x53, 0x11, 0x47, 0x42, 0xde, 0x2f, 0x43, 0x3b, 0x8f, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x50, 0x6b, 0x1d, 0xe3, 0xaf, 0x01, 0x00, 0x00,
}
