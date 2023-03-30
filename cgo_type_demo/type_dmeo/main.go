package main

/*
   #include <stdint.h>
    int8_t a = -128;
    uint8_t b= 255;
    int16_t c= -32768;
    uint16_t d= 65535;
    int32_t e = -2147483648;
    uint32_t f = 4294967295;
    int64_t g = -9223372036854776001;
    uint64_t h = 9223372036854775999;
*/
import "C"
import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("---------------int8-------------")
	a1 := int8(C.a)
	fmt.Println(C.a)
	fmt.Println(a1)
	fmt.Println("类型为:", reflect.TypeOf(C.a))
	fmt.Println("转后的类型:", reflect.TypeOf(a1))
	fmt.Println("---------------uint8-------------")
	b1 := uint8(C.b)
	fmt.Println(C.b)
	fmt.Println(b1)
	fmt.Println("类型为:", reflect.TypeOf(C.b))
	fmt.Println("转后的类型:", reflect.TypeOf(b1))

	fmt.Println("---------------int16-------------")
	c1 := int16(C.c)
	fmt.Println(C.c)
	fmt.Println(c1)
	fmt.Println("类型为:", reflect.TypeOf(C.c))
	fmt.Println("转后的类型:", reflect.TypeOf(c1))
	fmt.Println("---------------uint16-------------")
	d1 := uint16(C.d)
	fmt.Println(C.d)
	fmt.Println(d1)
	fmt.Println("类型为:", reflect.TypeOf(C.d))
	fmt.Println("转后的类型:", reflect.TypeOf(d1))

	fmt.Println("---------------int32-------------")
	e1 := int32(C.e)
	fmt.Println(C.e)
	fmt.Println(e1)
	fmt.Println("类型为:", reflect.TypeOf(C.e))
	fmt.Println("转后的类型:", reflect.TypeOf(e1))
	fmt.Println("---------------uint32-------------")
	f1 := uint32(C.f)
	fmt.Println(C.f)
	fmt.Println(f1)
	fmt.Println("类型为:", reflect.TypeOf(C.f))
	fmt.Println("转后的类型:", reflect.TypeOf(f1))

	fmt.Println("---------------int64-------------")
	g1 := int64(C.g)
	fmt.Println(C.g)
	fmt.Println(g1)
	fmt.Println("类型为:", reflect.TypeOf(C.g))
	fmt.Println("转后的类型:", reflect.TypeOf(g1))
	fmt.Println("---------------uint64-------------")
	h1 := uint64(C.h)
	fmt.Println(C.h)
	fmt.Println(h1)
	fmt.Println("类型为:", reflect.TypeOf(C.h))
	fmt.Println("转后的类型:", reflect.TypeOf(h1))
}
