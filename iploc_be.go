// +build ppc64 mips mips64 mips64p32

package iploc

import (
	_ "embed" // for ip data
)

//go:embed ipv4be.bin
var ip4bin []byte

//go:embed ipv6be.gz
var ip6bin []byte
