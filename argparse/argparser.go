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
func ParseArgs()(ports [2]int){
	ports = getArgs()
	logrus.Debug(ports)
	return 
}

func getArgs() [2]int {
	portsStringPtr := flag.String("ports", "0-3000", "A port range, separated by '-'. Defaults: 0-65000")
	flag.Parse()
	return parsePorts(*portsStringPtr)
}

func parsePorts(ps string) [2]int {
	validatePorts(ps)
	portsSliceString := strings.Split(ps, "-")
	start, _ := strconv.Atoi(portsSliceString[0]) 
	end, _ := strconv.Atoi(portsSliceString[1])
	
	portsSlice := [2]int{start, end}
	return portsSlice
}

func validatePorts(p string) {
	l, err := regexp.MatchString(`^\d*-\d*$`, p)
	if l == false || err != nil {
		logrus.Fatal("Invalid value passes in port range argument")
		os.Exit(1)
	}
}