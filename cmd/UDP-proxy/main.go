package main

import (
	"flag"
	"fmt"
	p "github.com/nbari/UDP-proxy"
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
		raddr_tcp *net.TCPAddr
		raddr_udp *net.UDPAddr
		buffer    = make([]byte, 0xffff)
		err       error
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
		raddr_tcp, err = net.ResolveTCPAddr("tcp", *r)
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

	fmt.Printf("UDP-Proxy listening on %v port %d", addr.IP, addr.Port)

	// wait for connections
	var buffer = make([]byte, 0xffff)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
		}
		if *d {
			log.Printf("new connection from %s", clientAddr.String())
		}
		// make new connection to remote server
		proxy = p.New()

	}

}
