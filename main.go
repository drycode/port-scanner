package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"

	ap "github.com/drypycode/portscanner/argparse"

	pb "github.com/drypycode/portscanner/progressbar"
	. "github.com/drypycode/portscanner/scanner"
	. "github.com/drypycode/portscanner/utils"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func getBatchSize(totalPorts int) int {
	batchSize, err := GetUlimit()
	if err != nil {
		batchSize = 2000
		logrus.Info(fmt.Sprintf("Trouble locating ulimit on %v...using default batch size 2000", exec.Command("uname -rs")))
	}
	portRangeSize := totalPorts
	if batchSize > portRangeSize {
		batchSize = portRangeSize
	}
	return batchSize
}

func reportOpenPorts(totalPorts int, ss *SafeSlice, timer time.Duration) {
	fmt.Println()
	fmt.Printf("GoScan done: %d ports scanned in %v seconds. \n", totalPorts, math.Round(timer.Seconds()*100)/100)
	for port := range ss.OpenPorts {
		fmt.Println(ss.OpenPorts[port])
	}
	fmt.Println()
}

func welcome() {
	fmt.Println("Starting Golang GoScan v0.1.0 ( github.com/drypycode/portscanner/v0.1.0 ) at", time.Now().Format(time.RFC1123))
	fmt.Println()
}

func main() {
	welcome()
	cliArgs := ap.ParseArgs()
	batchSize := getBatchSize(cliArgs.TotalPorts)
	for _, host := range cliArgs.Hosts {
		fmt.Println("Scanning ports on ", host)
		bar := pb.NewProgressBar(cliArgs.TotalPorts)
		scanner := Scanner{Config: cliArgs, BatchSize: batchSize, Display: &bar}
		openPorts := new(SafeSlice)
		startTime := time.Now()
		scanner.Scan(host, openPorts)
		elapsed := time.Since(startTime)
		reportOpenPorts(cliArgs.TotalPorts, openPorts, elapsed)
	}

}
