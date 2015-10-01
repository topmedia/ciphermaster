# ciphermaster

Set KEMP LoadMaster ciphers for all SSL services

## Usage:

```
Usage of ./ciphermaster:
  -cipers="ECDHE-RSA-AES256-GCM-SHA384": Ciphers to assign to services, colon-separated
  -host="192.168.110.123": Load Balancer Hostname/IP
  -listciphers=false: List available ciphers and exit
  -pass="2fourall": Administrator password
  -skip="": If the service matches this string, no ciphers will be changed.
  -user="bal": Administrator username
```

## Binary Download

Pre-compiled binaries for Windows, Linux, and OS X are available at https://github.com/topmedia/ciphermaster/releases
