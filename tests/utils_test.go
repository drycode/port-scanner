package tests

import (
	"strconv"
	"sync"
	"testing"

	. "github.com/drypycode/portscanner/utils"
)

func TestSafeSlice(t *testing.T) {
	ss := SafeSlice{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go ss.Append(strconv.Itoa(i), &wg)
	}
	wg.Wait()
	AssertEquals(t, "Check all ports evaluated", ss.Length(), 1000)
}

func TestParseIPRange(t *testing.T) {
	type test struct {
		input    string
		expected []string
	}
	testValidInput := []test{
		{input: "123.45.1.34-123.45.1.38", expected: []string{
			"123.45.1.34",
			"123.45.1.35",
			"123.45.1.36",
			"123.45.1.37",
		}},
		{input: "123.45.1.255-123.45.2.2", expected: []string{
			"123.45.1.255",
			"123.45.2.1",
		}},
		{input: "222.255.255.254-223.1.1.2", expected: []string{
			"222.255.255.254",
			"222.255.255.255",
			"223.1.1.1",
		}},
	}
	for _, test := range testValidInput {
		ipRange := ParseIPRange(test.input)
		AssertEquals(t, "Check parse IP range", test.expected, ipRange)
	}

}
