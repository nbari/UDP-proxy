package main

import (
	"flag"
	"fmt"
	"github.com/nbari/UDP-proxy"
	"net"
	"os"
)

var version, githash string

func main() {
	var b = flag.String("b", ":1514", "bind to host:port")
	var d = flag.Bool("d", false, "Debug mode.")
	var f = flag.String("f", "", "forward to host:port")
	var u = flag.Bool("u", false, "forward using UDP instead of TCP")
	var v = flag.Bool("v", false, fmt.Sprintf("Print version: %s", version))

	flag.Parse()

	if *v {
		if githash != "" {
			fmt.Printf("%s+%s\n", version, githash)
		} else {
			fmt.Printf("%s\n", version)
		}
		os.Exit(0)
	}

	var proxy *UDPProxy.UDPProxy

	// UDP or TCP
	if *u {
		addr, err := net.ResolveUDPAddr("udp", *f)
		if err != nil {
			panic(err)
		}
		proxy = UDPProxy.New(*b, nil, addr)
	} else {
		addr, err := net.ResolveTCPAddr("tcp", *f)
		if err != nil {
			panic(err)
		}
		proxy = UDPProxy.New(*b, addr, nil)
	}

	// start
	proxy.Start(*d)
}
