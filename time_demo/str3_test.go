package time_demo

import "testing"

func TestStrPool2(t *testing.T) {
	var str1 = "hello"
	var str2 = "hello"
	println(&str1, &str2)
}

func TestBytePool2(t *testing.T) {
	var str1 = 'h'
	var str2 = 'h'
	println(&str1, &str2)
}

func TestIntPool2(t *testing.T) {
	var str1 = 0
	var str2 = 0
	println(&str1, &str2)
}
