package main

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	fmt.Println(333)
	defer fmt.Println(111)
	defer fmt.Println(222)

	defer func() {
		if err := recover(); err != nil {
			//println(err.(string))
			fmt.Println(err)
			fmt.Println(777)
		}
	}()

	fmt.Println(666)

	defer fmt.Println("444")

	//panic("panic error!")
	panic("panic error!")

	defer fmt.Println(999)
	defer fmt.Println(888)

	fmt.Println(555)
}

func TestPanic2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("异常B")
	}()

	panic("异常C")
}

func TestPanic3(t *testing.T) {
	f1()
	r := recover()
	fmt.Printf("%s \n", r)
	fmt.Println("main func end")
}

func TestPanic4(t *testing.T) {
	defer func() {
		fmt.Println("defer func start")
		if r := recover(); r != nil {
			fmt.Printf("%s \n", r)
		}
		fmt.Println("defer func end")
	}()
	f1()
	fmt.Println("main func end")
}

func f1() {
	fmt.Println("func f1 start")
	arr := []int{}
	fmt.Println(arr[10])
	fmt.Println("func f1 end")
}
