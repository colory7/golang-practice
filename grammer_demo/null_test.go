package grammer_demo

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	bb := aa()
	fmt.Println(bb)
	fmt.Println(bb == nil)
}

func aa() interface{} {
	return dNull{}
}

type dNull struct{}
