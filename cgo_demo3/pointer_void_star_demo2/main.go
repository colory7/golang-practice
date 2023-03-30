package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var user User

	pointer := unsafe.Pointer(&user)
	setUser(pointer)

	fmt.Println(user)

}
