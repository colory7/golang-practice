package generic_demo

import (
	"fmt"
	"testing"
)

type vestor[T any] []T

// [T any]表示支持任何类型的参数  （s []T表示形参s是一个T类型的切片）
func printslice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v\n", v)
	}
}

func TestFunc(txx *testing.T) {
	//常规调用
	printslice[int]([]int{66, 77, 88, 99, 100})
	printslice[string]([]string{"dudu", "yiyi", "8号"})
	//省略显示类型
	printslice([]int{88, 99, 100})
}
