package main

import (
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 128)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("read data failed...")
		return
	}
	fmt.Printf("read len: %d\n", ReadLen)
	fmt.Println(buff)

	msgBuf := buf[0:n]
	reciveMsg := &msg.Msg{}

	err = proto.Unmarshal(msgBuf, reciveMsg)
	if err != nil {
		fmt.Printf("unmarshaling error: ", reciveMsg)
	}
	fmt.Printf("msg id: %d\n", reciveMsg.GetMsgId())
	fmt.Printf("msg info: %s\n", reciveMsg.GetMsgInfo())
	fmt.Printf("msg from id: %s\n", reciveMsg.GetMsgFrom())
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":2121")
	if err != nil {
		fmt.Println("get tcp addr failed...")
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("listen tcp failed...")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}
