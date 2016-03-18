package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) handlePacketUDP(i int, buf []byte, c *Client) {
	var err error
	c.conn, err = net.DialUDP("udp", nil, self.udp)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := c.conn.Write(buf[0:i]); err != nil {
		log.Printf("Client: %s err: %s", c.addr.String(), err)
		return
	}

	self.txBytes += uint64(i)

	go self.rw(c)
}
