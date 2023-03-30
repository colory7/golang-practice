package pointer_demo

import (
	"fmt"
	"testing"
	"unsafe"
)

func Int64To8(f int64) int8 {
	return *(*int8)(unsafe.Pointer(&f))
}
func Int8To64(f int8) int64 {
	return *(*int64)(unsafe.Pointer(&f))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func Test(t *testing.T) {
	// 错误
	fmt.Println("int8 => int64", Int8To64(5))
	fmt.Println("int64 => int8", Int64To8(5))

}
