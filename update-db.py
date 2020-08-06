#!/usr/bin/env python3
# pylint: disable=too-many-statements, line-too-long, W0703

import io
import urllib.request
import struct
import base64
import zipfile
import zlib


def get(url: str) -> (list, list):
    """extract country list and ip list from ip2loc url"""
    zfile = zipfile.ZipFile(io.BytesIO(urllib.request.urlopen(url).read()))
    text = zfile.read(next(x.filename for x in zfile.filelist if x.filename.endswith('.CSV')))
    geo, ips = '', []
    for line in io.BytesIO(text):
        parts = line.strip().decode().split(',')
        ip = parts[0].strip('"')
        country = parts[2].strip('"')
        if country == '-':
            country = 'ZZ'
        geo += country
        ips.append(ip)
    return geo, ips


def gen6() -> str:
    """generate ipv4 to geoip_db.go"""
    geo, ips = get('http://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.IPV6.CSV.ZIP')
    pack = lambda ip: struct.pack('<Q', ip >> 64) + struct.pack('<Q', ip & 0xFFFFFFFFFFFFFFFF)
    ips = base64.b64encode(zlib.compress(b''.join(pack(int(x)) for x in ips))).decode()
    return '''package geoip

import (
    "compress/zlib"
    "encoding/base64"
    "io/ioutil"
    "strings"
)

var geo6 = []byte("%s")
var ips6s = "%s"
var ips6 = func() []byte {
    r, _ := zlib.NewReader(base64.NewDecoder(base64.StdEncoding, strings.NewReader(ips6s)))
    b, _ := ioutil.ReadAll(r)
    return b
}()
''' % (geo, ips)


def gen4() -> str:
    """generate ipv4 to geoip_db.go"""
    geo, ips = get('http://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.CSV.ZIP')
    return '''
var geo = []byte("%s")
var ips = []uint32{%s}
''' % (geo, ','.join(ips))


def main():
    """convert ip2location country csv to geoip_db.go"""
    text = gen6() + gen4()
    with open('geoip_db.go', 'wb') as file:
        file.write(text.encode())


if __name__ == '__main__':
    main()

