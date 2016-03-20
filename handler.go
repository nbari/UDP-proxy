package UDPProxy

import (
	"log"
)

func (self *UDPProxy) handlePacket() {
	defer self.rconn.Close()
	var buffer = make([]byte, 0xffff)
	for {
		// Read from server
		n, err := self.rconn.Read(buffer[0:])
		log.Printf("reading..... n: %d, b= %x\n", n, buffer[0:n])

		if err != nil {
			log.Printf("Client: %s err: %s", self.caddr.String(), err)
			return
		}

		self.rxBytes += uint64(n)

		// Relay it to client
		_, err = self.lconn.WriteToUDP(buffer[0:n], self.caddr)
		if err != nil {
			log.Printf("Client: %s err: %s", self.caddr.String(), err)
			return
		}
	}
}
