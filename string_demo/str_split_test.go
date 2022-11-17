package string_demo

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit3(t *testing.T) {
	str := "-100 123 200"

	//指定分隔符
	countSplit := strings.Split(str, " ")
	fmt.Println(countSplit, len(countSplit))

	//指定分割符号，指定分割次数
	countSplit = strings.SplitN(str, " ", 2)
	fmt.Println(countSplit, len(countSplit))

	fmt.Println(countSplit[0])
	fmt.Println(countSplit[1])
	//fmt.Println(countSplit[2])

	fmt.Println("=====")
	for k, v := range countSplit {
		t.Log(k, v)
	}

	fmt.Println("=====")
	for k, v := range countSplit {
		fmt.Println(k, v)
	}

}
