# iploc - fastest ip country library

[![godoc][godoc-img]][godoc] [![release][release-img]][release] [![goreport][goreport-img]][goreport] [![coverage][coverage-img]][coverage]

### Getting Started

try on https://play.golang.org/p/T_7jfSr0cE1
```go
package main

import (
	"fmt"
	"net"
	"github.com/phuslu/iploc"
)

func main() {
	fmt.Printf("%s", iploc.Country(net.IP{1,1,1,1}))
}

// Output: US
```

### Command Tool
```bash
$ go get github.com/phuslu/iploc/cmd/iploc
$ iploc 2001:4860:4860::8888
US
```

### Benchmarks
```
BenchmarkCountryForIPv4-2   	34826096	        32.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountryForIPv6-2   	29270091	        40.93 ns/op	       0 B/op	       0 allocs/op
```

### Acknowledgment
This site or product includes IP2Location LITE data available from http://www.ip2location.com.

### How often are the database(iploc_db.go) updated?
Following IP2Location Lite Database, usually **monthly**.

[godoc-img]: http://img.shields.io/badge/godoc-reference-blue.svg
[godoc]: https://godoc.org/github.com/phuslu/iploc
[release-img]: https://img.shields.io/github/v/tag/phuslu/iploc?label=release
[release]: https://github.com/phuslu/iploc/releases
[goreport-img]: https://goreportcard.com/badge/github.com/phuslu/iploc
[goreport]: https://goreportcard.com/report/github.com/phuslu/iploc
[coverage-img]: http://gocover.io/_badge/github.com/phuslu/iploc
[coverage]: https://gocover.io/github.com/phuslu/iploc
