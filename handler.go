package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) handlePacket(c *net.UDPAddr, r *net.UDPConn) {
	var buffer = make([]byte, 0xffff)

	for {
		// Read from server
		n, err := r.Read(buffer[0:])
		log.Printf("reading..... n: %d, b= %x\n", n, buffer[0:n])

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
