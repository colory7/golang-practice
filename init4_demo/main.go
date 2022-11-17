package main

import (
	"fmt"
	"golang_practice/init4_demo/p1"
	"golang_practice/init4_demo/p1/p3"
	"golang_practice/init4_demo/p2"
)

func main() {
	fmt.Println("main.main()")
	fmt.Println(p3.P3Value)
	fmt.Println(p2.P2Value)
	fmt.Println(p1.P1Value)
}
