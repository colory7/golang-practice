package type_demo

import (
	"fmt"
	"reflect"
	"testing"
)

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
