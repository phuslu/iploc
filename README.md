# geoip - fastest geoip country library

[![godoc][godoc-img]][godoc] [![release][release-img]][release] [![goreport][goreport-img]][goreport] [![coverage][coverage-img]][coverage]

### Getting Started

try on https://play.golang.org/p/0quZRSXdTjq
```go
package main

import (
	"fmt"
	"net"
	"github.com/phuslu/geoip"
)

func main() {
	fmt.Printf("%s", geoip.Country(net.ParseIP("2001:4860:4860::8888")))
}

// Output: US
```

### Command Tool
```bash
$ go get github.com/phuslu/geoip/cmd/geoip
$ geoip 1.1.1.1
US
```

### Benchmarks
```
BenchmarkGeoIpCountryForIPv4-8   	20.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGeoIpCountryForIPv6-8   	38.2 ns/op	       0 B/op	       0 allocs/op
```

### Acknowledgment
This site or product includes IP2Location LITE data available from http://www.ip2location.com.

### How often are the database(geo_db.go) updated?
Following IP2Location Lite Database, usually **monthly**.

[godoc-img]: http://img.shields.io/badge/godoc-reference-blue.svg
[godoc]: https://godoc.org/github.com/phuslu/geoip
[release-img]: https://img.shields.io/github/v/tag/phuslu/geoip?label=release
[release]: https://github.com/phuslu/geoip/releases
[goreport-img]: https://goreportcard.com/badge/github.com/phuslu/geoip
[goreport]: https://goreportcard.com/report/github.com/phuslu/geoip
[coverage-img]: https://img.shields.io/badge/coverage-100%25-brightgreen
[coverage]: https://gocover.io/github.com/phuslu/geoip
