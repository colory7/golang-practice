package main

/*
#include <stdio.h>
//void TestArray(char* s[],int sLen);
static void DDD(char* s[],int sLen){
printf(s[sLen-1]);

//TestArray(s,sLen);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	//char site[7] = {'R', 'U', 'N', 'O', 'O', 'B', '\0'};

	var s = []*C.char{
		C.CString("88888888888"),
		C.CString("123456789"),
		C.CString("1234567890"),
		C.CString("12345678911"),
		C.CString("12345678912"),
	}
	fmt.Println(s)
	C.DDD((**C.char)(unsafe.Pointer(&s[0])), C.int(len(s)))

	//C.TestArray((**C.char)(unsafe.Pointer(&s[0])),C.int(len(s)))
	//C.DDD(s)
}

//export TestArray
func TestArray(t []*C.char, arrLen int) {
	fmt.Println("==>", C.GoString(t[arrLen-1]))
}
