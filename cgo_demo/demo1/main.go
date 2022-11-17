package main

/*
#include "stdio.h"
void hello() {
 printf("Hello world !");
}
*/
import "C"

func main() {
	C.hello()
}
