package type_demo

import (
	"fmt"
	"reflect"
	"testing"
)

type FM_DCH string

const (
	DCH_MINUS FM_DCH = "-"
	DCH_SLASH        = "/"
)

func TestConstType(t *testing.T) {
	fmt.Println(reflect.TypeOf(DCH_MINUS).String())
	fmt.Println(reflect.TypeOf(DCH_SLASH).String())
}

func TestType(t *testing.T) {
	num := 222
	fmt.Println(reflect.TypeOf(num))

	name := "333"
	fmt.Println(reflect.TypeOf(name))

	f := func(data int) {
		fmt.Println(fmt.Println("hello", data))
	}

	fmt.Println(reflect.TypeOf(f))

	var x interface{}
	fmt.Println(reflect.TypeOf(x))

	x = 66
	fmt.Println(reflect.TypeOf(x))
}
