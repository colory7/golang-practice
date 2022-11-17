package p1

import (
	"fmt"
	"golang_practice/init2_demo/p2"
)

var AAValue1 string = AAValue1Get()

//var AAValue1 string = p2.BBValue1Get()

var AAValue2 string = p2.BBValue1

func AAValue1Get() string {
	fmt.Println("AAValue1Get() ")
	return "5"
}

func init() {
	//a := test_util.BBValue1
	//fmt.Println("pack init pack ", a)

	fmt.Println("aa.go init() pack ")
	fmt.Println("aa.go init() pack " + AAValue2)
}
