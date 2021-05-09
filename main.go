package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	ap "github.com/drypycode/portscanner/argparse"
	"github.com/drypycode/portscanner/ssh"

	pb "github.com/drypycode/portscanner/progressbar"
	. "github.com/drypycode/portscanner/scanner"
	. "github.com/drypycode/portscanner/ssh"
	. "github.com/drypycode/portscanner/utils"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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

type output struct {
	Host  string
	Ports []string
}

func printJSON(ops map[string]*SafeSlice, f os.File) {
	var final []output
	for host, ss := range ops {
		hostPorts := &output{
			Host:  host,
			Ports: ss.OpenPorts,
		}
		final = append(final, *hostPorts)
	}
	marshalled, _ := json.MarshalIndent(final, "", "  ")
	f.Write(marshalled)
}

func printResultToFile(ops map[string]*SafeSlice, path string) {
	f, err := os.Create(path)
	check(err)
	if filepath.Ext(path) == ".json" {
		printJSON(ops, *f)
	} else {
		for host, ss := range ops {
			for port := range ss.OpenPorts {
				f.WriteString(fmt.Sprintf("%s:%v\n", host, ss.OpenPorts[port]))
			}
		}
	}

}

func reportOpenPorts(op map[string]*SafeSlice, timer time.Duration, cliArgs ap.UnmarshalledCommandLineArgs) {
	fmt.Println()
	fmt.Printf("GoScan done: %d ports scanned in %v seconds. \n", cliArgs.TotalPorts, math.Round(timer.Seconds()*100)/100)
	fmt.Println()

	if strings.Compare(cliArgs.FilePath, "") != 0 {
		printResultToFile(op, cliArgs.FilePath)
	} else {
		fmt.Println("Open Ports")
		for host, ss := range op {
			for port := range ss.OpenPorts {
				fmt.Printf("%s:%v\n", host, ss.OpenPorts[port])
			}
		}
	}
	fmt.Println()
}

func welcome(hosts []string) {
	fmt.Println("Starting Golang GoScan v0.1.0 ( github.com/drypycode/portscanner/v0.1.0 ) at", time.Now().Format(time.RFC1123))
	fmt.Println()
}

func main() {
	cliArgs := ap.ParseArgs()
	batchSize := getBatchSize(cliArgs.TotalPorts)
	hosts := cliArgs.Hosts
	if cliArgs.Jump {
		// Tunnel and runCommand on remote server
		Jump(ssh.SSHConfig{Key: cliArgs.PrivateKey, User: cliArgs.RemoteUser, RemoteHost: cliArgs.RemoteHost, Port: "22"})
	} else {
		welcome(hosts)
		bar := pb.NewProgressBar(cliArgs.TotalPorts)
		scanner := Scanner{Config: cliArgs, BatchSize: batchSize, Display: &bar}
		finalReport := make(map[string]*SafeSlice)
		scanner.PreScanCheck()
		startTime := time.Now()
		scanner.Scan(finalReport)
		elapsed := time.Since(startTime)
		reportOpenPorts(finalReport, elapsed, cliArgs)
	}

}
