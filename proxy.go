package UDPProxy

import (
	"log"
	"net"
)

type UDPProxy struct {
	lconn   *net.UDPConn
	rconn   *net.UDPConn
	txBytes uint64
	rxBytes uint64
	debug   bool
	Counter uint64
}

func New(l *net.UDPConn, r *net.UDPAddr, d bool) (*UDPProxy, error) {
	rconn, err := net.DialUDP("udp", nil, r)
	if err != nil {
		return nil, err
	}
	return &UDPProxy{
		lconn: l,
		rconn: rconn,
		debug: d,
	}, nil
}

func (self *UDPProxy) Start(data []byte) {
	n, err := self.rconn.Write(data)
	if err != nil {
		log.Println(err)
	}
	self.txBytes += uint64(n)
	log.Printf("Sent Bytes: %d", self.txBytes)
	//	go self.handlePacket()
}
