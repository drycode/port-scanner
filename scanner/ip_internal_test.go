package scanner

import (
	"net"
	"testing"

	. "github.com/drypycode/portscanner/utils"
)

func TestParseIPRange(t *testing.T) {
	input := "123.45.1.34-123.45.1.38"
	ipRange := parseIPRange(input)
	AssertEquals(t, "Check parse IP range", []net.IP{
		net.IP{123, 45, 1, 34}.To4(),
		net.IP{123, 45, 1, 35}.To4(),
		net.IP{123, 45, 1, 36}.To4(),
		net.IP{123, 45, 1, 37}.To4(),
	}, ipRange)
}
