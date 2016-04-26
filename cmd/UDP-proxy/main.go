package main

import (
	"flag"
	"fmt"
	"github.com/nbari/UDP-proxy"
	"log"
	"net"
	"os"
)

var version, githash string

func main() {
	var (
		b         = flag.String("b", ":1514", "bind to host:port")
		r         = flag.String("r", "", "remote host:port")
		f         = flag.Bool("f", false, "forward only UDP -> TCP")
		v         = flag.Bool("v", false, fmt.Sprintf("Print version: %s", version))
		d         = flag.Bool("d", false, "Debug mode")
		buffer    = make([]byte, 1400)
		raddr_udp *net.UDPAddr
		//		raddr_tcp *net.TCPAddr
		err error
	)

	flag.Parse()

	if *v {
		if githash != "" {
			fmt.Printf("%s+%s\n", version, githash)
		} else {
			fmt.Printf("%s\n", version)
		}
		os.Exit(0)
	}

	if *r == "" {
		fmt.Println("-r remote host:port required")
		os.Exit(1)
	}

	// UDP or TCP
	if *f {
		//		raddr_tcp, err = net.ResolveTCPAddr("tcp", *r)
	} else {
		raddr_udp, err = net.ResolveUDPAddr("udp", *r)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// open local port to listen for incoming connections
	addr, err := net.ResolveUDPAddr("udp", *b)
	if err != nil {
		fmt.Printf("Failed to resolve local address: %s", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("Failed to open local port to listen: %s", err)
		os.Exit(1)
	}

	p, err := UDPProxy.New(conn, raddr_udp, *d)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Printf("UDP-Proxy listening on %s\n", addr.String())

	// wait for connections
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error: UDP read error: %v", err)
			continue
		}
		p.Counter++
		if *d {
			log.Printf("New connection from %s read bytes: %d connections: %d", clientAddr.String(), n, p.Counter)
		}
		go p.HandlePack(UDPProxy.Packet{clientAddr, buffer[0:n]})
	}
}
