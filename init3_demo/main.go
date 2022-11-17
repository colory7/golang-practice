package main

import "fmt"

var _ int64 = s()

func init() {
	fmt.Println("init in main.go")
}
func s() int64 {
	fmt.Println("calling s() in main.go")
	return 1
}
func main() {
	fmt.Println("main")
}
