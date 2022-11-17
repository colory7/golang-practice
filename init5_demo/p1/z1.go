package p1

import (
	"fmt"
	"golang_practice/init5_demo/p2"
)

var P1Value string = "P1Value"

func init() {
	fmt.Println("p1.z1 init()")
}

func P1Func1() {
	fmt.Println("P1Func1")
	p2.P2Func1()
}
