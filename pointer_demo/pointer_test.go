package pointer_demo

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {
	a := 10
	var b *int //int类型指针
	b = &a
	fmt.Println("a=", a)  // a=10
	fmt.Println("b=", b)  // b=0xc000225530 输出的是a的地址
	fmt.Println("c=", *b) // c=10  对a的地址进行取值
}

func TestPointer2Uintptr(t *testing.T) {
	a := int64(100)
	var ptr *int
	ptr = (*int)(unsafe.Pointer(&a))
	fmt.Printf("%d\n", *ptr) //输出100 如果是int32输出一个很大值

}
