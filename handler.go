package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) handlePacketTCP(i int, buf []byte) {
	rConn, err := net.DialTCP("tcp", nil, self.tcp)
	if err != nil {
		panic(err)
	}
	defer rConn.Close()

	if _, err := rConn.Write(buf[0:i]); err != nil {
		panic(err)
	}

	if self.debug {
		log.Printf("sent:\n%s", buf[0:i])
	}
	return
}

func (self *UDPProxy) handlePacketUDP(i int, buf []byte) {
	rConn, err := net.DialUDP("tcp", nil, self.udp)
	if err != nil {
		panic(err)
	}
	defer rConn.Close()

	if _, err := rConn.Write(buf[0:i]); err != nil {
		panic(err)
	}

	if self.debug {
		log.Printf("sent:\n%s", buf[0:i])
	}
	return
}
