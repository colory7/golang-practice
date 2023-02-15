package generic2_demo

import (
	"fmt"
	"testing"
)

type super any

func print2[T super](s T) {
	fmt.Printf("%v\n", s)
}

func TestFunc(txx *testing.T) {
	print2[int](11)
	print2[string]("22")
	print2("33")
	print2(44)
}
