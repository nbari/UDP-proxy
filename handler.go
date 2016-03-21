package UDPProxy

import (
	"log"
)

func (self *UDPProxy) handlePacket() {
	defer self.rconn.Close()
	var buffer = make([]byte, 1500)
	for {
		// Read from server
		n, err := self.rconn.Read(buffer[0:])
		print("reading...")
		if err != nil {
			log.Printf("READ - client: %s err: %s", self.caddr.String(), err)
			return
		}
		self.rxBytes += uint64(n)
		log.Printf("reading size %d rxBytes: %d \n", n, self.rxBytes)

		// Relay it to client
		_, err = self.lconn.WriteToUDP(buffer[0:n], self.caddr)
		print("writing...")
		if err != nil {
			log.Printf("WRITE - client: %s err: %s", self.caddr.String(), err)
			return
		}
	}
}
