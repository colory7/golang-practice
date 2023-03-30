package main

/*
unsigned char* buf;
unsigned char arr[10] = {0};
unsigned char *getBuf() { return buf 或者 arr; }
unsigned char *setBuf(unsigned char** pbuf) {  buf = *pbuf; 或者复制内容到arr}
uint32_t size = 10;
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	gofp := (*C.char)(unsafe.Pointer(C.getBuf()))
	gostr := C.GoStringN(gofp, C.size)
	fmt.Println(gostr)

	s := []byte{97, 97, 97, 97, 98}
	cs := C.CString(string(s)) //这里会在c的runtime分配内存。
	C.setBuf(cs)
	C.free(unsafe.Pointer(cs)) //这里更安全的方式是去显示释放上面定义的buf，这里是个例子。
}
