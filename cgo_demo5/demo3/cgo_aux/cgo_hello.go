package cgo_aux

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* printInCFunc( char* s) {
	strcpy ( s , "abc " );

	printf("c print s: %s\n", s);
	char* ee="efg";
	printf("c print ee: %s\n", ee);

	return ee;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func PrintString(s string) {
	var cs *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	fmt.Println("gostring before:" + C.GoString((*C.char)(cs)))
	fmt.Println(*(*[]byte)(unsafe.Pointer(cs)))

	C.printInCFunc(cs)
	fmt.Println("gostring after:" + C.GoString(cs))
	fmt.Println("gostring after:" + C.GoString((*C.char)(cs)))

	//fmt.Println("go return : " + *(*string)(unsafe.Pointer(ret)))
	fmt.Println("go PrintString s: " + s)
	fmt.Println("go PrintString cs: " + *(*string)(unsafe.Pointer(cs)))

	fmt.Println(cs)
	ptr := uintptr(unsafe.Pointer(cs))
	fmt.Println(unsafe.Pointer(cs))
	fmt.Println(unsafe.Sizeof(ptr))
	fmt.Println(ptr)

	//C.free(unsafe.Pointer(ret))

}
