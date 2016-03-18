package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) handlePacketUDP(i int, buf []byte, client *net.UDPAddr) {
	rConn, err := net.DialUDP("udp", nil, self.udp)
	if err != nil {
		log.Fatalln(err)
	}
	defer rConn.Close()

	if _, err := rConn.Write(buf[0:i]); err != nil {
		log.Fatalln(err)
	}

	if self.debug {
		log.Printf("Size: %d data: \n%s", i, buf[0:i])
	}

	self.txBytes += uint64(i)

	var buffer = make([]byte, 0xffff)
	for {
		// Read from server
		n, err := rConn.Read(buffer[0:])
		if err != nil {
			log.Fatalln(err)
		}
		self.rxBytes += uint64(n)

		// Relay it to client
		_, err = self.conn.WriteToUDP(buffer[0:n], client)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return
}
