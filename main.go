package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type progressBar struct {
	portRange     [2]int
	current       int
	percentage    float64
	lastDisplayed string
}

func (p *progressBar) setPercentage(i float64) {
	perc := i / float64((*p).portRange[1]-(*p).portRange[0])
	(*p).percentage = perc
}

// Gives go time to close up some of the open connections
func batchCalls(r [2]int, c chan string, pb *progressBar) {
	for i := r[0]; i < r[1]; i++ {
		go pingServerPort(i, c)
	}
	j := r[0]
	for l := range c {
		j++
		go func(l string, j int) {
			if l != "." {
				fmt.Println(l)
			}
		}(l, j)
		percentageHelper(pb, float64(j-r[0]))
		if j >= r[1] {
			return
		}
	}

}

func pingServerPort(p int, c chan string) {
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

func getUlimit() int {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}
	return int(i)
}

func percentageHelper(pb *progressBar, n float64) {
	pb.setPercentage(n)
	if strconv.FormatInt(int64((*pb).percentage*100), 10) != (*pb).lastDisplayed {
		(*pb).lastDisplayed = strconv.FormatInt(int64((*pb).percentage*100), 10)
		fmt.Println((*pb).lastDisplayed + "%")
	}

}

func main() {
	// dial tcp 127.0.0.1:8000: socket: too many open files
	// This ^ error is occurring when trying to check a larger range of ports
	portRange := [2]int{0, 9000}
	pb := progressBar{portRange, 0, 0, ""}
	// ulimit := getUlimit()

	var b int
	// if ulimit > portRange[1] {
	// 	b = portRange[1]
	// } else {
	// 	b = ulimit
	// }
	b = 1000
	c := make(chan string)

	for i := portRange[0]; i < portRange[1]; i += b {
		batchCalls([2]int{i, i + b}, c, &pb)
	}

}
