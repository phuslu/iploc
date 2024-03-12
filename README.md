# iploc - fastest ip country library

[![godoc][godoc-img]][godoc] [![release][release-img]][release] [![goreport][goreport-img]][goreport]

### Getting Started

try on https://play.golang.org/p/T_7jfSr0cE1
```go
package main

import (
	"fmt"
	"net/netip"
	"github.com/phuslu/iploc"
)

func main() {
	fmt.Printf("%s", iploc.IPCountry(netip.MustParseAddr("1.1.1.1"))
}

// Output: US
```

### Benchmarks
```
goos: windows
goarch: amd64
pkg: github.com/phuslu/iploc
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkIPCountryForIPv4
BenchmarkIPCountryForIPv4-8     80750439                13.57 ns/op            0 B/op          0 allocs/op
BenchmarkIPCountryForIPv6
BenchmarkIPCountryForIPv6-8     57166812                20.44 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/phuslu/iploc 2.360s
```

### Acknowledgment
This site or product includes IP2Location LITE data available from http://www.ip2location.com.

### How often are the inlined ip data updated?
Following IP2Location Lite Database, usually **monthly**.

[godoc-img]: http://img.shields.io/badge/godoc-reference-blue.svg
[godoc]: https://godoc.org/github.com/phuslu/iploc
[release-img]: https://img.shields.io/github/v/tag/phuslu/iploc?label=release
[release]: https://github.com/phuslu/iploc/releases
[goreport-img]: https://goreportcard.com/badge/github.com/phuslu/iploc
[goreport]: https://goreportcard.com/report/github.com/phuslu/iploc
[coverage-img]: http://gocover.io/_badge/github.com/phuslu/iploc
[coverage]: https://gocover.io/github.com/phuslu/iploc
