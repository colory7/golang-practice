package interface3_demo

import (
	"fmt"
	"testing"
)

func TestInterfacePassValue(t *testing.T) {
	dog := Dog{name: "花花"}
	fmt.Printf("dog: %p\n", &dog)

	fmt.Println("===================")

	dog.eat()
	fmt.Printf("dog: %v\n", dog)
}

func TestInterfacePassPointer(t *testing.T) {
	dog := &Dog{name: "花花"}
	fmt.Printf("dog: %p\n", dog)

	fmt.Println("===================")
	dog.eat2()
	fmt.Printf("dog: %v\n", dog)
}
