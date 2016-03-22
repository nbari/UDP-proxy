package UDPProxy

import (
	"log"
	"net"
)

type Packet struct {
	Addr *net.UDPAddr
	Data []byte
}

func (self *UDPProxy) HandlePack(p Packet) {
	// -> proxy send
	var (
		buffer = make([]byte, 1500)
		n      int
		err    error
	)

	_, err = self.rconn.Write(p.Data)
	if err != nil {
		panic(err)
	}

	n, err = self.rconn.Read(buffer)
	log.Println("read size:", n)

	// <- proxy read
	_, err = self.lconn.WriteToUDP(buffer[0:n], p.Addr)
	if err != nil {
		log.Println(err)
	}
}
