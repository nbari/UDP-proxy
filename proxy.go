package UDPProxy

import (
	//"log"
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
	if rconn, err := net.DialUDP("udp", nil, r); err != nil {
		return nil, err
	} else {
		return &UDPProxy{
			lconn: l,
			rconn: rconn,
			debug: d,
		}, nil
	}
}
