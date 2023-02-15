package inherit1_demo

import (
	"fmt"
	"testing"
)

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Println("sleep")
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func TestInherit(t *testing.T) {
	dog := Dog{
		Animal{
			name: "dog",
			age:  2,
		},
	}

	cat := Cat{
		Animal{
			name: "cat",
			age:  3,
		},
	}

	dog.eat()
	dog.sleep()

	cat.eat()
	dog.sleep()
}
