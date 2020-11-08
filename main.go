package main

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"time"

	ap "github.com/drypycode/port-scanner/argparse"
	ps "github.com/drypycode/port-scanner/portscanner"
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
	logrus.SetLevel(logrus.InfoLevel)
}

func printInitialization(protocol string, host string) {
	fmt.Println("Starting Golang GoScan v0.1.0 ( github.com/drypycode/port-scanner/v0.1.0 ) at", time.Now().Format(time.RFC1123))
	fmt.Println("Scanning ports on", host)
	
}

func getBatchSize(totalPorts int) int{
	batchSize := getUlimit()
	portRangeSize := totalPorts
	if batchSize > portRangeSize{
		batchSize = portRangeSize
	}		
	return batchSize
}

func reportOpenPorts(totalPorts int, ss *ps.SafeSlice, timer time.Duration) {
	fmt.Println()
	fmt.Printf("GoScan done: %d ports scanned in %v seconds. \n", totalPorts,  math.Round(timer.Seconds()*100) / 100)
	fmt.Println()
	for port := range ss.OpenPorts {
		fmt.Println(ss.OpenPorts[port])
	}
}



func main() {
	
	cliArgs := ap.ParseArgs()
	printInitialization(cliArgs.Protocol,cliArgs.Host)	
	totalPorts := cliArgs.PortRange[1] - cliArgs.PortRange[0] 
	logrus.Debug(totalPorts)
	batchSize := getBatchSize(totalPorts)
	bar := pb.NewProgressBar(totalPorts)
	scanner := ps.Scanner{Config:cliArgs, BatchSize: batchSize, Display: &bar}
	
	openPorts := ps.SafeSlice{}
	
	startTime := time.Now()
	scanner.Scan(&openPorts)
	elapsed := time.Since(startTime)

	reportOpenPorts(totalPorts, &openPorts, elapsed)
}
