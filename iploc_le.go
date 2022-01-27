// +build 386 amd64 arm amd64p32 arm64 ppc64le mipsle mips64le mips64p32le

package iploc

import (
	_ "embed" // for ip data
)

//go:embed ipv4le.bin
var ip4bin []byte

//go:embed ipv6le.gz
var ip6bin []byte
