# geoip - fastest geoip country library

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/phuslu/geoip) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/phuslu/geoip/master/LICENSE) [![goreport](https://goreportcard.com/badge/github.com/phuslu/geoip)](https://goreportcard.com/report/github.com/phuslu/geoip) [![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://gocover.io/github.com/phuslu/geoip)

### Getting Started

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

### Benchmarks
```
BenchmarkGeoIpCountryForIPv4-8   	20.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGeoIpCountryForIPv6-8   	38.2 ns/op	       0 B/op	       0 allocs/op
```

### Acknowledgment
This site or product includes IP2Location LITE data available from http://www.ip2location.com.
