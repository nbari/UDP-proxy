package UDPProxy

import (
	"log"
	"net"
)

type UDPProxy struct {
	local *net.UDPConn
	tcp   *net.TCPAddr
	udp   *net.UDPAddr
	debug bool
}

func New(bind string, tcp *net.TCPAddr, udp *net.UDPAddr) *UDPProxy {
	addr, err := net.ResolveUDPAddr("udp", bind)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	proxy := &UDPProxy{}
	proxy.local = conn
	proxy.tcp = tcp
	proxy.udp = udp

	return proxy
}

func (self *UDPProxy) Start(debug bool) {
	defer self.local.Close()

	if debug {
		self.debug = true
	}

	buf := make([]byte, 1024)
	for {
		n, _, err := self.local.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error: ", err)
		} else {
			if self.udp != nil {
				go self.handlePacketUDP(n, buf)
			} else {
				go self.handlePacketTCP(n, buf)
			}
		}
	}
}
