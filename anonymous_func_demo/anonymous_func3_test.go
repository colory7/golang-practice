package anonymous_func_demo

import (
	"fmt"
	"testing"
)

func TestAnonymousFunc4(t *testing.T) {
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	f := func(data int) {
		fmt.Println("hello", data)
	}
	f(200)
	f(300)
}
