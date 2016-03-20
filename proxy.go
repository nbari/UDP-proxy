package UDPProxy

import (
	"log"
	"net"
)

type UDPProxy struct {
	lconn   *net.UDPConn
	rconn   *net.UDPConn
	caddr   *net.UDPAddr
	txBytes uint64
	rxBytes uint64
	debug   bool
}

func New(lconn *net.UDPConn, c, r *net.UDPAddr, d bool) *UDPProxy {
	rconn, err := net.DialUDP("udp", nil, r)
	if err != nil {
		log.Println(err)
	}
	return &UDPProxy{
		lconn: lconn,
		rconn: rconn,
		caddr: c,
		debug: d,
	}
}

func (self *UDPProxy) Start(data []byte) {
	_, err := self.rconn.Write(data)
	if err != nil {
		log.Println(err)
	}
	go self.handlePacket()
}
