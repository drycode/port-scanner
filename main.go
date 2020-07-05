package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	// dial tcp 127.0.0.1:8000: socket: too many open files
	// This ^ error is occurring when trying to check a larger range of ports
	portRange := [2]int{5755, 8001}
	c := make(chan string)
	for i := portRange[0]; i < portRange[1]; i++ {
		go pingServerPort(i, c)
		// Gives go time to close up some of the open connections ?
		if i%200 == 0 {
			time.Sleep(time.Second)
		}

	}

	j := portRange[0]
	for l := range c {
		j++
		go func(l string) {
			if l != "." {
				fmt.Println(l)
			}
		}(l)
		if j >= portRange[1] {
			return
		}
	}

}

func pingServerPort(p int, c chan string) {
	port := strconv.FormatInt(int64(p), 10)
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)
	if p == 8000 {
		fmt.Println(port)
		fmt.Println(conn)
		fmt.Println(err)
	}
	if err == nil {
		c <- "\nPort " + port + " is open"
		conn.Close()
		return
	}
	c <- "."
	return
}
