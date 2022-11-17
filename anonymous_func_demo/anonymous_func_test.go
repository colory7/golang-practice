package anonymous_func_demo

import (
	"fmt"
	"testing"
)

func TestAnonymousFunc(t *testing.T) {
	result := func(a int, b int) int {
		c := a + b
		return c
	}(1, 2) //括号表示调用，并传入参数1和2

	fmt.Println(result)
}
