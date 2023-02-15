package inherit4_demo

import (
	"fmt"
	"testing"
	"unsafe"
)

type Object struct {
}

func (o Object) this() Object {
	return o
}

var root = new(Object)

type A struct {
	age  int
	root *Object
}

func (a A) this() A {
	return a
}

type B struct {
	name string
	root *Object
}

type C struct {
	height float32
	root   *Object
}

func TestInherit(t *testing.T) {
	aa := new(A)
	aa.age = 23
	fmt.Println(aa.age)
	fmt.Println(aa.this().age)

	var cc *Object = (*Object)(unsafe.Pointer(aa))
	fmt.Println((*cc).this())
}

func TestUnsafe(t *testing.T) {
	var a int = 10
	var b *int = &a
	var c *int64 = (*int64)(unsafe.Pointer(b))
	fmt.Println(*c)
}
