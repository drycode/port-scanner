package argparse

import (
	"flag"
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
	Hosts          []string
	Protocol       string
	Timeout        int
	SpecifiedPorts []int
	AllPorts       []int
	TotalPorts     int
}

func getArgs() UnmarshalledCommandLineArgs {
	hostStringPtr := flag.String("hosts", "127.0.0.1", "Hostname or IP address, local or remote.")
	protocolStringPtr := flag.String("protocol", "TCP", "Specify the protocol for the scanned ports.")
	timeout := flag.Int("timeout", 5000, "Specify the timeout to wait on a port on the server.")
	specifiedPortsPtr := flag.String("ports", "", "A list of specific ports delimited by ','. Can be used w/ or w/o port range.")

	flag.Parse()
	hosts := parseHosts(*hostStringPtr)
	specifiedPorts, _ := parseSpecifiedPorts(*specifiedPortsPtr)
	totalPorts := specifiedPorts
	cla := UnmarshalledCommandLineArgs{
		Hosts:          hosts,
		Protocol:       *protocolStringPtr,
		Timeout:        *timeout,
		SpecifiedPorts: specifiedPorts,
		AllPorts:       totalPorts,
		TotalPorts:     len(totalPorts),
	}
	return cla
}

func parseHosts(ps string) []string {
	hosts := strings.Split(ps, ",")
	for i := 0; i < len(hosts); i++ {
		if i < len(hosts)-1 && hosts[i] == "" {
			hosts[i], hosts[i+1] = hosts[i+1], hosts[i]
		}
	}

	i := len(hosts) - 1
	for i > 0 && hosts[i] == "" {
		i--
	}
	return hosts[0 : i+1]
}

func parsePortRange(ps string) [2]int {
	portsSliceString := strings.Split(ps, "-")
	start, _ := strconv.Atoi(portsSliceString[0])
	end, _ := strconv.Atoi(portsSliceString[1])

	portsSlice := [2]int{start, end}
	return portsSlice
}

func in(char string, s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == char[0] {
			return true
		}
	}
	return false
}

func parseSpecifiedPorts(ps string) ([]int, error) {
	portsSlice := strings.Split(ps, ",")
	var specifiedPorts []int
	var err error

	addPortsFromRange := func(prs string) {
		portRange := parsePortRange(prs)
		logrus.Debug(portRange)
		for i := portRange[0]; i < portRange[1]; i++ {
			specifiedPorts = append(specifiedPorts, i)
		}
	}

	if len(portsSlice) == 0 {
		return make([]int, 0), nil
	}

	for _, portString := range portsSlice {
		portString = strings.TrimSpace(portString)
		if portString == "" {
			continue
		} else if in("-", portString[1:]) {
			addPortsFromRange(portString)
		} else {
			val, err := strconv.Atoi(portString)
			if val < 0 {
				continue
			}
			if err != nil {
				logrus.Errorf("Trouble decoding specified ports of %s", portString)
				return make([]int, 0), err
			}
			specifiedPorts = append(specifiedPorts, val)
		}
	}
	sort.Ints(specifiedPorts)
	return specifiedPorts, err
}

func validatePorts(p string) {
	l, err := regexp.MatchString(`^\d*-\d*$`, p)
	if l == false || err != nil {
		logrus.Fatal("Invalid value passes in port range argument")
		os.Exit(1)
	}
}
