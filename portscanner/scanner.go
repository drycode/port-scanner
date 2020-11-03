package portscanner

import (
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// PingServerPort ...
func PingServerPort(p int, c chan string) {
	port := strconv.FormatInt(int64(p), 10)
	conn, err := net.DialTimeout("tcp", "127.0.0.1:"+port, 5000*time.Millisecond)

	if err == nil {
		c <- "Port " + port + " is open"
		conn.Close()
		return
	}
	c <- "."
	return
}

// GetUlimit ...
func GetUlimit(ports int) int {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}
	if int(i) > ports{
		return ports
	}		
	return int(i)
}
