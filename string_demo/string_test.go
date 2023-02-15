package string_demo

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	s := " aa cc  "
	s2 := strings.TrimSpace(s)
	fmt.Println(len(s))
	fmt.Println(len(s2))

}

func TestReverseBytes(t *testing.T) {
	fmt.Println(ReverseBytes("1010"))
}

func TestUtf8Index(t *testing.T) {
	s1 := "北京人efg蓝天安门白云abc最美丽天安门"

	fmt.Println(Utf8Index(s1, "天安门"))
	fmt.Println(s1[Utf8Index(s1, "天安门"):])

	fmt.Println(strings.Index(s1, "天安门"))
	fmt.Println(s1[strings.Index(s1, "天安门"):])

	fmt.Println(Utf8Index(s1, "aa"))
	fmt.Println(strings.Index(s1, "aa"))
}

func TestUtf8Index2(t *testing.T) {
	s1 := "北京人efg蓝天安门白云abc最美丽天安门"

	fmt.Println(Utf8Index2(s1, "天安门"))

	pos, asciiPos := Utf8Index2(s1, "天安门")
	fmt.Println(pos)
	fmt.Println(s1[asciiPos:])
	fmt.Println(s1[asciiPos+len("天安门"):])
}

func TestUtf8(t *testing.T) {
	str := "Hello,世界"
	fmt.Println("方法一  格式化打印")
	for _, ch1 := range str {
		fmt.Printf("%q", ch1) //单引号围绕的字符字面值，由go语法安全的转义
	}
	fmt.Println("方法二  转化输出格式")
	for _, ch2 := range str {
		fmt.Println(string(ch2))
	}

}
