package main

import (
	"log"
	"net"
)

type Packet struct {
	addr *net.UDPAddr
	data []byte
}

const UDP_PACKET_SIZE = 1500

func send(conn *net.UDPConn, outbound chan Packet) {
	for packet := range outbound {
		_, err := conn.WriteToUDP(packet.data, packet.addr)
		if err != nil {
			log.Println("Error on write: ", err)
			continue
		}
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":1514")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	inbound := make(chan Packet)

	go send(conn, inbound)

	var connections uint64
	for {
		b := make([]byte, UDP_PACKET_SIZE)
		n, addr, err := conn.ReadFromUDP(b)
		if err != nil {
			log.Printf("Error: UDP read error: %v", err)
			continue
		}
		connections++
		log.Printf("Connections: %d read: %d client: %s", connections, n, addr.String())
		inbound <- Packet{addr, b[:n]}
	}
}
