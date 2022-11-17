package anonymous_func_demo

import (
	"fmt"
	"testing"
)

func TestAnonymousFunc2(t *testing.T) {
	result := func(a int, b int) int {
		c := a + b
		return c
	}
	r := result(2, 3)

	fmt.Println(r)

}
