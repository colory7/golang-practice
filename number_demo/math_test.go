package number_demo

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestMax(t *testing.T) {
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(strconv.FormatFloat(math.MaxFloat32, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64))
}

func TestRound(t *testing.T) {
	fmt.Println(math.Round(2.4))
	fmt.Println(math.Round(2.5))
	fmt.Println(math.Round(2.9))

	fmt.Println(math.Round(2.45))

	fmt.Println(math.Trunc(math.Round(2.45)))

	fmt.Println(math.Round(889898.999666))

}

func TestNumSqrt(t *testing.T) {
	n := 10006
	fmt.Println(int(math.Sqrt(float64(n))))

}
