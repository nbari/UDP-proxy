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
	var r = flag.String("r", "", "remote host:port")
	var f = flag.Bool("f", false, "forward only UDP -> TCP")
	var v = flag.Bool("v", false, fmt.Sprintf("Print version: %s", version))
	var d = flag.Bool("d", false, "Debug mode")

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

	if *r == "" {
		fmt.Println("-r remote host:port required")
		os.Exit(1)
	}

	// UDP or TCP
	if *f {
		addr, err := net.ResolveTCPAddr("tcp", *r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		proxy = UDPProxy.New(*b, addr, nil)
	} else {
		addr, err := net.ResolveUDPAddr("udp", *r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		proxy = UDPProxy.New(*b, nil, addr)
	}

	// start
	proxy.Start(*d)
}
