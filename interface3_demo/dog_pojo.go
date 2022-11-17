package interface3_demo

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) eat() {
	fmt.Println("start eat>>>>>>>>")
	fmt.Printf("dog: %p\n", &dog)
	fmt.Println("dog eat.. ")
	dog.name = "嘿嘿"

	fmt.Println(dog.name)
	fmt.Println("end eat<<<<<<<<<")
}

func (dog *Dog) eat2() {
	fmt.Println("start eat>>>>>>>>")
	fmt.Printf("dog: %p\n", dog)
	fmt.Println("dog eat..")
	dog.name = "小白"
	fmt.Println(dog.name)
	fmt.Println("end eat<<<<<<<<<")
}
