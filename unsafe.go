// +build 386 amd64 amd64p32 arm64 ppc64le mipsle mips64le mips64p32le

package geoip

import (
	"unsafe"
)

func ip6uint(i int) uint64 {
	return *(*uint64)(unsafe.Pointer(&ips6[i]))
}
