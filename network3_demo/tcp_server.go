package network3_demo

import (
	"fmt"
	"net"
)

func startTcpServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:9003")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// go语言闭包函数对外部变量是以引用的方式引用，所以这里用函数传递conn
		go func(conn net.Conn) {
			for {
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err)
					break
				}
				msg := string(buf[:n])
				fmt.Println(conn.RemoteAddr(), ":", msg)
				if msg == "quit" {
					conn.Write([]byte("Bye!"))
					conn.Close()
					break
				} else {
					conn.Write([]byte("OK"))
				}
			}
		}(conn)
	}
}
