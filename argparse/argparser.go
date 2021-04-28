package argparse

import (
	"flag"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	. "github.com/drypycode/portscanner/utils"
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
	Hosts      []string
	Protocol   string
	Timeout    int
	AllPorts   []int
	TotalPorts int
}

func getArgs() UnmarshalledCommandLineArgs {
	hostStringPtr := flag.String(
		"hosts",
		"127.0.0.1", "A list DNS names or IP addresses (local or remote) delimited by ','. Additionally, for IP addresses the user can provide a valid \nCIDR notation block, and the range of IPs defined in that block will be scanned. \nEx. '127.0.0.1, www.google.com, 192.0.0.0/24, 100.0.0.0-100.0.1.0' \n\nWARNING: A large range of IP addresses compounds exponentially against the list of ports to scan. \n10 hosts @ 10k ports == 100k total scans\n",
	)
	protocolStringPtr := flag.String("protocol", "TCP", "Specify the protocol for the scanned ports.")
	timeout := flag.Int("timeout", 5000, "Specify the timeout to wait on a port on the server.")
	specifiedPortsPtr := flag.String("ports", "", "A list of specific ports delimited by ','. Optionally: A range of ports can be provided in addition to to comma delimited \nspecific ports.\nEx. '80, 443, 100-200, 6543'")

	flag.Parse()
	hosts := parseHosts(*hostStringPtr)
	allPorts, _ := parseSpecifiedPorts(*specifiedPortsPtr)

	cla := UnmarshalledCommandLineArgs{
		Hosts:      hosts,
		Protocol:   *protocolStringPtr,
		Timeout:    *timeout,
		AllPorts:   allPorts,
		TotalPorts: len(allPorts),
	}
	return cla
}

func stripWhitespaceFromSliceOfStrings(sof []string) []string {
	for i := 0; i < len(sof); i++ {
		if i < len(sof)-1 && sof[i] == "" {
			sof[i], sof[i+1] = sof[i+1], sof[i]
		}
	}
	i := len(sof) - 1
	for i > 0 && sof[i] == "" {
		i--
	}
	return sof[0 : i+1]
}

func parseHosts(ps string) []string {
	hosts := strings.Split(ps, ",")

	for i, host := range hosts {
		if in("-", host) {
			ips := ParseIPRange(host)
			hosts[i] = ""
			hosts = append(hosts, ips...)
		}
	}
	hosts = stripWhitespaceFromSliceOfStrings(hosts)
	sort.Strings(hosts)
	return hosts
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

func validatePorts(p string) {
	l, err := regexp.MatchString(`^\d*-\d*$`, p)
	if l == false || err != nil {
		logrus.Fatal("Invalid value passes in port range argument")
		os.Exit(1)
	}
}
