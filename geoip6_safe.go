// +build !386
// +build !amd64
// +build !amd64p32
// +build !arm64
// +build !ppc64le
// +build !mipsle
// +build !mips64le
// +build !mips64p32le

package geoip

import (
	"encoding/binary"
)

func country6(high, low uint64) (country []byte) {
	i, j := 0, len(ips6)/16
	for i < j {
		h := int(uint(i+j) >> 1)
		hi := binary.LittleEndian.Uint64(ips6[h*16:])
		lo := binary.LittleEndian.Uint64(ips6[h*16+8:])
		if hi > high || (hi == high && lo > low) {
			j = h
		} else {
			i = h + 1
		}
	}

	country = geo6[i*2-2 : i*2]

	return
}
