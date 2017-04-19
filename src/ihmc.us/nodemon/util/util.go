package util

import (
	"net"
	"encoding/binary"
	"runtime"
	"bytes"
	"strconv"
	"time"
)

// Gets the ID of the goroutine through an hack of the Go runtime,
// use only to print for debug purposes
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// Returns current GID as string using GetGID()
func GetGIDString() string {
	return strconv.Itoa(int(GetGID()))
}

// Converts a net.IP type to uint32
func IPToUInt32(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// Converts a uint32 type to net.IP
func UInt32ToIP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}
