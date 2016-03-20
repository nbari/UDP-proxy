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
		raddr_udp *net.UDPAddr
		buffer    = make([]byte, 0xffff)
		err       error
		clients   map[string]*UDPProxy.UDPProxy = make(map[string]*UDPProxy.UDPProxy)
		proxy     *UDPProxy.UDPProxy
		found     bool
	)
	//raddr_tcp *net.TCPAddr

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
		// raddr_tcp, err = net.ResolveTCPAddr("tcp", *r)
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

	log.Printf("UDP-Proxy listening on %s\n", addr.String())

	// wait for connections
	var counter uint64
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
		}
		counter++
		if *d {
			log.Printf("new connection from %s", clientAddr.String())
		}
		fmt.Printf("Connections: %d, clients: %d\n", counter, len(clients))
		proxy, found = clients[clientAddr.String()]
		if !found {
			// make new connection to remote server
			proxy = UDPProxy.New(conn, clientAddr, raddr_udp, *d)
			clients[clientAddr.String()] = proxy
		}
		go proxy.Start(buffer[0:n])
	}

}
