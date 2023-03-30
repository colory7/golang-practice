package main

// #cgo CFLAGS: -I${SRCDIR}/library
// #cgo LDFLAGS: -lstdc++ -L./library -llibrary
// #include "./library-bridge.h"
import "C"
import "unsafe"
import "fmt"

type Foo struct {
	ptr unsafe.Pointer
}

func NewFoo(value int) Foo {
	var foo Foo
	foo.ptr = C.LIB_NewFoo(C.int(value))
	return foo
}

func (foo Foo) Free() {
	C.LIB_DestroyFoo(foo.ptr)
}

func (foo Foo) value() int {
	return int(C.LIB_FooValue(foo.ptr))
}

func (foo Foo) testString(params string) string {
	return C.GoString(C.LIB_TestString(C.CString(params)))
}

func main() {
	foo := NewFoo(42)
	defer foo.Free() // The Go analog to C++'s RAII
	fmt.Println("[go]", foo.value())
	fmt.Println("----------------------")
	params := "gohello"
	result := foo.testString(params)
	fmt.Println("string result = ", result)
}
