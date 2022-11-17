package number_demo

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestA1(t *testing.T) {
	var ff float64
	ff = -1.355923156
	ff = FloatTruncate(ff, 4)
	fmt.Println(ff) // 输出 -1.3551
}

func TestA2(t *testing.T) {
	var ff float64
	ff = 231.355923156
	ff = FloatTruncate(ff, 4)
	fmt.Println(ff) // 输出 -1.3551
}

func TestA3(t *testing.T) {
	var ff float64
	ff = 231.355923156
	ff = FloatTruncate(ff, 0)
	fmt.Println(ff) // 输出 -1.3551
}

func TestA4(t *testing.T) {
	var ff float64
	ff = 231.355923156
	ff = FloatTruncate(ff, -1)
	fmt.Println(ff) // 输出 -1.3551
}

func TestA5(t *testing.T) {
	var ff float64
	ff = 231.355923156
	ff = FloatTruncate(ff, -2)
	fmt.Println(ff) // 输出 -1.3551
}

func TestA6(t *testing.T) {
	var ff float64
	ff = 231.355923156
	ff = FloatTruncate(ff, -12)
	fmt.Println(ff) // 输出 -1.3551
}

func TestA7(t *testing.T) {
	x := 231
	k := 2
	fmt.Println(x & (^(1 << (k - 1))))
}

func TestA8(t *testing.T) {
	var d uint8 = 2
	fmt.Printf("%08b\n", d)  // 00000010
	fmt.Printf("%08b\n", ^d) // 11111101
}

func Test9(t *testing.T) {
	x := 231
	fmt.Println(x | 1 - 1)

	x = 230
	fmt.Println(x | 1 - 1)
}

func Test10(t *testing.T) {
	x := 231
	k := 2
	fmt.Println(x & ^(1 << (k - 1)))
}

//func Test11(t *testing.T) {
//	x := 231.0
//	k := 2.0
//
//	base := math.Pow(10, k)
//	math.Mod(10)
//	format := "%." + strconv.Itoa(x%base*base) + "f"
//	fmt.Sprintf("%08b\n")
//
//}

func Test12(t *testing.T) {
	fmt.Println(1 % 2)
	fmt.Println(2 % 2)
	fmt.Println(3 % 2)
	fmt.Println(4 % 2)
}

//func Test13(t *testing.T) {
//	fmt.Println(1.2 % 2.2)
//	fmt.Println(2 % 2)
//	fmt.Println(3 % 2)
//	fmt.Println(4 % 2)
//}

func Test14(t *testing.T) {
	//math.Trunc(±0) = ±0
	//math.Trunc(±Inf) = ±Inf
	//math.Trunc(NaN) = NaN
}

func Test15(t *testing.T) {
	//x := 2
	//y := 10
	//
	////fmt.Printf("%.f", math.Trunc(x/y)*y)
	//fmt.Printf("%.f", x)
	//fmt.Println(3.2)
}

func Test16(t *testing.T) {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.2f\n", x, math.Exp(float64(x)))
	}
}

func Test17(t *testing.T) {
	fmt.Println(math.Mod(2.3, 10))
	fmt.Println(math.Mod(2.3, 100))
	fmt.Println(math.Modf(2.3))

	fmt.Println("==================")
	fmt.Println(math.Modf(231.325 / 10))
	fmt.Println(math.Modf(231.325 / 100))
	fmt.Println(math.Modf(231.325 / 1000))
	fmt.Println(math.Modf(231.325 / 10000))

	fmt.Println("==================")

	fmt.Println(math.Modf(-231.325 / 10))
	fmt.Println(math.Modf(-231.325 / 100))
	fmt.Println(math.Modf(-231.325 / 1000))
	fmt.Println(math.Modf(-231.325 / 10000))
}

func Test18(t *testing.T) {
	base := math.Pow10(2)
	fmt.Println(base)

	base2 := math.Pow10(3)
	fmt.Println(base2)

}

func Test19(t *testing.T) {
	knownNum := 231.325
	n := 2
	base10 := math.Pow10(n)
	remainder, _ := math.Modf(knownNum / base10)

	result := remainder * base10
	fmt.Println(result)
	fmt.Println(math.Modf(231.325 / math.Pow10(3)))
	fmt.Println(math.Modf(231.325 / math.Pow10(4)))
	fmt.Println(math.Modf(231.325 / math.Pow10(5)))

	fmt.Println("==================")

	fmt.Println(math.Modf(-231.325 / 10))
	fmt.Println(math.Modf(-231.325 / 100))
	fmt.Println(math.Modf(-231.325 / 1000))
	fmt.Println(math.Modf(-231.325 / 10000))
}

func Test20(t *testing.T) {
	num := 231.325
	fmt.Println(truncatePositive(num, 2))
	fmt.Println(truncatePositive(num, 3))
	fmt.Println(truncatePositive(num, 4))
	fmt.Println(truncatePositive(num, 5))

	fmt.Println("==================")
	num = -231.325

	fmt.Println(truncatePositive(num, 1))
	fmt.Println(truncatePositive(num, 2))
	fmt.Println(truncatePositive(num, 3))
	fmt.Println(truncatePositive(num, 4))
	fmt.Println(truncatePositive(num, 5))
}

func truncatePositive(num float64, n int) int64 {
	base10 := math.Pow10(n)
	remainder, _ := math.Modf(num / base10)

	return int64(remainder * base10)
}

// 截取小数位数
func FloatTruncate(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}

func TestDecimal(t *testing.T) {
	value := 2.3
	fmt.Println(math.Trunc(value*1e2+0.5) * 1e-2)
}

func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

func Decimal2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
