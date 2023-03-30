package main

/*
#include <stdlib.h>

typedef struct User {
	int id;
	int age;
	int number;
} User;

static void createUser(void **pUser) {
	if(pUser) *pUser = malloc(sizeof(User));
	struct User aa={1,1,1};
	*pUser=&aa;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type User C.User

func main() {
	pointer := unsafe.Pointer(nil)
	C.createUser(&pointer)
	user := (*User)(pointer)
	fmt.Println(user)
}
