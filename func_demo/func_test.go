package func_demo

import (
	"fmt"
	"testing"
)

func show(a func()) {
	fmt.Printf("type:%T value:%p\n", a, a)
}

func TestFunc(t *testing.T) {
	show(func() {
		fmt.Println("show func test")
	})

}

func TestCallback(t *testing.T) {
	b := false

	var ff func(int) int
	//var ff interface{}
	if b {
		ff = aa
	} else {
		ff = bb
	}

	println(ff(333))
	println(ff(555))
}

func aa(a1 int) int {
	return 1
}
func bb(b1 int) int {
	return 2
}
