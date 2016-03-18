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

type Client struct {
	addr *net.UDPAddr
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
	defer self.conn.Close()

	if debug {
		self.debug = true
	}

	var buffer = make([]byte, 0xffff)
	for {
		n, clientAddr, err := self.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error: ", err)
		} else {
			client := &Client{}
			client.addr = clientAddr
			if self.udp != nil {
				go self.handlePacketUDP(n, buffer, client)
			} else {
				go self.handlePacketTCP(n, buffer, client)
			}
		}
	}
}
