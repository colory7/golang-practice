package slice_demo

import (
	"fmt"
	"testing"
)

func TestSliceCutAndErgodic(t *testing.T) {
	intArr := [...]int{1, 2, 3, 4, 5, 6, 7, 9}
	//方式1
	s := intArr[1:3]
	fmt.Println(len(s), cap(s))
	fmt.Println("====")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println("====")

	//方式2
	s = intArr[1:3:6]
	fmt.Println(len(s), cap(s))
	fmt.Println("====")
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Println("====")

}

func Test(t *testing.T) {

}
