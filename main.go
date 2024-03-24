package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("run")

	addr := net.UDPAddr{
		Port: 31193,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic("zhopa")
	}

	defer conn.Close()

	var buf [1024]byte
	for {
		rlen, remote, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("err %v", err)
			continue
		}
		fmt.Printf("new packet, len %d, ip %v\n", rlen, remote.IP)
		remote.Port = 56789

		resp := []byte("zhopa")
		conn.WriteTo(resp, remote)
	}
}
