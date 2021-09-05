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
    ipinfo = []
    for line in io.BytesIO(text):
        parts = line.strip().decode().split(',')
        ip = parts[0].strip('"')
        country = parts[2].strip('"')
        if country == '-':
            country = 'ZZ'
        ipinfo.append((ip, country.encode()))
    return ipinfo


def gen_ip4_data() -> bytes:
    """generate ipv4 to binary data"""
    ipinfo = get('http://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.CSV.ZIP')
    pack = lambda ip: struct.pack('<I', ip)
    data = b''.join(pack(int(ip))+country for (ip, country) in ipinfo)
    with open('ipv4', 'wb') as file:
        file.write(data)


def gen_ip6_data() -> bytes:
    """generate ipv6 to binary data"""
    ipinfo = get('http://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.IPV6.CSV.ZIP')
    pack = lambda ip: struct.pack('<Q', ip >> 64) + struct.pack('<Q', ip & 0xFFFFFFFFFFFFFFFF)
    data = b''.join(pack(int(ip))+country for (ip, country) in ipinfo)
    with gzip.GzipFile('ipv6.gz', 'wb') as file:
        file.write(data)


def main():
    """convert ip2location country csv to binary data"""
    gen_ip4_data()
    gen_ip6_data()


if __name__ == '__main__':
    main()
