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
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkIPCountryForIPv4
BenchmarkIPCountryForIPv4-8     74145014                14.96 ns/op            0 B/op          0 allocs/op
BenchmarkIPCountryForIPv6
BenchmarkIPCountryForIPv6-8     59173639                22.73 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/phuslu/iploc 2.637s
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
