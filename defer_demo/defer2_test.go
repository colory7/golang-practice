package defer_demo

import (
	"fmt"
	"testing"
)

func aa() {
	i := 0
	defer fmt.Printf("first defer %v \n", i) //输出0，因为i此时就是0
	i++
	defer fmt.Printf("second defer %v \n", i) //输出1，因为i此时就是1
	return
}

func TestDefer2(t *testing.T) {
	aa()
}
