package bit_demo

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestBinary(t *testing.T) {
	fmt.Printf("%b\n", 3)
	fmt.Printf("%b\n", -3)
	fmt.Printf("%b\n", 3.0)

	fmt.Printf("%b\n", math.Float64bits(3))
	fmt.Printf("%b\n", math.Float64bits(3.0))
	fmt.Printf("%b\n", math.Float64frombits(3))
}

func TestBinary2(t *testing.T) {
	fmt.Printf("%b\n", math.Float64bits(3))
	fmt.Printf("%b\n", math.Float64bits(3.0))
	fmt.Printf("%b\n", math.Float32bits(3))
	fmt.Printf("%b\n", math.Float32bits(3.0))

	fmt.Printf("%b\n", math.Float64bits(-3))
	fmt.Printf("%b\n", math.Float64bits(-3.0))
	fmt.Printf("%b\n", math.Float32bits(-3))
	fmt.Printf("%b\n", math.Float32bits(-3.0))
}

func TestAB(t *testing.T) {
	ch := "-10"
	aa, e := strconv.ParseInt(ch, 2, 64)
	if e != nil {
		panic(e)
	}
	fmt.Println(aa)
}
