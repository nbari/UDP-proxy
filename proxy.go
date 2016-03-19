package UDPProxy

import (
	"log"
	"net"
)

type UDPProxy struct {
	conn    *net.UDPConn
	tcp     *net.TCPAddr
	udp     *net.UDPAddr
	txBytes uint64
	rxBytes uint64
	debug   bool
}

type Backend struct {
	conn interface{}
}

func New(bind string, tcp *net.TCPAddr, udp *net.UDPAddr) *UDPProxy {
	addr, err := net.ResolveUDPAddr("udp", bind)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	proxy := &UDPProxy{}
	proxy.conn = conn
	proxy.tcp = tcp
	proxy.udp = udp

	return proxy
}

func (self *UDPProxy) Start(debug bool) {
	if debug {
		self.debug = true
	}

	var (
		backend = &Backend{}
		err     error
	)

	if self.udp != nil {
		backend.conn, err = net.DialUDP("udp", nil, self.udp)
	}
	if err != nil {
		log.Fatalln(err)
	}

	var buffer = make([]byte, 0xffff)
	for {
		n, clientAddr, err := self.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
		} else {
			go self.handlePacket(buffer[0:n], clientAddr, backend)
		}
	}
}
