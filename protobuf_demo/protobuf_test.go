package protobuf_demo

import (
	"fmt"
	test "golang_practice/protobuf_demo/aaa"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestMarshal(t *testing.T) {
	msgDataEncoding := OrderMarshal()
	fmt.Println(msgDataEncoding)
}

func TestProtobuf(t *testing.T) {
	msgDataEncoding := OrderMarshal()
	fmt.Println(msgDataEncoding)

	OrderUnmarshal(msgDataEncoding)
}

func OrderMarshal() []byte {
	//创建消息对象
	msg_test := &test.Order{
		OrderId: *proto.Int32(12),
		Num:     *proto.Int64(22),
		Name:    *proto.String("test str"),
	}

	//序列化
	msgDataEncoding, err := proto.Marshal(msg_test)
	if err != nil {
		fmt.Println("err:", err.Error())
		panic(err)
		return nil
	}

	return msgDataEncoding
}

func OrderUnmarshal(msgDataEncoding []byte) {
	//反序列化
	msgEmpty := test.Order{}
	empty := test.Order{}
	proto.Unmarshal(msgDataEncoding, &empty)

	fmt.Println("order_id:", msgEmpty.GetOrderId())
	fmt.Println("num:", msgEmpty.GetNum())
	fmt.Println("name:", msgEmpty.GetName())
}
