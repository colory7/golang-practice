package interface2_demo

import (
	"fmt"
	"testing"
)

type OpenClose interface {
	open()
	close()
}

type Door struct {
}

func (d Door) open() {
	fmt.Println("open door...")
}

//func (d Door) close() {
//	fmt.Println("close door...")
//}

func TestInterface(t *testing.T) {
	var oc OpenClose
	// 这里编译错误，提示只实现了一个接口
	//oc = Door{}
	oc.open()
}
