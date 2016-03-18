package UDPProxy

import (
	"log"
)

func (self *UDPProxy) rw(c *Client) {
	var buffer = make([]byte, 0xffff)
	for {
		// Read from server
		defer c.conn.Close()
		n, err := c.conn.Read(buffer[0:])

		if err != nil {
			log.Printf("Client: %s err: %s", c.addr.String(), err)
			return
		}

		self.rxBytes += uint64(n)

		// Relay it to client
		_, err = self.conn.WriteToUDP(buffer[0:n], c.addr)
		if err != nil {
			log.Printf("Client: %s err: %s", c.addr.String(), err)
			return
		}
	}

	return
}
