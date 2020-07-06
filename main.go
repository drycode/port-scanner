package main

import (
	"fmt"
	"net"
	"strconv"
)

type progressBar struct {
	portRange  [2]int
	current    int
	percentage float64
}

func main() {
	// dial tcp 127.0.0.1:8000: socket: too many open files
	// This ^ error is occurring when trying to check a larger range of ports
	portRange := [2]int{0, 9000}
	pb := progressBar{portRange, 0, 0}

	c := make(chan string)
	for i := portRange[0]; i < portRange[1]; i += 100 {
		batchCalls([2]int{i, i + 100}, c)
		pb.setPercentage(float64(i - portRange[0] + 100))
		fmt.Println(strconv.FormatInt(int64(pb.percentage*100), 10) + "%")
	}
}

func (p *progressBar) setPercentage(i float64) {
	perc := i / float64((*p).portRange[1]-(*p).portRange[0])
	(*p).percentage = perc
}

// Gives go time to close up some of the open connections
func batchCalls(r [2]int, c chan string) {
	if r[1]-r[0] >= 250 {
		fmt.Println("WARNING: Your batch size might exceed your ulimit for open files.")
	}
	for i := r[0]; i < r[1]; i++ {
		go pingServerPort(i, c)
	}

	j := r[0]

	for l := range c {
		j++
		go func(l string) {
			if l != "." {
				fmt.Println(l)
			}
		}(l)
		if j >= r[1] {
			return
		}
	}
}

func pingServerPort(p int, c chan string) {
	port := strconv.FormatInt(int64(p), 10)
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)

	if err == nil {
		c <- "Port " + port + " is open"
		conn.Close()
		return
	}
	c <- "."
	return
}
