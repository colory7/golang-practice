package convert_demo

import (
	"fmt"
	"testing"
)

type Inter1 interface {
	getName() string
}

type Inter2 interface {
	printName()
}

//person 来实现这两个接口
type Person struct {
	name string
}

func (p Person) getName() string {
	return p.name
}
func (p Person) printName() {
	fmt.Println(p.name)
}

func check(t Inter1) {
	if f, ok1 := t.(Person); ok1 {
		fmt.Printf("f: %T，：%s\n", f, f.getName())
	}

	if t, ok2 := t.(Inter2); ok2 {
		check2(t)
	}
	if f1, ok2 := t.(Person); ok2 {
		fmt.Printf("f1: %T，%s\n", f1, f1.getName())
	}

	if f1, ok2 := t.(Inter1); ok2 {
		fmt.Printf("f1: %T，%s\n", f1, f1.getName())
	}
}

func check2(t Inter2) { //参数类型必须为inter2,否则会报错
	t.printName()
}

func TestInter(tx *testing.T) {
	var t Inter1
	t = Person{"xiaoming"}
	check(t)
}
