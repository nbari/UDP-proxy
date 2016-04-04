package main

import (
	"log"
	"net"
	"time"
)

func close(c *net.UDPConn) {
	log.Printf("closing: %s -> %s", c.LocalAddr().String(), c.RemoteAddr().String())
	c.Close()
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:1514")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer close(conn)

	buf := []byte("hola")
	_, err = conn.Write(buf)
	if err != nil {
		log.Println(err)
	}

	// receive message from server
	buffer := make([]byte, 1400)
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	n, addr, err := conn.ReadFromUDP(buffer)

	log.Printf("Received from UDP server [%s]: size: %d msg: %s", addr.String(), n, buffer[:n])
}
