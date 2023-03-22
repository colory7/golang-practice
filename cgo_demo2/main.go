package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./lib -lhi
#include "hi.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	d := make([]byte, 16)
	d = []byte("aabbccdd")
	str := C.CString(string(d))
	fmt.Println(C.demo(str))
	fmt.Println(string(d))
	C.free(unsafe.Pointer(str))
}
