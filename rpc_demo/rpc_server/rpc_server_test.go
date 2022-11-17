package rpc_server

import (
	"errors"
	"fmt"
	. "golang_practice/rpc_demo/rpc_common"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"testing"
)

// 算数运算结构体
type Arith struct {
}

// 乘法运算方法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算方法
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func TestRpcServer(t *testing.T) {
	rpc.Register(new(Arith)) // 注册rpc服务
	rpc.HandleHTTP()         // 采用http协议作为rpc载体

	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection")

	http.Serve(lis, nil)
}
