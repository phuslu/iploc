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
	"encoding/binary"
)

func ip6uint(i int) uint64 {
	return binary.LittleEndian.Uint64(ips6[i:])
}
