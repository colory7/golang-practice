//go:build gss
// +build gss

package main

import "C"
import "unsafe"

func main() {
	s := "abc"
	cs := C.CString(s)
	C.fputs(cs, (*C.FILE)(C.stdout))
	C.free(unsafe.Pointer(cs)) // 释放分配在C中的内存，否则会造成内存泄露
}
