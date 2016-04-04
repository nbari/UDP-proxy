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
	var (
		buffer = make([]byte, 1400)
		n      int
		err    error
	)

	// -> proxy send
	if _, err = self.rconn.Write(p.Data); err != nil {
		log.Println(err)
		return
	}

	if n, err = self.rconn.Read(buffer); err != nil {
		log.Println(err)
		return
	}
	log.Printf("N: %d, read Buffer: %s, %x", n, buffer[0:n], buffer[0:n])

	// <- proxy read
	if _, err = self.lconn.WriteToUDP(buffer[0:n], p.Addr); err != nil {
		log.Println(err)
		return
	}
}
