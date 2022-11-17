package main

//#include "stdio.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.CString("cgo Hello World")))
}
