package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lqrng
#include <stdlib.h>
#include <stdio.h>
#include "qrng.h"
*/
import "C"
import "unsafe"

func main() {
	buffer := C.CString("abc")
	C.free(unsafe.Pointer(buffer))

	//var buffer = [1024]*C.ulonglong{}
	//ret := C.qrng_read_random_sched((**C.char)(buffer), C.int(1024))
	//fmt.Println(ret)

}
