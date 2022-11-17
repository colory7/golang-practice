package enum_demo

import (
	"fmt"
	"testing"
)

const (
	a = 1
	b
	c
	d = 2
	e
	f = "aa"
	g
	h = iota
	i
	j
)

func TestEnum(t *testing.T) {
	// iota 一行加1 ，除了字符串
	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}
