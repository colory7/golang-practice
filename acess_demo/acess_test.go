package acess_demo

import (
	"cobra_demo/acess_demo/mycom"
	"testing"
)

func Test(t *testing.T) {

	mycom.IsBlank("hello")

	//mycom.isEmpty("1");  // 会报错
}
