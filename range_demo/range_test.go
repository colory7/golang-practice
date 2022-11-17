package range_demo

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	var a = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i: %d, v: %v\n", i, v)
	}

	fmt.Println("==============================")
	for i := range a {
		fmt.Printf("i: %d \n", i)
	}
}
