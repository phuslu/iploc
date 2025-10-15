// Package iploc provides fastest IP to Country library for Go.
//
//	package main
//
//	import (
//		"fmt"
//		"net/netip"
//		"github.com/phuslu/iploc"
//	)
//
//	func main() {
//		fmt.Printf("%s", iploc.IPCountry(netip.MustParseAddr("1.1.1.1"))
//	}
//
//	// Output: US
package iploc

import (
	"bytes"
	"compress/gzip"
	_ "embed" // for ip data
	"io"
	"net"
	"net/netip"
	"os"
	"reflect"
	"sync"
	"unsafe"
)

// Version is iplocation database version.
const Version = "v1.0.20251015"

//go:embed ipv4.txt
var ip4txt []byte

//go:embed ipv6.txt
var ip6txt []byte

var ipv6once sync.Once
var ipv4only bool

// IPCountry return ISO 3166-1 alpha-2 country code of IP.
func IPCountry(ip netip.Addr) (country string) {
	if ip.Is4() {
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
		// country = ip4txt[i*2-2 : i*2]
		sh := (*reflect.StringHeader)(unsafe.Pointer(&country))
		sh.Data = uintptr(unsafe.Add(unsafe.Pointer(&ip4txt[0]), uintptr(i*2-2)))
		sh.Len = 2
		return
	}

	ipv6once.Do(func() {
		ipv4only = os.Getenv("IPLOC_IPV4ONLY") != ""
		if !ipv4only {
			r, _ := gzip.NewReader(bytes.NewReader(ip6bin))
			ip6bin, _ = io.ReadAll(r)
		}
	})

	if ipv4only {
		return
	}

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
	// country = ip6txt[i-2 : i]
	sh := (*reflect.StringHeader)(unsafe.Pointer(&country))
	sh.Data = uintptr(unsafe.Add(unsafe.Pointer(&ip6txt[0]), uintptr(i-2)))
	sh.Len = 2
	return
}

// Country return ISO 3166-1 alpha-2 country code of IP.
func Country(ip net.IP) (country string) {
	if addr, ok := netip.AddrFromSlice(ip); ok {
		return IPCountry(addr)
	}
	return
}
