package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("start TestHello()")

	var a int = 1
	var b int = 1

	var expectedResult int = 2

	ret := Sum(a, b)

	if ret != expectedResult {
		t.Error("expected 2, but x actually!!!")
	}

	fmt.Println("end TestHello()")
}
