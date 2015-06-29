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
Example:
```
λ ciphermaster-windows-0.2.exe -host=hostname.net.tld -pass=fancypass -only=WWW_1 -ciphers=ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384 -dryrun
2015/06/29 09:22:21 Skipping Service WWW_01 because it does not match -only
2015/06/29 09:22:21 Skipping Service WWW_02 because it does not match -only
2015/06/29 09:22:21 Skipping Service WWW_03 because it does not match -only
2015/06/29 09:22:21 Skipping Service WWW_04 because it does not match -only
2015/06/29 09:22:21 Skipping Service WWW_05 because it does not match -only
2015/06/29 09:22:21 Skipping Service WWW_06 because it does not match -only
2015/06/29 09:22:21 Skipping Service WWW_07 because it does not match -only
2015/06/29 09:22:21 Skipping Service WWW_08 because it does not match -only
2015/06/29 09:22:21 Skipping Service WWW_09 because it does not match -only
2015/06/29 09:22:21 Dry run. Would be setting ciphers of WWW_10 to ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384
2015/06/29 09:22:21 Dry run. Would be setting ciphers of WWW_11 to ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384
```