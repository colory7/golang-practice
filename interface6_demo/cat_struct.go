package interface6_demo

import "fmt"

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (cat Cat) sleep() {
	fmt.Println("cat sleep...")
}
