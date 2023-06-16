package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	ipPtr := flag.String("ip", "", "IP address to scan")
	flag.Parse()

	if *ipPtr == "" {
		fmt.Println("Please specify an IP address with the -ip argument")
		return
	}

	var wg sync.WaitGroup

	for port := 0; port <= 65535; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", *ipPtr, p)
			conn, err := net.DialTimeout("tcp", address, time.Second)
			if err == nil {
				defer conn.Close()
				fmt.Printf("Port %d is open\n", p)
			}
		}(port)
	}

	wg.Wait()
}
