package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
)

var assertionStatement = "%s -- expected: %v | actual: %v \n"

// Equals... fails the test if exp is not equal to act.
func AssertEquals(tb testing.TB, name string, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		tb.Errorf(assertionStatement, name, exp, act)
	}
}

type SafeSlice struct {
	mu        sync.RWMutex
	OpenPorts []string
}

func (ss *SafeSlice) Length() int {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	return len(ss.OpenPorts)
}

func (ss *SafeSlice) Append(val string, wg *sync.WaitGroup) {
	ss.mu.Lock()
	ss.OpenPorts = append((*ss).OpenPorts, val)
	ss.mu.Unlock()
	wg.Done()
}

// GetUlimit ...
func GetUlimit() (int, error) {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		return -1, err
	}
	ulimit := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(ulimit, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(i), nil
}

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

func IPRangeFromFirstLast(first net.IP, last net.IP) []string {
	var ip4s []string
	for 0 != bytes.Compare(first, last) {
		ip4s = append(ip4s, first.String())
		first = nextIP(first)
	}
	return ip4s
}

// ParseIpRange ...
func ParseIPRange(ipRange string) []string {
	// Assume input of IPs is valid and separated by -
	stringRange := strings.Split(ipRange, "-")
	startIP, endIP := net.ParseIP(stringRange[0]).To4(), net.ParseIP(stringRange[1]).To4()
	return IPRangeFromFirstLast(startIP, endIP)
}

func intToIP(ip uint32) net.IP {
	result := make(net.IP, 4)
	result[3] = byte(ip)
	result[2] = byte(ip >> 8)
	result[1] = byte(ip >> 16)
	result[0] = byte(ip >> 24)
	return result
}

func DeriveFromCIDR(cidr string) [2]net.IP {
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		logrus.Debug(err)
	}
	first := subnet.IP
	firstInt := binary.BigEndian.Uint32(first)
	maskInt := binary.BigEndian.Uint32(subnet.Mask)
	lastInt := (firstInt & maskInt) | (maskInt ^ 0xffffffff)

	last := intToIP(lastInt)

	fmt.Println(first, last)
	return [2]net.IP{first, last}
}
