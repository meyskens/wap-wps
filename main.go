package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/meyskens/uintvar"
)

func main() {
	fmt.Println("Enjoy nostalgia gone too far")

	udpAddress, err := net.ResolveUDPAddr("udp", ":9201")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		panic(err)
	}

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buffer)
		go parseRequest(buffer[:n], addr)
		if err != nil {
			go parseRequest(buffer[:n], addr)
		}
	}
}

func parseRequest(buffer []byte, addr net.Addr) {
	fmt.Println(buffer)
	c := 0
	// [0] == 00
	c++

	subPDULength := getSubPDULength(buffer[c:(c + 2)])
	c += 2
	fmt.Println("subpdulength", subPDULength)

	// 4 bytes Wireless Transaction Protocol, no use for now
	c += 4

	wspHeader, n := parseWSPHeaders(buffer[c:])
	c += n

	data, _ := json.Marshal(wspHeader)
	fmt.Print(string(data))

	//subPULength2 := getSubPDULength(buffer[c:(c + 1)])
	c++

	// Ignore the next WTP Header for now
	c += 4

	verb := verbs[buffer[c]]
	c++
	urlLength, n, _ := uintvar.Parse(buffer[c:])
	c += n

	url := string(buffer[c:(c + int(urlLength))])
	c += int(urlLength)

	fmt.Println(verb, url)
}
