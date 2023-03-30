package cgo_aux

/*
#include <stdio.h>
#include <stdlib.h>
void printString(const char* s) {
	printf("%s", s);
}
*/
import "C"
import "unsafe"

func PrintString(s string) {
	var cs *C.char = C.CString(s)
	C.printString(cs)
	C.free(unsafe.Pointer(cs))
}
