package scanner

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	"github.com/sirupsen/logrus"
)

func nextIP(ip net.IP) net.IP {
	ip4 := ip.To4()
	var updateRange func(i int)
	updateRange = func(i int) {
		if i < 0 {
			logrus.Fatal()
		}
		if ip4[i] < 255 {
			ip4[i] += 1
		} else {
			ip4[i] = 1
			updateRange(i - 1)
		}
	}
	updateRange(3)
	return ip4
}

func dupIP(ip net.IP) net.IP {
	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}

func parseIPRange(ipRange string) []net.IP {
	// Assume input of IPs is valid and separated by -
	stringRange := strings.Split(ipRange, "-")
	startIP, endIP := net.ParseIP(stringRange[0]).To4(), net.ParseIP(stringRange[1]).To4()
	var ip4s []net.IP
	for 0 != bytes.Compare(startIP, endIP) {
		ip4s = append(ip4s, dupIP(startIP))
		startIP = nextIP(startIP)
	}
	fmt.Println(endIP)
	return ip4s
}
