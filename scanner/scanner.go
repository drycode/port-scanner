package scanner

import (
	"fmt"
	"net"
	"os"
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
	Display   Progress
	scanned   int
}

type hostPortPair struct {
	host string
	port int
}

func buildHostPortPairs(hosts []string, allPortsToScan []int) []hostPortPair {
	var pairs []hostPortPair
	for _, host := range hosts {
		for _, port := range allPortsToScan {
			pairs = append(pairs, hostPortPair{host, port})
		}
	}
	return pairs
}

func (s *Scanner) PreScanCheck() {
	total := int64(len((*s).Config.Hosts)) * int64((*s).Config.TotalPorts)
	if total > 300_000 {
		fmt.Printf("It looks like you have selected a considerable amount of hosts / ports to scan (total: %v). Would you like to continue?\n", total)

		// var then variable name then variable type
		var input string

		// Taking input from user
		fmt.Scanln(&input)
		for _, val := range []string{"yes", "y", "Y"} {
			if strings.Compare(input, val) == 0 {
				return
			}
		}
		os.Exit(0)
	}
}

// Scan ...
func (s *Scanner) Scan(openPorts map[string]*SafeSlice) {
	(*s).scanned = 0
	fmt.Println("Scanning ports on ", (*s).Config.Hosts)
	allPortsToScan := (*s).Config.AllPorts
	hostPortPairs := buildHostPortPairs((*s).Config.Hosts, allPortsToScan)
	totalPorts := len(hostPortPairs)

	for batchStart := 0; batchStart < totalPorts; batchStart += (*s).BatchSize {
		start := batchStart
		var end int
		if (batchStart + (*s).BatchSize) < totalPorts {
			end = (batchStart + (*s).BatchSize)
		} else {
			end = totalPorts
		}
		(*s).BatchCalls(hostPortPairs[start:end], openPorts)
	}
}

// PingServerPort ...
func (s *Scanner) PingServerPort(host string, p int, c chan hostPortPair) {

	port := strconv.FormatInt(int64(p), 10)
	conn, err := net.DialTimeout(
		strings.ToLower((*s).Config.Protocol),
		host+":"+port,
		time.Duration(int64((*s).Config.Timeout))*time.Millisecond,
	)

	if err == nil {
		port, _ := strconv.Atoi(port)
		c <- hostPortPair{host: host, port: port}
		conn.Close()
		return
	}
	c <- hostPortPair{host: ".", port: 0}
	return
}

// BatchCalls ...
func (s *Scanner) BatchCalls(hpps []hostPortPair, ops map[string]*SafeSlice) {
	c := make(chan hostPortPair)
	totalPorts := len(hpps)
	scannedInBatch := 0
	var logFromChannel = func(c chan hostPortPair) {
		wg := sync.WaitGroup{}
		for hpp := range c {
			(*s).scanned++
			(*s).Display.UpdatePercentage((*s).scanned)
			scannedInBatch++
			go func(hpp hostPortPair, openPorts map[string]*SafeSlice) {
				if hpp.host != "." {
					wg.Add(1)
					if _, ok := openPorts[hpp.host]; !ok {
						openPorts[hpp.host] = new(SafeSlice)
					}
					openPorts[hpp.host].Append(strconv.Itoa(hpp.port), &wg)
				}

			}(hpp, ops)

			if scannedInBatch >= totalPorts {
				return
			}
		}
	}

	var pingPorts = func(c chan hostPortPair) {
		for _, hpp := range hpps {
			go (*s).PingServerPort(hpp.host, hpp.port, c)
		}
	}

	pingPorts(c)
	logFromChannel(c)
	close(c)
}
