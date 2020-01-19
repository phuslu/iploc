// Package geoip provides fastest GeoIP Country library for Go.
//
// eg.
//
//      country := geoip.Country(net.ParseIP("1.1.1.1"))
//      fmt.Printf("%s\n", country)
//
//      // Output: US
package geoip

import (
	"encoding/binary"
	"net"
)

// Country find iso3166 country code of IP.
func Country(ip net.IP) (country []byte) {
	if ip == nil {
		return
	}
	if ip4 := ip.To4(); ip4 != nil {
		country = country4(binary.BigEndian.Uint32(ip4))
	} else {
		country = country6(binary.BigEndian.Uint64(ip), binary.BigEndian.Uint64(ip[8:]))
	}
	return
}

func country4(n uint32) (country []byte) {
	i, j := 0, len(ips)
	for i < j {
		h := int(uint(i+j) >> 1)
		if ips[h] <= n {
			i = h + 1
		} else {
			j = h
		}
	}

	country = geo[i*2-2 : i*2]

	return
}
