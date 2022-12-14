package struct_demo

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

// https://zhuanlan.zhihu.com/p/403683508
func TestStructCreate(t *testing.T) {
	//方法一
	var person Person
	person.Name = "zhangsan"
	person.Age = 20
	fmt.Println(person)

	//方法二
	person1 := Person{"zhangsan1", 21}
	fmt.Println(person1)

	//方法三，此处(*person2).Name 等同于 person2.Name ，其他属性同理,因为go语言设计者在底层做了相关处理
	var person2 = new(Person)
	//(*person2).Name = "zhangsan2"
	person2.Name = "zhangsan2"
	(*person2).Age = 22
	fmt.Println(*person2)

	//方法四，此处(*person3).Name 等同于 person3.Name ，其他属性同理
	var person3 = &Person{}
	(*person3).Name = "zhangsan3"
	(*person3).Age = 23
	fmt.Println(*person3)
}
