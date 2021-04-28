package scanner

import (
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	ap "github.com/drypycode/portscanner/argparse"
	. "github.com/drypycode/portscanner/progressbar"
	. "github.com/drypycode/portscanner/utils"
)

// Scanner ...
type Scanner struct {
	Config    ap.UnmarshalledCommandLineArgs
	BatchSize int
	Display   *ProgressBar
	scanned   int
}

// Scan ...
func (s *Scanner) Scan(host string, openPorts *SafeSlice) {
	(*s).scanned = 0

	allPortsToScan := (*s).Config.AllPorts
	totalPorts := len(allPortsToScan)
	for batchStart := 0; batchStart < totalPorts; batchStart += (*s).BatchSize {
		start := batchStart
		var end int
		if (batchStart + (*s).BatchSize) < totalPorts {
			end = (batchStart + (*s).BatchSize)
		} else {
			end = totalPorts
		}
		(*s).BatchCalls(host, allPortsToScan[start:end], openPorts)
	}
}

// PingServerPort ...
func (s *Scanner) PingServerPort(host string, p int, c chan string) {

	port := strconv.FormatInt(int64(p), 10)
	conn, err := net.DialTimeout(
		strings.ToLower((*s).Config.Protocol),
		host+":"+port,
		time.Duration(int64((*s).Config.Timeout))*time.Millisecond,
	)

	if err == nil {
		c <- port
		conn.Close()
		return
	}
	c <- "."
	return
}

// BatchCalls ...
func (s *Scanner) BatchCalls(host string, ports []int, ops *SafeSlice) {
	c := make(chan string)
	totalPorts := len(ports)
	scannedInBatch := 0
	var logFromChannel = func(c chan string) {
		wg := sync.WaitGroup{}
		for l := range c {
			(*s).scanned++
			(*s).Display.UpdatePercentage((*s).scanned)
			scannedInBatch++
			go func(l string, openPorts *SafeSlice) {
				if l != "." {
					wg.Add(1)
					openPorts.Append(l, &wg)
				}

			}(l, ops)

			if scannedInBatch >= totalPorts {
				return
			}
		}
	}

	var pingPorts = func(c chan string) {
		for _, port := range ports {
			go (*s).PingServerPort(host, port, c)
		}
	}

	pingPorts(c)
	logFromChannel(c)
	close(c)
}
