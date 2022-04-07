/*
* Run: GODEBUG=netdns=go go run gordns.go
 */
package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <x.x.x.x/x>\n", os.Args[0])
		return
	}

	cidr := os.Args[1]
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err)
	}

	// network address is in CIDR notation, so get just the address part
	networkstart := strings.Split(network.String(), "/")[0]

	// figure out which IPs should be included
	strip := networkstart
	ips := []string{}
	for {
		if network.Contains(net.ParseIP(strip)) {
			ips = append(ips, strip)
		} else {
			break
		}
		strip = NextIP(strip)
	}

	// lookup each PTR in turn
	for _, v := range ips {
		ptr, err := net.LookupAddr(v)
		if err != nil {
			continue
		}

		for _, dnsname := range ptr {
			fmt.Printf("%s - %s\n", v, dnsname)
		}

	}

}

// Oh, I feel dirty doing this...
func NextIP(ip string) string {
	stroctets := strings.Split(ip, ".")

	// convert from string IP to decimal
	ioctets := [4]int{0, 0, 0, 0}
	ioctets[0], _ = strconv.Atoi(stroctets[0])
	ioctets[1], _ = strconv.Atoi(stroctets[1])
	ioctets[2], _ = strconv.Atoi(stroctets[2])
	ioctets[3], _ = strconv.Atoi(stroctets[3])

	// increment address by 1
	if ioctets[3] == 255 {
		ioctets[3] = 0
		if ioctets[2] == 255 {
			ioctets[2] = 0
			if ioctets[1] == 255 {
				ioctets[1] = 0
				if ioctets[0] == 255 {
					ioctets[0] = 0
				} else {
					ioctets[0]++
				}
			} else {
				ioctets[1]++
			}
		} else {
			ioctets[2]++
		}
	} else {
		ioctets[3]++
	}

	// back to string
	return fmt.Sprintf("%d.%d.%d.%d", ioctets[0], ioctets[1], ioctets[2], ioctets[3])
}
