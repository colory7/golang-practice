package slice_demo

import (
	"fmt"
	"testing"
)

func TestSliceCreate5(t *testing.T) {
	var s1 []int
	fmt.Println(s1)

	s2 := new([]int)
	fmt.Println(s2)

	var s3 = []int{}
	fmt.Println(s3)

	var s4 = make([]int, 0)
	fmt.Println(s4)

	var s5 = make([]int, 3)
	fmt.Println(s5)
}

func TestSliceDelete(t *testing.T) {
	var s1 = []int{1, 2, 3, 4}               // 初始化一个切片
	var index = 2                            // 要删除的下标
	s1 = append(s1[:index], s1[index+1:]...) //删除下标为index的元素

	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}
}

func TestSliceDeepCopy(t *testing.T) {
	var s1 = []int{1, 2}    // 初始化一个切片
	var s2 = make([]int, 2) // 初始化一个空的切片，cap为2
	copy(s2, s1)            // 将s1拷贝给s2
	s2[0] = 99              // 改变s2[0]
	fmt.Println(s1[0])      // 打印s1[0]
}

func TestTwoDimonSlice(t *testing.T) {
	n := 10
	m := 20
	var s = make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, m)
	}

	fmt.Println(len(s), cap(s))
}
