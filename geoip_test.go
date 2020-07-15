package geoip

import (
	"net"
	"testing"
)

func TestGeoIpCountry(t *testing.T) {
	cases := []struct {
		IP      string
		Country string
	}{
		{"", ""},
		{"0.0.0.0", "ZZ"},
		{"0.1.1.1", "ZZ"},
		{"1.1.1.1", "US"},
		{"121.229.143.64", "CN"},
		{"122.96.43.186", "CN"},
		// {"123.249.20.198", "CN"},
		{"153.3.123.160", "CN"},
		{"153.3.131.201", "CN"},
		{"180.109.81.198", "CN"},
		{"180.111.103.88", "CN"},
		{"183.206.11.225", "CN"},
		{"192.210.171.249", "US"},
		{"223.112.9.2", "CN"},
		{"23.16.28.232", "CA"},
		{"58.240.115.210", "CN"},
		{"61.155.4.66", "CN"},
		{"255.255.255.255", "ZZ"},
		{"2001:41d0:701:1100::29c8", "FR"},
		{"2001:4860:4860::8888", "US"},
		{"2001::6ca0:a535", "ZZ"},
		{"2001:dc7:1000::1", "CN"},
		{"2400:3200::1", "CN"},
		{"2400:da00::6666", "CN"},
		{"2404:6800:4008:801::2004", "TW"},
		{"2404:6800:4012:1::200e", "AU"},
		{"240C::6666", "CN"},
		{"240e:4c:4008::1", "CN"},
		{"240e:e8:f089:4877:70d2:775c:91d1:ab12", "CN"},
		{"2620:0:2d0:200::7", "US"},
		{"2a04:4e42:600::223", "NL"},
		{"::", "ZZ"},
		{"::1", "ZZ"},
	}

	for _, c := range cases {
		country := Country(net.ParseIP(c.IP))
		if string(country) != c.Country {
			t.Errorf("Country(%#v) return \"%s\", expect %#v", c.IP, country, c.Country)
		}
	}
}

func BenchmarkGeoIpCountryForIPv4(b *testing.B) {
	ip := net.IP{8, 8, 8, 8}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Country(ip)
	}
}

func BenchmarkGeoIpCountryForIPv6(b *testing.B) {
	ip := net.ParseIP("2001:4860:4860::8888")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Country(ip)
	}
}
