package UDPProxy

import (
	"log"
	"net"
)

type UDPProxy struct {
	rconn   *net.UDPConn
	caddr   *net.UDPAddr
	txBytes uint64
	rxBytes uint64
	debug   bool
}

func New(conn *.net.UDPConn, tcp *net.TCPAddr, udp *net.UDPAddr) *UDPProxy {

//	return proxy
//}

func (self *UDPProxy) Start(debug bool) {
	if debug {
		self.debug = true
	}

	var buffer = make([]byte, 0xffff)
	for {

	}
}
