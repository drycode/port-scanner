package portscanner

import (
	"net"
	"strconv"
	"strings"
	"time"

	ap "github.com/drypycode/port-scanner/argparse"
	"github.com/drypycode/port-scanner/progressbar"
	. "github.com/drypycode/port-scanner/progressbar"
)

// Scanner ...
type Scanner struct {
	Config		ap.UnmarshalledCommandLineArgs
	BatchSize	int
}

// Scan ...
func (s *Scanner) Scan(bar *ProgressBar, openPorts *[]string) {
	for batchStart := (*s).Config.PortRange[0]; batchStart < (*s).Config.PortRange[1]; batchStart += (*s).BatchSize {
		s.BatchCalls(bar, openPorts)
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
func (s *Scanner) BatchCalls(pb *ProgressBar, ops *[]string) {
	c := make(chan string)
	var start = (*s).Config.PortRange[0]
	var end = (*s).Config.PortRange[1]
	
	
	var logFromChannel = func (c chan string) {
		scanned := 0
		for l := range c {
			scanned++
			
			go func(l string, openPorts *[]string) {
				if l != "." {
					// fmt.Println(l)
					*openPorts = append(*openPorts, l)
				}
				
			}(l, ops)
			progressbar.PercentageHelper(pb, scanned-start)		
			if scanned >= end - start {
				return
			}
		}
		
	}
	
	var pingPorts = func(c chan string) {
		for port := start; port < end; port++ {
			go s.PingServerPort(port, c)
		}
	}

	pingPorts(c)
	logFromChannel(c)
	close(c)
}