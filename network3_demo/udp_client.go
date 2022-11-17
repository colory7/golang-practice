package network3_demo

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func startUdpClient() {
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
