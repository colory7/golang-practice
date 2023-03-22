package main

/*

   //假设这段定义在外部C源码中对cgo不可见
   typedef struct User {
   	int id;
   	int age;
   	int number;
   } User;



   static void User_SetId(void *user, int id) {
   	((User *)user)->id = id;
   }

   static void User_SetAge(void *user, int age) {
   	((User *)user)->age = age;
   }

   static void User_SetNumber(void *user, int number) {
   	((User *)user)->number = number;
   }
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type User struct {
	Id     int32
	Age    int32
	Number int32
}

func main() {
	var user User

	pointer := unsafe.Pointer(&user)

	C.User_SetId(pointer, C.int(1))
	C.User_SetAge(pointer, C.int(25))
	C.User_SetNumber(pointer, C.int(10001))

	fmt.Println(&user)

}
