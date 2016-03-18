package UDPProxy

import (
	"log"
	"net"
)

func (self *UDPProxy) handlePacketTCP(i int, buf []byte, c *Client) {
	rConn, err := net.DialTCP("tcp", nil, self.tcp)
	if err != nil {
		log.Fatalln(err)
	}

	c.conn = rConn

	if _, err := rConn.Write(buf[0:i]); err != nil {
		log.Printf("Client: %s err: %s", c.addr.String(), err)
		return
	}

	self.txBytes += uint64(i)

	go self.rw(c)
}
