package defer_demo

import (
	"fmt"
	"testing"
)

func deferfunc() int {
	a := 1
	b := 2

	defer fmt.Printf("因为a=%v，b=%v \n", a, b)
	fmt.Println("a+b=3")
	return func() int {
		fmt.Println("如果这一行比defer先输出，则说明使用defer语句的函数中，先返回return后执行defer语句，否则相反。")
		return 1
	}()
}

func TestDefer(t *testing.T) {
	c := deferfunc()
	fmt.Println(c)
}
