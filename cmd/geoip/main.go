package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/phuslu/geoip"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <host>\n", filepath.Base(os.Args[0]))
		return
	}

	s := os.Args[1]

	switch s {
	case "-v", "-version", "--version":
		fmt.Printf("%s\n", geoip.Version)
		return
	}

	if ip := net.ParseIP(s); ip != nil {
		fmt.Printf("%s\n", geoip.Country(ip))
		return
	}

	ips, err := net.LookupIP(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", geoip.Country(ips[0]))
}
