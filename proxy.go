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

type Connection struct {
	ClientAddr *net.UDPAddr // Address of the client
	ServerConn *net.UDPConn // UDP connection to server
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

	var ClientDict map[string]*Connection = make(map[string]*Connection)
	var buffer = make([]byte, 0xffff)
	for {
		n, clientAddr, err := self.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			log.Printf("Received: %s From: %s", buffer[0:n], clientAddr.String())
		}
		saddr := clientAddr.String()
		conn, found := ClientDict[saddr]
		if !found {
			conn := new(Connection)
			conn.ClientAddr = clientAddr
			conn.ServerConn, err = net.DialUDP("udp", nil, self.udp)
			if err != nil {
				log.Fatal(err)
			}
			go self.handlePacket(clientAddr, conn.ServerConn)
		}
		print(conn)
		//		_, err = conn.ServerConn.Write(buffer[0:n])
		//		if err != nil {
		//		log.Fatal(err)
		//	}

	}
}
