// +build 386 amd64 amd64p32 arm64 ppc64le mipsle mips64le mips64p32le

package geoip

import (
	"unsafe"
)

func country6(high, low uint64) (country []byte) {
	i, j := 0, len(ips6)/16
	for i < j {
		h := int(uint(i+j) >> 1)
		hi := *(*uint64)(unsafe.Pointer(&ips6[h*16]))
		lo := *(*uint64)(unsafe.Pointer(&ips6[h*16+8]))
		if hi > high || (hi == high && lo > low) {
			j = h
		} else {
			i = h + 1
		}
	}

	country = geo6[i*2-2 : i*2]

	return
}
