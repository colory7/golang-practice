package slice_demo

import (
	"fmt"
	"testing"
)

func TestSliceCreate2(t *testing.T) {
	t1 := []int{2, 3}
	fmt.Println(t1[0])

	t2 := make([]int, 5, 8)
	fmt.Println(len(t2), cap(t2))

}

func TestSliceCreate3(t *testing.T) {
	test := make([]int, 0)            // 创建一个长度为0，容量为0 的数组
	fmt.Println(len(test), cap(test)) // 0 0

	test = append(test, 1)
	fmt.Println(len(test), cap(test)) // 1 1

	test = append(test, 1)
	fmt.Println(len(test), cap(test)) // 2 2

	test = append(test, 1)
	fmt.Println(len(test), cap(test)) // 3 4

	test = append(test, 1)
	fmt.Println(len(test), cap(test)) // 4 4

	test = append(test, 1)
	fmt.Println(len(test), cap(test)) // 5 8
}
