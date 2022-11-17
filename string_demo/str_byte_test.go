package string_demo

import (
	"fmt"
	"testing"
	"unsafe"
)

func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func StringToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

func TestStringByte(t *testing.T) {
	str := "hello world!"

	fmt.Println(str)

	a := StringToBytes(str)

	fmt.Println(a)

	b := BytesToString(a)

	fmt.Println(b)
}
