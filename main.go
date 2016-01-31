package UDPProxy

import (
	"log"
	"net"
)

func handlePacket(i int, addr *net.UDPAddr, buf []byte) {
	log.Println("Received ", string(buf[0:i]), " from ", addr)
	return
}

func Start(bind string) {

	UDPAddr, err := net.ResolveUDPAddr("udp", bind)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", UDPAddr)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		go handlePacket(n, addr, buf)
		if err != nil {
			log.Println("Error: ", err)
		}
	}
}
