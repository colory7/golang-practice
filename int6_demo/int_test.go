package int6_demo

// https://blog.csdn.net/u011582922/article/details/121732193
import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

func TestIntSize(t *testing.T) {
	fmt.Println("不同int类型占用的字节数大小：")
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Printf("int    : %v\n", unsafe.Sizeof(i1))
	fmt.Printf("int8   : %v\n", unsafe.Sizeof(i2))
	fmt.Printf("int16  : %v\n", unsafe.Sizeof(i3))
	fmt.Printf("int32  : %v\n", unsafe.Sizeof(i4))
	fmt.Printf("int64  : %v\n", unsafe.Sizeof(i5))
}

func TestUIntSize(t *testing.T) {
	fmt.Println("不同无符号int类型占用的字节数大小：")
	var i1 uint = 1
	var i2 uint8 = 2
	var i3 uint16 = 3
	var i4 uint32 = 4
	var i5 uint64 = 5
	fmt.Printf("uint    : %v\n", unsafe.Sizeof(i1))
	fmt.Printf("uint8   : %v\n", unsafe.Sizeof(i2))
	fmt.Printf("uint16  : %v\n", unsafe.Sizeof(i3))
	fmt.Printf("uint32  : %v\n", unsafe.Sizeof(i4))
	fmt.Printf("uint64  : %v\n", unsafe.Sizeof(i5))
}

func TestIntRange(t *testing.T) {
	// 不同int类型的取值范围
	fmt.Println("不同int类型的取值范围：")
	//fmt.Println("int:", math.MinInt, "~", math.MaxInt) 报错，没有 math.MinInt math.MaxInt
	fmt.Println("int8:", math.MinInt8, "~", math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, "~", math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, "~", math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, "~", math.MaxInt64)
	fmt.Println()
}

func TestUintRange(t *testing.T) {
	fmt.Println("uint8:", 0, "~", math.MaxUint8)
	fmt.Println("uint16:", 0, "~", math.MaxUint16)
	fmt.Println("uint32:", 0, "~", math.MaxUint32)
	//fmt.Println("uint64:", 0, "~", math.MaxUint64)
	fmt.Println()
}
