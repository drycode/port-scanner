package argparse

import (
	"testing"

	. "github.com/drypycode/portscanner/utils"
)

func TestArgParsingDefaults(t *testing.T) {
	cla := getArgs()
	if cla.SpecifiedPorts == nil {
		t.Errorf("Expected non-null value")
	}
	AssertEquals(t, "Default Port Range", cla.PortRange, [2]int{0, 3000})
	AssertEquals(t, "Default Host", cla.Host, "127.0.0.1")
	AssertEquals(t, "Default Protocol", cla.Protocol, "TCP")
	AssertEquals(t, "Default Timeout", cla.Timeout, 5000)
}

func TestParsePorts(t *testing.T) {
	portRange := "134-345"
	ports := parsePorts(portRange)
	AssertEquals(t, "", 2, len(ports))
}

func TestParseSpecifiedPorts(t *testing.T) {

}

func TestValidatePorts(t *testing.T) {
}
