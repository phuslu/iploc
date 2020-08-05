// Package geoip provides fastest GeoIP Country library for Go.
//
//      package main
//
//      import (
//      	"fmt"
//      	"net"
//      	"github.com/phuslu/geoip"
//      )
//
//      func main() {
//      	fmt.Printf("%s", geoip.Country(net.ParseIP("2001:4860:4860::8888")))
//      }
//
//      // Output: US
package geoip

import (
	"encoding/binary"
	"net"
)

const Version = "v1.0.20200730"

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
