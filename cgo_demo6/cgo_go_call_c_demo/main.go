package main

// go调用c,会弹出一个框

/*
#include <stdio.h>

void say_hi(){
    printf("hello\n");
}

*/
import "C"

func main() {
	C.say_hi()
}
