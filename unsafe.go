// +build 386 amd64 amd64p32 arm64 ppc64le mipsle mips64le mips64p32le

package iploc

import (
	_ "embed"
	"unsafe"
)

//go:embed ipv4
var ip4data []byte

//go:embed ipv6.gz
var ip6data []byte

func ip4uint(i int) uint32 {
	return *(*uint32)(unsafe.Pointer(&ip4data[i*6]))
}

func ip6uint(i int) uint64 {
	return *(*uint64)(unsafe.Pointer(&ip6data[i*18]))
}
