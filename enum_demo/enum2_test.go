package enum_demo

import (
	"fmt"
	"testing"
)

const (
	a1 = iota
	b1 = iota
	c1 = iota
	d1 = 2
	e1
	f1 = "aa"
	g1
	h1 = iota
	i1
	j1
)

func TestEnum2(t *testing.T) {
	// iota  从0开始，一行加1 ，遇到自定义的值，从新计算
	// const 中的非自定义的值，顺延上一个的值
	fmt.Println(a1, b1, c1, d1, e1, f1, g1, h1, i1, j1)
}
