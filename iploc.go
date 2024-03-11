// Package iploc provides fastest Geolocation Country library for Go.
//
//	package main
//
//	import (
//		"fmt"
//		"net"
//		"github.com/phuslu/iploc"
//	)
//
//	func main() {
//		fmt.Printf("%s", iploc.Country(net.ParseIP("2001:4860:4860::8888")))
//	}
//
//	// Output: US
package iploc

import (
	"bytes"
	"compress/gzip"
	_ "embed" // for ip data
	"encoding/binary"
	"io"
	"net"
	"net/netip"
	"os"
	"reflect"
	"unsafe"
)

// Version is iplocation database version.
const Version = "v1.0.20240229"

var ip4uint []uint32
var ip6uint []uint64

//go:embed ipv4.txt
var ip4txt []byte

//go:embed ipv6.txt
var ip6txt []byte

func init() {
	// ipv4
	ip4uint = *(*[]uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&ip4bin[0])),
		Len:  len(ip4bin) / 4,
		Cap:  len(ip4bin) / 4,
	}))

	// ipv6
	if os.Getenv("IPLOC_IPV4ONLY") == "" {
		r, _ := gzip.NewReader(bytes.NewReader(ip6bin))
		ip6bin, _ = io.ReadAll(r)
		ip6uint = *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(&ip6bin[0])),
			Len:  len(ip6bin) / 8,
			Cap:  len(ip6bin) / 8,
		}))
	} else {
		ip6uint = []uint64{0, 0}
	}
}

// Country return ISO 3166-1 alpha-2 country code of IP.
func Country(ip net.IP) (country []byte) {
	if ip == nil {
		return
	}

	if ip4 := ip.To4(); ip4 != nil {
		// ipv4
		n := binary.BigEndian.Uint32(ip4)
		i, j := 0, len(ip4uint)
		_ = ip4uint[j-1]
		for i < j {
			h := (i + j) >> 1
			if ip4uint[h] > n {
				j = h
			} else {
				i = h + 1
			}
		}
		country = ip4txt[i*2-2 : i*2]
	} else {
		// ipv6
		high := binary.BigEndian.Uint64(ip)
		low := binary.BigEndian.Uint64(ip[8:])
		i, j := 0, len(ip6uint)
		_ = ip6uint[j-1]
		for i < j {
			h := (i + j) >> 1 & ^1
			n := ip6uint[h]
			if n > high || (n == high && ip6uint[h+1] > low) {
				j = h
			} else {
				i = h + 2
			}
		}
		country = ip6txt[i-2 : i]
	}

	return
}

// IPCountry return ISO 3166-1 alpha-2 country code of IP.
func IPCountry(ip netip.Addr) (country []byte) {
	if ip.Is4() {
		// ipv4
		b := ip.As4()
		n := uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
		i, j := 0, len(ip4uint)
		_ = ip4uint[j-1]
		for i < j {
			h := (i + j) >> 1
			if ip4uint[h] > n {
				j = h
			} else {
				i = h + 1
			}
		}
		country = ip4txt[i*2-2 : i*2]
	} else {
		// ipv6
		b := ip.As16()
		high := uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
		low := uint64(b[8])<<56 | uint64(b[9])<<48 | uint64(b[10])<<40 | uint64(b[11])<<32 | uint64(b[12])<<24 | uint64(b[13])<<16 | uint64(b[14])<<8 | uint64(b[15])
		i, j := 0, len(ip6uint)
		_ = ip6uint[j-1]
		for i < j {
			h := (i + j) >> 1 & ^1
			n := ip6uint[h]
			if n > high || (n == high && ip6uint[h+1] > low) {
				j = h
			} else {
				i = h + 2
			}
		}
		country = ip6txt[i-2 : i]
	}

	return
}
