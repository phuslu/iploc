#!/usr/bin/env python3
# pylint: disable=too-many-statements, line-too-long, W0703

import gzip
import io
import struct
import urllib.request
import zipfile


def get(url: str) -> (list, list):
    """extract country list and ip list from ip2loc url"""
    zfile = zipfile.ZipFile(io.BytesIO(urllib.request.urlopen(url).read()))
    text = zfile.read(next(x.filename for x in zfile.filelist if x.filename.endswith('.CSV')))
    iplist, geo = [], b''
    for line in io.BytesIO(text):
        parts = line.strip().decode().split(',')
        ip = parts[0].strip('"')
        country = parts[2].strip('"')
        if country == '-':
            country = 'ZZ'
        iplist.append(ip)
        geo += country.encode()
    return iplist, geo


def gen_ip4_data() -> bytes:
    """generate ipv4 to binary data"""
    iplist, geo = get('http://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.CSV.ZIP')
    pack = lambda e, ip: struct.pack(e+'I', ip)
    with open('ipv4.txt', 'wb') as file:
        file.write(geo)
    with open('ipv4le.bin', 'wb') as file:
        file.write(b''.join(pack('<', int(ip)) for ip in iplist))
    with open('ipv4be.bin', 'wb') as file:
        file.write(b''.join(pack('>', int(ip)) for ip in iplist))


def gen_ip6_data() -> bytes:
    """generate ipv6 to binary data"""
    iplist, geo = get('http://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.IPV6.CSV.ZIP')
    pack = lambda e, ip: struct.pack(e+'Q', ip >> 64) + struct.pack(e+'Q', ip & 0xFFFFFFFFFFFFFFFF)
    with open('ipv6.txt', 'wb') as file:
        file.write(geo)
    with gzip.GzipFile('ipv6le.gz', 'wb') as file:
        file.write(b''.join(pack('<', int(ip)) for ip in iplist))
    with gzip.GzipFile('ipv6be.gz', 'wb') as file:
        file.write(b''.join(pack('>', int(ip)) for ip in iplist))


def main():
    """convert ip2location country csv to binary data"""
    gen_ip4_data()
    gen_ip6_data()


if __name__ == '__main__':
    main()
