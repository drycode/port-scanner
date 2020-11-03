package main

import (
	"fmt"

	. "github.com/drypycode/port-scanner/portscanner"
	"github.com/drypycode/port-scanner/progressbar"
	log "github.com/sirupsen/logrus"
)

type progressBar = progressbar.ProgressBar


type scannerConfig struct {
	portRange [2]int
	ulimit int 
}


// Gives go time to close up some of the open connections
func batchCalls(r [2]int, pb *progressBar) {
	c := make(chan string)
	var start = r[0]
	var end = r[1]
	for port := start; port < end; port++ {
		go PingServerPort(port, c)
	}
	
	scanned := 0
	
	for l := range c {
		scanned++
		progressbar.PercentageHelper(pb, scanned-start)
		go func(l string) {
			if l != "." {
				fmt.Println(l)
			}
			
		}(l)
		
		// log.Info(scanned)
		
		if scanned >= end - start {
			return
		}
	}

}


func main() {
	// dial tcp 127.0.0.1:8000: socket: too many open files
	// This ^ error is occurring when trying to check a larger range of ports
	log.Info("Scanning ports...")
	portRange := [2]int{0, 8000}
	config := scannerConfig{
		portRange, 
		// 200,
		GetUlimit(portRange[1] - portRange[0]),
	}
	bar := progressBar{
		TotalPorts: config.portRange[1] - config.portRange[0], 
		Current: 0, 
		Percentage: 0, 
		LastDisplayed: "",
	}
	
	

	for batchStart := config.portRange[0]; batchStart < config.portRange[1]; batchStart += config.ulimit {
		batchCalls([2]int{batchStart, batchStart + config.ulimit}, &bar)
	}

}
