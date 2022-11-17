package server

import (
	"context"
	"errors"
	"fmt"
	"golang_practice/rpc_protobuff_grpc_demo/pb"
	"google.golang.org/grpc"
	"net"
	"net/rpc"
	"testing"
	"time"
)

type OrderServiceImpl struct {
}

func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
	//创建订单信息
	orderMap := map[string]message.OrderInfo{
		"1": {OrderId: "1", OrderName: "衣服", OrderStatus: "已付款"},
		"2": {OrderId: "2", OrderName: "手机", OrderStatus: "未付款"},
		"3": {OrderId: "3", OrderName: "电脑", OrderStatus: "已付款"},
	}

	//得到当前时间
	current := time.Now().Unix()
	if request.TimeStamp > current {
		//返回异常
		return &message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}, errors.New("timeout")
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			//找到对应的d订单id，返回正常数据
			fmt.Println("result:", result)
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}

	return nil, nil
}

func TestServer(t *testing.T) {
	//1.初始化grpc
	server := grpc.NewServer()

	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	//3.下面的函数可以把 machUtil里面包含的功能函数注册到HTTP协议中,调用者可以使用http方式进行数据传输
	rpc.HandleHTTP()
	//4.监听

	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}

	//5.处理请求
	_ = server.Serve(listen)
}
