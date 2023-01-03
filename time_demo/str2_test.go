package time_demo

import "testing"

func TestStrPool(t *testing.T) {
	var str1 = "hello"
	var str2 = "hello"
	println(&str1, &str2)
}

func TestBytePool(t *testing.T) {
	var str1 = 'h'
	var str2 = 'h'
	println(&str1, &str2)
}

func TestIntPool(t *testing.T) {
	var str1 = 0
	var str2 = 0
	println(&str1, &str2)
}
