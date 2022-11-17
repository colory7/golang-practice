package simulate

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// 专门处理客户端连接
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [2048]byte
	for {
		n, err := reader.Read(buf[:])
		// 如果客户端关闭，则退出本协程
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("reader.Read error :", err)
			break
		}
		recvStr := string(buf[:n])
		// 打印收到的数据，稍后我们主要是看这里输出的数据是否是我们期望的
		fmt.Printf("received data：%s\n\n", recvStr)
	}
}

func startServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen error : ", err)
		return
	}
	defer listen.Close()
	fmt.Println("server start ...  ")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error :", err)
			continue
		}
		go process(conn)
	}
}
