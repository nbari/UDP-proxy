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
		log.Println(err)
		return
	}

	self.txBytes += uint64(i)

	return
}
