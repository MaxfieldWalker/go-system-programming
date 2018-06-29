package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address)
	// defer listener.Close()
	fmt.Println("aaa")
	buffer := make([]byte, 1500)
	fmt.Println("aaa")
	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		fmt.Println("aaa")
		if err != nil {
			panic(err)
		}
		fmt.Println("aaa")
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now  %s\n", string(buffer[:length]))
	}
}
