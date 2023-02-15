package inherit2_demo

import (
	"fmt"
	"testing"
)

type A struct {
	age int
}
type C struct {
	age int
}
type B struct {
	C
	A
	name string
}

func Test(t *testing.T) {
	var b B
	b.A.age = 20
	fmt.Println(b.A.age)
	b.C.age = 40
	fmt.Println(b.C.age)
	fmt.Println(b)
}
