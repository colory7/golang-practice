package bit_demo

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"testing"
	"unsafe"
)

// https://tingzi.vip/417847.html

// AND 运算符具有选择性的把整型数据的位清除为 0 的好的效果
func TestBitAnd(t *testing.T) {
	var x uint8 = 0xAC // x = 10101100
	fmt.Printf("%b\n", x)

	x = x & 0xF0 // x = 10100000
	fmt.Printf("%b\n", x)
}

// 可以用 & 操作去测试一个数字是奇数还是偶数。原因是当一个数字的二进制的最低位是 1 的时候，那他就是奇数。
// 我们可以用一个数字和 1 进行 & 操作，然后在和 1 做 AND 运算，如果的到的结果是 1 ，那么这个原始的数字就是奇数
func TestAndOddEven(t *testing.T) {
	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}
}

// 我们可以利用按位或操作符为给定的整数有选择地设置单个位
func TestOr(t *testing.T) {
	// 第3，7，8位为1
	fmt.Printf("%b\n", 196)
	// fmt.Printf("%b\n", 3)

	var a uint8 = 0
	a |= 196
	//a |= 3
	fmt.Printf("%b", a)
}

func TestOr2(t *testing.T) {
	fmt.Printf("%b\n", 196)

	var a uint8 = 0
	a |= 196
	a |= 3
	fmt.Printf("%b", a)
}

// 异或运算的这个特性可以用来把二进制位的一个值变成另外一个值
func TestXOR(t *testing.T) {
	fmt.Printf("%b\n", 0xCEFF)

	var a uint16 = 0xCEFF
	a ^= 0xFF00 // same a = a ^ 0xFF00

	fmt.Printf("%b\n", a)
}

// 与 1 进行异或的位被翻转
func TestXOR2(t *testing.T) {
	a, b := -12, 25
	fmt.Println("a and b have same sign?", (a^b) >= 0)
}

// 用2个操作数为1的位将第第1个操作数的对应位设置为0
func TestAndNot(t *testing.T) {
	var a byte = 0xAB
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", 0x0F)
	a &^= 0x0F
	fmt.Printf("%08b\n", a)
}

// 取反
func TestNot(t *testing.T) {
	var a byte = 0x0F
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", ^a)

	a = 3
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", ^a)
}

func TestShift(t *testing.T) {
	var a int8 = 3
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", a<<1)
	fmt.Printf("%08b\n", a<<2)
	fmt.Printf("%08b\n", a<<3)
}

func TestShift2(t *testing.T) {
	a := 12
	fmt.Printf("%08b\n", a)
	fmt.Printf("%d\n", a<<2)
	fmt.Printf("%08b\n", a<<2)
}

// |配合位移:设置第n位的值
func TestShift3(t *testing.T) {
	var a int8 = 8
	fmt.Printf("%08b\n", a)
	a = a | (1 << 2)
	fmt.Printf("%08b\n", a)
}

// &配合位移: 测试是否设置了第n位
func TestShift4(t *testing.T) {
	var a int8 = 12
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", 1<<2)
	fmt.Printf("%08b\n", 4)
	if a&1 != 0 {
		fmt.Println("二进制位第1位是 1")
	} else {
		fmt.Println("二进制位第1位是 0")
	}
	if a&(1<<1) != 0 {
		fmt.Println("二进制位第2位是 1")
	} else {
		fmt.Println("二进制位第2位是 0")
	}
	if a&(1<<2) != 0 {
		fmt.Println("二进制位第3位是 1")
	} else {
		fmt.Println("二进制位第1位是 0")
	}
}

// &^配合位移: 将第n位设置为0
func TestShift5(t *testing.T) {
	var a int8 = 13
	fmt.Printf("%04b\n", a)
	a = a &^ (1 << 2)
	fmt.Printf("%04b\n", a)
}

func TestShift6(t *testing.T) {
	fmt.Println(runtime.GOARCH)  //CPU型号
	fmt.Println(strconv.IntSize) //int位数

	var a int = 12
	fmt.Println(unsafe.Sizeof(a))

	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", 1<<2)
	fmt.Printf("%08b\n", 4)
	if a&1 != 0 {
		fmt.Println("二进制位第1位是 1")
	} else {
		fmt.Println("二进制位第1位是 0")
	}
	if a&(1<<1) != 0 {
		fmt.Println("二进制位第2位是 1")
	} else {
		fmt.Println("二进制位第2位是 0")
	}
	if a&(1<<2) != 0 {
		fmt.Println("二进制位第3位是 1")
	} else {
		fmt.Println("二进制位第1位是 0")
	}

}
