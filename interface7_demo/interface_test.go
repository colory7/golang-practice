package interface7_demo

import (
	"fmt"
	"testing"
)

func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func TestInterfacePassParam(t *testing.T) {
	s := "zhang"
	show(s)

	b := true
	show(b)
}

func TestInterfaceAssign(t *testing.T) {
	var x interface{} // 定义一个空接口
	s := "zhang"
	x = s

	fmt.Printf("type:%T value:%v\n", x, x)

	i := 10
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)

	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

func TestPredicate(t *testing.T) {
	var x interface{}
	x = 888

	switch val := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", val)
	case int:
		fmt.Printf("x is a int，value is %v\n", val)
	case bool:
		fmt.Printf("x is a bool，value is %v\n", val)
	default:
		fmt.Println("unsupport type！")
	}
}
