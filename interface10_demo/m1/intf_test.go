package m1

import (
	"fmt"
	"testing"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// struct 中实现接口中的同名函数
func Test(t *testing.T) {
	var s1 = Square{5}
	fmt.Printf("The square has area: %f\n", s1.Area())
}

func Test2(t *testing.T) {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
