package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5514")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := conn.Write(buf)
		if err != nil {
			log.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
