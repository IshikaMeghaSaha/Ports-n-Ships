package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type scanResult struct {
	Protocol string
	Port     int
	State    string
}

func scanPort(protocol, host string, port int) scanResult {
	result := scanResult{Port: port, Protocol: protocol}
	address := host + strconv.Itoa(port)

	conn, err := net.DialTimeout(protocol, address, 6000*time.Second)
	if err != nil {
		result.State = "Closed"
		return result
	}

	defer conn.Close()
	result.State = "Open"
	return result
}

func scan(hostname string) []scanResult {
	var results []scanResult

	for i := 1; i <= 65535; i++ {
		results = append(results, scanPort("tcp", hostname, i))
		results = append(results, scanPort("udp", hostname, i))
	}
	return results
}

func main() {
	res := scan("localhost")
	fmt.Println(res)
}
