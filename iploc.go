// Package iploc provides fastest Geolocation Country library for Go.
//
//      package main
//
//      import (
//      	"fmt"
//      	"net"
//      	"github.com/phuslu/iploc"
//      )
//
//      func main() {
//      	fmt.Printf("%s", iploc.Country(net.ParseIP("2001:4860:4860::8888")))
//      }
//
//      // Output: US
package iploc

import (
	"encoding/binary"
	"net"
)

const Version = "v1.0.20201130"

// Country return ISO 3166-1 alpha-2 country code of IP.
func Country(ip net.IP) (country []byte) {
	if ip == nil {
		return
	}

	if ip4 := ip.To4(); ip4 != nil {
		// ipv4
		n := binary.BigEndian.Uint32(ip4)
		i, j := 0, len(ips)
		for i < j {
			h := (i + j) >> 1
			if ips[h] > n {
				j = h
			} else {
				i = h + 1
			}
		}
		country = geo[i<<1-2 : i<<1]
	} else {
		// ipv6
		high := binary.BigEndian.Uint64(ip)
		low := binary.BigEndian.Uint64(ip[8:])
		i, j := 0, len(ips6)
		for i < j {
			h := (i + j) >> 1 & ^0xf
			n := ip6uint(h)
			if n > high || (n == high && ip6uint(h+8) > low) {
				j = h
			} else {
				i = h + 16
			}
		}
		country = geo6[i>>3-2 : i>>3]
	}

	return
}

// IsReservedIPv4 detects a net.IP is a reserved address, return fasle if IPv6
func IsReservedIPv4(ip net.IP) bool {
	if ip4 := ip.To4(); ip4 != nil {
		_ = ip4[3]
		switch ip4[0] {
		case 10:
			return true
		case 100:
			return ip4[1] >= 64 && ip4[1] <= 127
		case 127:
			return true
		case 169:
			return ip4[1] == 254
		case 172:
			return ip4[1] >= 16 && ip4[1] <= 31
		case 192:
			switch ip4[1] {
			case 0:
				switch ip4[2] {
				case 0, 2:
					return true
				}
			case 18, 19:
				return true
			case 51:
				return ip4[2] == 100
			case 88:
				return ip4[2] == 99
			case 168:
				return true
			}
		case 203:
			return ip4[1] == 0 && ip4[2] == 113
		case 224:
			return true
		case 240:
			return true
		}
	}
	return false
}
