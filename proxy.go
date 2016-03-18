package UDPProxy

import (
	"log"
	"net"
)

type UDPProxy struct {
	addr    *net.UDPAddr
	conn    *net.UDPConn
	tcp     *net.TCPAddr
	udp     *net.UDPAddr
	txBytes uint64
	rxBytes uint64
	debug   bool
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
	proxy.addr = addr
	proxy.conn = conn
	proxy.tcp = tcp
	proxy.udp = udp

	return proxy
}

func (self *UDPProxy) Start(debug bool) {
	defer self.conn.Close()

	if debug {
		self.debug = true
	}

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := self.conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error: ", err)
		} else {
			if self.udp != nil {
				go self.handlePacketUDP(n, buf, clientAddr)
			} else {
				go self.handlePacketTCP(n, buf)
			}
		}
	}
}
