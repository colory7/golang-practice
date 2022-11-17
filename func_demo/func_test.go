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
