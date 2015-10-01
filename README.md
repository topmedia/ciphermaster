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
  -only="": Only if the service matches this string ciphers will be changed.
  -user="bal": Administrator username
	-dryrun=false: Do not go through with changes, logging only
```

## Binary Download

Pre-compiled binaries for Windows, Linux, and OS X are available at https://github.com/topmedia/ciphermaster/releases

## Example

```
$ ciphermaster -host=vlm.host.tld -pass=2fourall -only=http -ciphers=ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384 -dryrun
2015/06/29 09:22:21 Skipping Service db01 because it does not match -only
2015/06/29 09:22:21 Skipping Service db02 because it does not match -only
2015/06/29 09:22:21 Dry run. Would be setting ciphers of http01 to ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384
2015/06/29 09:22:21 Dry run. Would be setting ciphers of http02 to ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384
```
