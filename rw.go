package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) rw(c *Client) {
	var (
		buffer = make([]byte, 0xffff)
		err    error
		n      int
	)

	for {
		// Read from server
		switch v := c.conn.(type) {
		case *net.UDPConn:
			defer v.Close()
			n, err = v.Read(buffer[0:])
		case *net.TCPConn:
			defer v.Close()
			n, err = v.Read(buffer[0:])
		}

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
