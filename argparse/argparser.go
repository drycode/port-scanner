package argparse

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// ParseArgs ...
func ParseArgs() UnmarshalledCommandLineArgs {
	cliArgs := getArgs()
	logrus.Debug(cliArgs)
	return cliArgs
}

// UnmarshalledCommandLineArgs ...
type UnmarshalledCommandLineArgs struct {
	PortRange      [2]int
	Host           string
	Protocol       string
	Timeout        int
	SpecifiedPorts []int
}

func (ucla *UnmarshalledCommandLineArgs) GetAllPorts() []int {
	var allPorts []int
	for i := ucla.PortRange[0]; i < ucla.PortRange[1]; i++ {
		allPorts = append(allPorts, i)
	}
	allPorts = append(allPorts, ucla.SpecifiedPorts...)
	sort.Ints(allPorts)
	return allPorts
}

func getArgs() UnmarshalledCommandLineArgs {
	portsStringPtr := flag.String("portrange", "0-0", "A port range, delimited by '-'. 65535")
	hostStringPtr := flag.String("host", "127.0.0.1", "Hostname or IP address, local or remote.")
	protocolStringPtr := flag.String("protocol", "TCP", "Specify the protocol for the scanned ports.")
	timeout := flag.Int("timeout", 5000, "Specify the timeout to wait on a port on the server.")
	specifiedPortsPtr := flag.String("portlist", "", "A list of specific ports delimited by ','. Can be used w/ or w/o port range.")

	flag.Parse()
	portRange := parsePorts(*portsStringPtr)
	specifiedPorts, _ := parseSpecifiedPorts(*specifiedPortsPtr)

	cla := UnmarshalledCommandLineArgs{portRange, *hostStringPtr, *protocolStringPtr, *timeout, specifiedPorts}
	return cla
}

func parsePorts(ps string) [2]int {
	portsSliceString := strings.Split(ps, "-")
	start, _ := strconv.Atoi(portsSliceString[0])
	end, _ := strconv.Atoi(portsSliceString[1])

	portsSlice := [2]int{start, end}
	return portsSlice
}

func parseSpecifiedPorts(ps string) ([]int, error) {
	portsSlice := strings.Split(ps, ",")
	var specifiedPorts []int
	var err error
	for _, port := range portsSlice {
		if port == "" {
			continue
		}
		val, err := strconv.Atoi(port)
		if err != nil {
			logrus.Error("Trouble decoding specified ports")
			return make([]int, 0), err
		}
		specifiedPorts = append(specifiedPorts, val)
	}

	fmt.Println(specifiedPorts)
	return specifiedPorts, err
}

func validatePorts(p string) {
	l, err := regexp.MatchString(`^\d*-\d*$`, p)
	if l == false || err != nil {
		logrus.Fatal("Invalid value passes in port range argument")
		os.Exit(1)
	}
}
