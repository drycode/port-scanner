package tests

import (
	"fmt"
	"testing"

	. "github.com/drypycode/portscanner/argparse"
	. "github.com/drypycode/portscanner/scanner"
	. "github.com/drypycode/portscanner/utils"
)

type TestProgressBar struct{}

func (tpb TestProgressBar) UpdatePercentage(_ int) {}

func TestSupportedConnectionTypes(t *testing.T) {
	type scannerConfigs struct {
		protocol string
		hosts    []string
		allPorts []int
	}
	testValidInput := []scannerConfigs{
		{protocol: "TCP", hosts: []string{"google.com"}, allPorts: []int{80, 443, 13337}},
		{protocol: "UDP", hosts: []string{"localhost"}, allPorts: []int{80, 443, 13337}},
	}

	for _, inputs := range testValidInput {
		output := make(map[string]*SafeSlice)

		scanner := Scanner{
			Config: UnmarshalledCommandLineArgs{
				Hosts:      inputs.hosts,
				Protocol:   inputs.protocol,
				Timeout:    5000,
				AllPorts:   inputs.allPorts,
				TotalPorts: len(inputs.allPorts) * len(inputs.hosts),
			},
			BatchSize: 100,
			Display:   &TestProgressBar{},
		}
		testName := fmt.Sprintf("Hosts: %v, Ports: %v", inputs.hosts, inputs.allPorts)

		scanner.Scan(output)

		for _, ss := range output {
			if len(ss.OpenPorts) == 0 {
				t.Error(testName)
			}
		}

	}
}
