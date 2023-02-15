package bit_demo

import (
	"fmt"
	"golang_practice/string_demo"
	"math/bits"
	"strconv"
	"testing"
)

func TestBit(t *testing.T) {
	fmt.Println(strconv.FormatInt(2, 2))
	fmt.Println("reverse: " + string_demo.ReverseBytes(strconv.FormatInt(2, 2)))

	fmt.Println(strconv.FormatInt(10, 2))
	fmt.Println("reverse: " + string_demo.ReverseBytes(strconv.FormatInt(10, 2)))

	fmt.Println(strconv.FormatInt(2561, 2))
	fmt.Println("reverse: " + string_demo.ReverseBytes(strconv.FormatInt(2561, 2)))

	fmt.Println(strconv.FormatInt(3087, 2))
	fmt.Println("reverse: " + string_demo.ReverseBytes(strconv.FormatInt(3087, 2)))

	fmt.Println(strconv.FormatInt(99999999, 2))
	fmt.Println("reverse: " + string_demo.ReverseBytes(strconv.FormatInt(99999999, 2)))

}

func TestBitScan(t *testing.T) {
	i := 3
	i1 := bits.Reverse64(uint64(i))
	i2 := bits.ReverseBytes64(uint64(i))
	i3 := bits.LeadingZeros(uint(i))
	i4 := bits.TrailingZeros(uint(i))

	fmt.Println(strconv.FormatUint(uint64(i), 2))
	fmt.Println(strconv.FormatUint(i1, 2))
	fmt.Println(strconv.FormatUint(i2, 2))
	fmt.Println(strconv.FormatUint(uint64(i3), 2))
	fmt.Println(strconv.FormatUint(uint64(i4), 2))

	fmt.Println(strconv.FormatUint(uint64(bits.LeadingZeros64(bits.Reverse64(uint64(i)))), 2))
	fmt.Println(bits.Len64(uint64(i)))
	fmt.Println(bits.RotateLeft64(uint64(i), 0))
	fmt.Println(bits.RotateLeft64(uint64(i), 2))
	fmt.Println(bits.RotateLeft64(uint64(i), 65))
}

func TestBitScan2(t *testing.T) {
	num := 19
	i1 := bits.Reverse64(uint64(num))

	len := bits.Len64(uint64(num))

	fmt.Println(strconv.FormatUint(uint64(num), 2))
	fmt.Println(strconv.FormatUint(i1, 2))

	for i := len - 1; i >= 0; i-- {
		if (num & (1 << i)) == 0 {
			fmt.Println("第" + fmt.Sprint(i) + "位是0")
		} else {
			fmt.Println("第" + fmt.Sprint(i) + "位是1")
		}
	}
}
