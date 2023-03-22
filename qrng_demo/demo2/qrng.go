package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lqrng
#include <stdlib.h>
#include <stdio.h>
#include "qrng.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var buffer = [1024]uint64{}
	ptr := unsafe.Pointer(&buffer[0])
	ret := C.qrng_read_random_sched(ptr, C.uint(5))

	//for i := 0; i < len(buffer); i++ {
	//	qq := *(*C.ulonglong)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(C.ulonglong(1))))
	//	fmt.Println(qq)
	//}
	for i := 0; i < len(buffer); i++ {
		fmt.Println(buffer[i])
	}
	fmt.Println(ret)

	//defer C.free(ptr)
}
