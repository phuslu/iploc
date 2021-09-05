// +build arm ppc64be mipsbe mips64be mips64p32be

package iploc

import (
	_ "embed" // for ip data
)

//go:embed ipv4be.bin
var ip4bin []byte

//go:embed ipv6be.gz
var ip6bin []byte
