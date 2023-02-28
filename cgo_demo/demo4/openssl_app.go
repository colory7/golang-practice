package main

/*
#include <openssl.rand.h>
*/
import "C"
import "fmt"

func main() {
	buffer := C.CString("123456")
	bufferSize := C.CInt(1)
	rc := C.RAND_bytes(buffer, bufferSize)
	fmt.Println(rc)
}
