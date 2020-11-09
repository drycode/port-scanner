package portscanner

import (
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	ap "github.com/drypycode/port-scanner/argparse"
	. "github.com/drypycode/port-scanner/progressbar"
)

// SafeSlice ...
type SafeSlice struct {
	mu 			sync.RWMutex
	OpenPorts 	[]string
}

func (ss *SafeSlice) length() int {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	return len(ss.OpenPorts)
}

func (ss *SafeSlice) append(val string, wg *sync.WaitGroup) {
	ss.mu.Lock()
	ss.OpenPorts = append((*ss).OpenPorts, val)
	ss.mu.Unlock()	
	wg.Done()
}

// Scanner ...
type Scanner struct {
	Config		ap.UnmarshalledCommandLineArgs
	BatchSize	int
	Display		*ProgressBar
	scanned 	int
}

// Scan ...
func (s *Scanner) Scan(openPorts *SafeSlice) {
	(*s).scanned = 0
	for batchStart := (*s).Config.PortRange[0]; batchStart < (*s).Config.PortRange[1]; batchStart += (*s).BatchSize {
		start := batchStart
		var end int
		if (batchStart + (*s).BatchSize) < (*s).Config.PortRange[1] {
			end = (batchStart + (*s).BatchSize)
		} else {
			end = (*s).Config.PortRange[1]
		}
		(*s).BatchCalls(start, end, openPorts)
	}
}

// PingServerPort ...
func (s *Scanner) PingServerPort(p int, c chan string) {
	
	port := strconv.FormatInt(int64(p), 10)
	conn, err := net.DialTimeout(
		strings.ToLower((*s).Config.Protocol), 
		(*s).Config.Host + ":" + port, 
		time.Duration(int64((*s).Config.Timeout)) * time.Millisecond,
	)

	if err == nil {
		c <- "Port " + port + " is open"
		conn.Close()
		return
	}
	c <- "."
	return
}


// BatchCalls ...
func (s *Scanner) BatchCalls(start int, end int, ops *SafeSlice) {
	c := make(chan string)
	// var start = batchStart
	// var end = (*s).Config.PortRange[1]
	scannedInBatch := 0
	var logFromChannel = func (c chan string) {
		wg := sync.WaitGroup{}
		for l := range c {
			(*s).scanned++
			(*s).Display.UpdatePercentage((*s).scanned)
			scannedInBatch++
			go func(l string, openPorts *SafeSlice) {
				if l != "." {
					wg.Add(1)
					openPorts.append(l, &wg)
				}
				
			}(l, ops)
					
			if scannedInBatch >= end - start {
				return
			}
		}
		// wg.Wait()
		
	}
	
	var pingPorts = func(c chan string) {
		for port := start; port < end; port++ {
			go (*s).PingServerPort(port, c)
		}
	}

	pingPorts(c)
	logFromChannel(c)
	close(c)
}