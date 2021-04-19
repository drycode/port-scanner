package argparse

import (
	"fmt"
	"testing"

	. "github.com/drypycode/portscanner/utils"
)

var nilSlice []int

func TestArgParsingDefaults(t *testing.T) {
	cla := getArgs()
	AssertEquals(t, "Default SpecifiedPorts", nilSlice, cla.SpecifiedPorts)
	AssertEquals(t, "Default Port Range", [2]int{0, 0}, cla.PortRange)
	AssertEquals(t, "Default Host", []string{"127.0.0.1"}, cla.Hosts)
	AssertEquals(t, "Default Protocol", "TCP", cla.Protocol)
	AssertEquals(t, "Default Timeout", 5000, cla.Timeout)
}

func TestParsePorts(t *testing.T) {
	portRange := "134-345"
	ports := parsePorts(portRange)
	AssertEquals(t, "", 2, len(ports))
}

func TestParseSpecifiedPorts(t *testing.T) {
	type test struct {
		input    string
		expected []int
	}
	testValidInput := []test{
		{input: "80,443", expected: []int{80, 443}},
		{input: "", expected: nilSlice},
		{input: "1,3,5,6,7", expected: []int{1, 3, 5, 6, 7}},
	}
	for _, param := range testValidInput {
		testname := fmt.Sprintf("%s, %d", param.input, param.expected)
		t.Run(testname, func(t *testing.T) {
			value, _ := parseSpecifiedPorts(param.input)
			AssertEquals(t, "Specified Ports", param.expected, value)
		})
	}

	testInvalidInput := []string{"-1, -2, 0", "banana", "ice,cream"}
	for _, param := range testInvalidInput {
		testname := fmt.Sprintf("parseSpecifiedPorts invalid input: %s", param)
		t.Run(testname, func(t *testing.T) {
			_, err := parseSpecifiedPorts(param)
			if err == nil {
				t.Error(testname + " did not return an error")
			}
		})
	}
}

func TestGetAllPorts(t *testing.T) {
	ports := sortAllPorts([2]int{90, 100}, []int{80, 443})
	expected := []int{80, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 443}
	for i := 0; i < len(ports); i++ {
		AssertEquals(t, "Checking Ports Equality", expected[i], ports[i])
	}

}

func TestParseHosts(t *testing.T) {
	hosts := parseHosts("google.com,facebook.com,localhost,192.0.0.1")
	AssertEquals(t, "Hosts parsed correctly", 4, len(hosts))
}

func TestValidatePorts(t *testing.T) {
}
