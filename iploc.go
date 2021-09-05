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
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io"
	"net"
)

// Version is iplocation database version.
const Version = "v1.0.20210830"

func init() {
	r, _ := gzip.NewReader(bytes.NewReader(ip6data))
	ip6data, _ = io.ReadAll(r)
}

// Country return ISO 3166-1 alpha-2 country code of IP.
func Country(ip net.IP) (country []byte) {
	if ip == nil {
		return
	}

	if ip4 := ip.To4(); ip4 != nil {
		// ipv4
		n := binary.BigEndian.Uint32(ip4)
		i, j := 0, len(ip4data)/6
		for i < j {
			h := (i + j) >> 1
			if ip4uint(h) > n {
				j = h
			} else {
				i = h + 1
			}
		}
		country = ip4data[i*6-2 : i*6]
	} else {

		// ipv6
		high := binary.BigEndian.Uint64(ip)
		low := binary.BigEndian.Uint64(ip[8:])
		i, j := 0, len(ip6data)/18
		for i < j {
			h := (i + j) >> 1
			n := ip6uint(h)
			if n > high || (n == high && ip6uint(h+1) > low) {
				j = h
			} else {
				i = h + 1
			}
		}
		country = ip6data[i*18-2 : i*18]
	}

	return
}
