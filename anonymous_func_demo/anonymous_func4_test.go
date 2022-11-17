package anonymous_func_demo

import (
	"fmt"
	"testing"
)

func prints(list []int, f func(int)) {
	for _, b := range list {
		f(b) //打印
	}
}

func TestAnonymousFunc6(t *testing.T) {
	prints([]int{1, 2, 3}, func(i int) {
		fmt.Println(i)
	})
}
