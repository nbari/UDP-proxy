package UDPProxy

import (
	"log"
	"net"
)

func handlePacket(i int, addr *net.UDPAddr, buf []byte, tcp *net.TCPAddr) {
	rConn, err := net.DialTCP("tcp", nil, tcp)
	if err != nil {
		panic(err)
	}
	defer rConn.Close()

	if _, err := rConn.Write(buf[0:i]); err != nil {
		panic(err)
	}
	log.Printf("sent:\n%s", buf[0:i])
	return
}

func Start(bind string, tcp *net.TCPAddr) {
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
		if err != nil {
			log.Println("Error: ", err)
		} else {
			go handlePacket(n, addr, buf, tcp)
		}
	}
}
