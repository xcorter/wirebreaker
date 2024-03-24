package main

import (
	"fmt"
	"net"
)

func main() {

	raddr := net.UDPAddr{
		Port: 8088,
		IP:   net.ParseIP("127.0.0.1"),
	}

	laddr := net.UDPAddr{
		Port: 8086,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, _ := net.DialUDP("udp", &laddr, &raddr)
	conn.Write([]byte("asd"))

	resp := make([]byte, 1024)
	conn.Read(resp)

	fmt.Println(resp)
}
