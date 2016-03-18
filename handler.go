package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) handlePacketTCP(i int, buf []byte) {
	rConn, err := net.DialTCP("tcp", nil, self.tcp)
	if err != nil {
		log.Fatalln(err)
	}
	defer rConn.Close()

	if _, err := rConn.Write(buf[0:i]); err != nil {
		log.Fatalln(err)
	}

	if self.debug {
		log.Println(string(buf[0:i]))
	}
	return
}

func (self *UDPProxy) handlePacketUDP(i int, buf []byte) {
	rConn, err := net.DialUDP("udp", nil, self.udp)
	if err != nil {
		log.Fatalln(err)
	}
	defer rConn.Close()

	if _, err := rConn.Write(buf[0:i]); err != nil {
		log.Fatalln(err)
	}

	if self.debug {
		log.Println(string(buf[0:i]))
	}
	return
}
