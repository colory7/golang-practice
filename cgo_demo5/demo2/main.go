package main

import "golang_practice/cgo_demo5/demo2/cgo_aux"

func main() {
	s := "hello"
	cgo_aux.PrintString(s)
}
