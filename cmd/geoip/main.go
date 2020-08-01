package main

import (
	"fmt"
	"net"
	"os"

	"github.com/phuslu/geoip"
)

func main() {
	fmt.Printf("%s\n", geoip.Country(net.ParseIP(os.Args[1])))
}
