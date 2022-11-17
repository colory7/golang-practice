package network_demo

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"testing"
)

func TestTcpServer(t *testing.T) {
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

func TestTcpClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:9003")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		n, err := conn.Write(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		buf := make([]byte, 1024)
		n, err = conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Rev:", string(buf[:n]))
		if string(data) == "quit" {
			break
		}
	}
}

func TestUdpServer(t *testing.T) {
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

func TestUdpClient(t *testing.T) {
	conn, err := net.Dial("udp", "127.0.0.1:9003")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		n, err := conn.Write(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		buf := make([]byte, 1024)
		n, err = conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Rev:", string(buf[:n]))
		if string(data) == "quit" {
			break
		}
	}
}
