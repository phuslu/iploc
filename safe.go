// +build !386
// +build !amd64
// +build !amd64p32
// +build !arm64
// +build !ppc64le
// +build !mipsle
// +build !mips64le
// +build !mips64p32le

package iploc

import (
	_ "embed" // for ip data
	"encoding/binary"
)

//go:embed ipv4
var ip4data []byte

//go:embed ipv6.gz
var ip6data []byte

func ip4uint(i int) uint32 {
	return binary.LittleEndian.Uint32(ip4data[i*6:])
}

func ip6uint(i int) uint64 {
	return binary.LittleEndian.Uint64(ip6data[i*18:])
}
