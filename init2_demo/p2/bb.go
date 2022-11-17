package p2

import "fmt"

var BBValue1 string = BBValue1Get()

func BBValue1Get() string {
	fmt.Println("test_util ttGet")
	return "5"
}

func init() {
	fmt.Println("bb.go init()")
}
