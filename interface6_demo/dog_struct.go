package interface6_demo

import "fmt"

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}
