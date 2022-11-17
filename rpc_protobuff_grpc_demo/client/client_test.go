package client

import (
	"context"
	"fmt"
	message "golang_practice/rpc_protobuff_grpc_demo/pb"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	//1.链接服务器
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	//创建连接
	orderServiceClient := message.NewOrderServiceClient(conn)
	//构造请求参数
	orderRequest := &message.OrderRequest{
		OrderId:   "3",
		TimeStamp: time.Now().Unix(),
	}
	//发送请求并且得到回参
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	if err != nil {
		panic(err.Error())
	}
	if orderInfo != nil {
		fmt.Println("orderID:", orderInfo.OrderId)
		fmt.Println("orderName:", orderInfo.OrderName)
		fmt.Println("orderName:", orderInfo.OrderStatus)
	}
}
