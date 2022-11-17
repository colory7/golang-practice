package rpc_protobuf_client

import (
	"fmt"
	message "golang_practice/rpc_protobuf_demo/pb"
	"net/rpc"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	//1.链接服务器
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	//2.封装请求
	timeStamp := time.Now().Unix()
	requset := message.OrderRequest{OrderId: "2", TimeStamp: timeStamp}

	var response *message.OrderInfo //声明接受参数

	//3.同步调用
	err = client.Call("routeName.GetOrderInfo", requset, &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("resp.OrderId:", response.OrderId)
	fmt.Println("resp.OrderName:", response.OrderName)
	fmt.Println("resp.OrderStatus:", response.OrderStatus)
}
