package number_demo

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestNumberCalc(t *testing.T) {
	i1 := math.MaxInt32
	fmt.Println(i1)
	fmt.Println(reflect.TypeOf(i1))

	i2 := i1 + 1
	fmt.Println(i2)
	fmt.Println(reflect.TypeOf(i2))

	i3 := i1 + i1
	fmt.Println(i3)
	fmt.Println(reflect.TypeOf(i3))

	i4 := math.MaxInt64 - 1
	fmt.Println(i4)
	fmt.Println(reflect.TypeOf(i4))

	i5 := math.MaxInt64 - math.MaxInt32
	fmt.Println(i5)
	fmt.Println(reflect.TypeOf(i5))

	//fmt.Println(math.MaxInt64 + 1)

}
