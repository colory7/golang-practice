package p1

import (
	"fmt"
)

var CCValue1 string = CCValue1Get()

func CCValue1Get() string {
	fmt.Println("CCValue1Get")
	return "6"
}

func init() {
	fmt.Println("cc.go init() " + CCValue1)
}
