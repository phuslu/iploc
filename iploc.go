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
	"unsafe"
)

// Version is iplocation database version.
const Version = "v1.0.20240311"

//go:embed ipv4.txt
var ip4txt []byte

//go:embed ipv6.txt
var ip6txt []byte

func init() {
	// ipv6
	if os.Getenv("IPLOC_IPV4ONLY") == "" {
		r, _ := gzip.NewReader(bytes.NewReader(ip6bin))
		ip6bin, _ = io.ReadAll(r)
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
		i, j := 0, len(ip4bin)/4
		for i < j {
			h := (i + j) >> 1
			if *(*uint32)(unsafe.Add(unsafe.Pointer(&ip4bin[0]), uintptr(h*4))) > n {
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
		i, j := 0, len(ip6bin)/8
		for i < j {
			h := (i + j) >> 1 & ^1
			n := *(*uint64)(unsafe.Add(unsafe.Pointer(&ip6bin[0]), uintptr(h*8)))
			if n > high || (n == high && *(*uint64)(unsafe.Add(unsafe.Pointer(&ip6bin[0]), uintptr((h+1)*8))) > low) {
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
		i, j := 0, len(ip4bin)/4
		for i < j {
			h := (i + j) >> 1
			if *(*uint32)(unsafe.Add(unsafe.Pointer(&ip4bin[0]), uintptr(h*4))) > n {
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
		i, j := 0, len(ip6bin)/8
		for i < j {
			h := (i + j) >> 1 & ^1
			n := *(*uint64)(unsafe.Add(unsafe.Pointer(&ip6bin[0]), uintptr(h*8)))
			if n > high || (n == high && *(*uint64)(unsafe.Add(unsafe.Pointer(&ip6bin[0]), uintptr((h+1)*8))) > low) {
				j = h
			} else {
				i = h + 2
			}
		}
		country = ip6txt[i-2 : i]
	}

	return
}
