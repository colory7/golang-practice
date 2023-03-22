package acess_demo

import (
	"golang_practice/acess_demo/mycom"
	"testing"
)

func Test(t *testing.T) {

	mycom.IsBlank("hello")

	//mycom.isEmpty("1");  // 会报错
}
