package main

import (
	"log"
	"net"
)

func main() {
	UDPAddr, err := net.ResolveUDPAddr("udp", ":5514")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", UDPAddr)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		log.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			log.Println("Error: ", err)
		}
	}

}
