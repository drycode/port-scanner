package utils

import (
	"fmt"
	"net"
	"testing"
)

func TestDeriveFromCIDR(t *testing.T) {
	type test struct {
		input    string
		expected [2]net.IP
	}
	cidrBlockTests := []test{
		{input: "3.58.1.97/12", expected: [2]net.IP{net.ParseIP("3.48.0.0"), net.ParseIP("3.63.255.255")}},
		{input: "123.0.0.1/31", expected: [2]net.IP{net.ParseIP("123.0.0.0"), net.ParseIP("123.0.0.1")}},
		{input: "10.0.0.1/17", expected: [2]net.IP{net.ParseIP("10.0.0.0"), net.ParseIP("10.0.127.255")}},
	}
	for _, block := range cidrBlockTests {
		testname := fmt.Sprintf("%s", block.input)
		t.Run(testname, func(t *testing.T) {
			range_ := DeriveFromCIDR(block.input)
			AssertEquals(t, "Check Range", block.expected[0].String(), range_[0].String())
			AssertEquals(t, "Check Range", block.expected[1].String(), range_[1].String())
		})

	}
}
