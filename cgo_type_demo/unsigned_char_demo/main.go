package main

/*
   #include<stdio.h>
    unsigned char a = 254;
*/
import "C"
import (
	"fmt"
	"reflect"
)

func main() {
	b := byte(C.a)
	fmt.Println(C.a)
	fmt.Println("类型为:", reflect.TypeOf(C.a))
	fmt.Println(b)
	fmt.Println("类型为:", reflect.TypeOf(b))
}
