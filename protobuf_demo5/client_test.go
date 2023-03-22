package protobuf_demo5

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"net"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	regMessage := &RegMessage{
		Id:       *proto.Int32(10001),
		Username: *proto.String("vicky"),
		Password: *proto.String("123456"),
		Email:    proto.String("eclipser@163.com"),
	}
	buffer, err := proto.Marshal(regMessage)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	pTCPAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:11111")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}
	pTCPConn, err := net.DialTCP("tcp", nil, pTCPAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}
	pTCPConn.Write(buffer)
}
