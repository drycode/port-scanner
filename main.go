package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	ap "github.com/drypycode/port-scanner/argparse"
	. "github.com/drypycode/port-scanner/portscanner"
	pb "github.com/drypycode/port-scanner/progressbar"
	"github.com/sirupsen/logrus"
)

// GetUlimit ...
func getUlimit() int {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}
	ulimit := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(ulimit, 10, 64)

	if err != nil {
		panic(err)
	}
	return int(i)
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func getBatchSize(totalPorts int) int{
	batchSize := getUlimit()
	portRangeSize := totalPorts
	if batchSize > portRangeSize{
		batchSize = portRangeSize
	}		
	return batchSize
}

func reportOpenPorts(openPorts []string) {
	for port := range openPorts {
		fmt.Println(openPorts[port])
	}
}

func main() {
	logrus.Info("Scanning ports...")
	cliArgs := ap.ParseArgs()
	totalPorts := cliArgs.PortRange[1] - cliArgs.PortRange[0] 
	batchSize := getBatchSize(totalPorts)
	bar := pb.NewProgressBar(totalPorts)
	scanner := Scanner{Config:cliArgs, BatchSize: batchSize}
	
	openPorts := make([]string, 2)

	scanner.Scan(&bar, &openPorts)
	reportOpenPorts(openPorts)
}
