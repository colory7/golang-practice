package convert_demo

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {
	var a int = 10
	var b *int = &a
	var c *int64 = (*int64)(unsafe.Pointer(b))
	fmt.Println(*c)
}

func TestInterface(tx *testing.T) {
	var a interface{} = 10
	switch a.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("string")
	}
}

func TestConvertInterface(tx *testing.T) {
	var a interface{} = 10
	t1, ok := a.(int)
	if ok {
		fmt.Println("int", t1)
	}
	t2, ok := a.(float32)
	if ok {
		fmt.Println("float32", t2)
	}

	t3 := a.(int)
	fmt.Println("t3: ", t3)

	t4 := a.(float32)
	fmt.Println("t4: ", t4)
}
