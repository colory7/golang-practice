package string_demo

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSprintf(t *testing.T) {
	fmt.Println(fmt.Sprintf("%x", "Hello"))

	//fmt.Println(fmt.Sprintf("%b", "48656c6c6f"))

	fmt.Println(strconv.Itoa(65))
	//fmt.Println(string(rune(x)))
}
