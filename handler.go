package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) handlePacket(d []byte, c *net.UDPAddr, b *Backend) {
	if _, err := b.conn.(*net.UDPConn).Write(d); err != nil {
		log.Printf("Client: %s err: %s", c.String(), err)
		return
	}

	var (
		n      int
		err    error
		buffer = make([]byte, 0xffff)
	)

	for {
		// Read from server
		switch v := b.conn.(type) {
		case *net.UDPConn:
			n, err = v.Read(buffer[0:])
		case *net.TCPConn:
			n, err = v.Read(buffer[0:])
		}

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
