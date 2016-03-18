package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) handlePacketUDP(i int, buf []byte, c *net.UDPAddr) {
	conn, err := net.DialUDP("udp", nil, self.udp)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := conn.Write(buf[0:i]); err != nil {
		log.Printf("Client: %s err: %s", c.String(), err)
		return
	}

	self.txBytes += uint64(i)

	defer conn.Close()

	var buffer = make([]byte, 0xffff)
	for {
		// Read from server
		n, err := conn.Read(buffer[0:])

		if err != nil {
			log.Printf("Client: %s err: %s", c.String(), err)
			return
		}

		self.rxBytes += uint64(n)

		// Relay it to client
		_, err = self.conn.WriteToUDP(buffer[0:n], c)
		if err != nil {
			log.Printf("Client: %s err: %s", c.String(), err)
			return
		}
	}
}
