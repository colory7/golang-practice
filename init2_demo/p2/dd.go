package p2

import (
	"fmt"
)

var DDValue1 string = DDValue1Get()

func DDValue1Get() string {
	fmt.Println("DDValue1Get()")
	return "ddv1"
}
func init() {
	fmt.Println("dd.go init()")
}
