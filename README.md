# gordns
Bulk PTR resolution by netblock. Provide network in CIDR format, tool will lookup the DNS name of every IP in that block. It will only output if a name is available. IPv4 ONLY!!

### Example
```
$ GODEBUG=netdns=go go run gordns.go 216.239.32.10/28
216.239.32.1 - any-in-2001.1e100.net.
216.239.32.2 - any-in-2002.1e100.net.
216.239.32.4 - any-in-2004.1e100.net.
216.239.32.5 - any-in-2005.1e100.net.
216.239.32.6 - any-in-2006.1e100.net.
216.239.32.8 - any-in-2008.1e100.net.
216.239.32.10 - ns1.google.com.
216.239.32.11 - ns.google.com.
216.239.32.15 - time1.google.com.
```
