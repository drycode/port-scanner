package tests

import (
	"testing"

	. "github.com/drypycode/portscanner/argparse"
	. "github.com/drypycode/portscanner/utils"
)

func TestGetAllPorts(t *testing.T) {
	ucla := UnmarshalledCommandLineArgs{
		PortRange:      [2]int{90, 100},
		SpecifiedPorts: []int{80, 443},
		Host:           "localhost",
		Protocol:       "TCP",
		Timeout:        100,
	}
	ports := ucla.GetAllPorts()
	expected := []int{80, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 443}
	for i := 0; i < len(ports); i++ {
		AssertEquals(t, "Checking Ports Equality", expected[i], ports[i])
	}

}
