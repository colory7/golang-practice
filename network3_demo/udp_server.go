package network3_demo

import (
	"fmt"
	"net"
)

func startUdpServer() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9003")

	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		// UDP面向无连接的通信，UDP Server接收报文。
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(remoteAddr, ": ", string(buf[:n]))
		conn.WriteToUDP([]byte("OK"), remoteAddr)
	}
}
