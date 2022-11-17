package main

import (
	"fmt"
	"golang_practice/init2_demo/p1"
	"golang_practice/init2_demo/p2"
)

var AMainValue string = AMainValueGet()

func AMainValueGet() string {
	fmt.Println("AMainValueGet()")
	return "7"
}

func main() {
	fmt.Println("main() " + p1.AAValue1)
	fmt.Println("main() " + p2.BBValue1)
	fmt.Println("main() " + AMainValue)
}
