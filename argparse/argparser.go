package argparse

import (
	"flag"
	"os"
	"regexp"
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

func getArgs() UnmarshalledCommandLineArgs {
	portsStringPtr := flag.String("portrange", "0-3000", "A port range, delimited by '-'. 65535")
	hostStringPtr := flag.String("host", "127.0.0.1", "Hostname or IP address, local or remote.")
	protocolStringPtr := flag.String("protocol", "TCP", "Specify the protocol for the scanned ports.")
	timeout := flag.Int("timeout", 5000, "Specify the timeout to wait on a port on the server.")
	specifiedPortsPtr := flag.String("portlist", "", "A list of specific ports delimited by ','. Can be used w/ or w/o port range.")

	flag.Parse()
	portRange := parsePorts(*portsStringPtr)
	specifiedPorts := parseSpecifiedPorts(*specifiedPortsPtr)

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

func parseSpecifiedPorts(ps string) []int {
	portsSliceString := strings.Split(ps, ",")
	specifiedPorts := make([]int, 0, len(portsSliceString))
	if portsSliceString[0] != "" {
		for i := range portsSliceString {
			val, err := strconv.Atoi(portsSliceString[i])
			if err != nil {
				logrus.Fatal("Trouble decoding specified ports")
			}
			specifiedPorts[i] = val
		}
	}
	return specifiedPorts
}

func validatePorts(p string) {
	l, err := regexp.MatchString(`^\d*-\d*$`, p)
	if l == false || err != nil {
		logrus.Fatal("Invalid value passes in port range argument")
		os.Exit(1)
	}
}
