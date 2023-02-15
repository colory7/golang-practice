package generic_demo

import (
	"fmt"
	"testing"
)

func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			fmt.Println(i)
			return 0
		}
	}
	return -1
}

func TestConstraint3(t *testing.T) {
	fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
	fmt.Println(findFunc([]string{"dudu", "yiyi", "8号"}, "dudu"))
}
