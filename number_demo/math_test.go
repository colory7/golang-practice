package number_demo

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestMax(t *testing.T) {
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(strconv.FormatFloat(math.MaxFloat32, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64))

	fmt.Println(strconv.FormatFloat(-math.MaxFloat32, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(-math.MaxFloat64, 'f', -1, 64))

	fmt.Println(1.4e-45)
	fmt.Println(4.9e-324)

	minFloat64 := strconv.FormatFloat(4.9e-324, 'f', -1, 64)
	fmt.Println(minFloat64)

	fmt.Println("====")
	minFloat64Str := "-1.7976931348623157999999"
	//minFloat64 := "1.79769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	minFloat64F, err := strconv.ParseFloat(minFloat64Str, 64)
	if err != nil {
		panic(err)
	}
	maxFloat64 := strconv.FormatFloat(minFloat64F, 'f', -1, 64)
	fmt.Println(maxFloat64)

}

func TestRound(t *testing.T) {
	fmt.Println(math.Round(2.4))
	fmt.Println(math.Round(2.5))
	fmt.Println(math.Round(2.9))

	fmt.Println(math.Round(2.45))

	fmt.Println(math.Trunc(math.Round(2.45)))

	fmt.Println(math.Round(889898.999666))

}

func TestNumSqrt(t *testing.T) {
	n := 10006
	fmt.Println(int(math.Sqrt(float64(n))))

}

func TestNumberConvert(t *testing.T) {
	// 十进制
	var a int64 = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)
	fmt.Printf("%o \n", a)
	fmt.Printf("%x \n", a)
	fmt.Println("-----------1")
	// 八进制  以0开头
	var b int64 = 077
	fmt.Printf("%d \n", b)
	fmt.Printf("%b \n", b)
	fmt.Printf("%o \n", b)
	fmt.Printf("%x \n", b)
	fmt.Println("-----------1")
	// 十六进制  以0x开头
	var c int64 = 0xff
	fmt.Printf("%d \n", c)
	fmt.Printf("%b \n", c)
	fmt.Printf("%o \n", c)
	fmt.Printf("%x \n", c)
}

func TestParseInt(t *testing.T) {
	s := "9223372036854775806"
	s1, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(s1)
}

func TestParseInt3(t *testing.T) {
	aa, err := strconv.ParseInt("2.3", 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(aa)

	//fmt.Println(int64(2.3))
}

func TestFormatInt(t *testing.T) {
	s := "9223372036854775806"
	s1, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	s2 := strconv.FormatInt(s1, 8)
	fmt.Println(s2)

	s3, err := strconv.ParseInt(s2, 8, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(s3)

	fmt.Println(math.MaxInt64)
}

func TestFormatInt3(t *testing.T) {
	//s := int64(float64(1.2))
	s := int64(math.Floor(1.2))
	s2 := strconv.FormatInt(s, 8)
	fmt.Println(s2)
}

func TestFormatInt2(t *testing.T) {
	s := "9223372036854775805.2"
	s1, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	s2 := strconv.FormatInt(s1, 8)
	fmt.Println(s2)

	s3, err := strconv.ParseInt(s2, 8, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(s3)

	fmt.Println(math.MaxInt64)
}

func TestOct2Decimal(t *testing.T) {
	r := strconv.FormatInt(9223372036854775806, 8)
	n, err := fmt.Printf("%v", r)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

}

func TestMaxFloat(t *testing.T) {
	f := math.MaxFloat32
	fmt.Println(math.MaxFloat32)
	minFloatStr := strconv.FormatFloat(f, 'f', -1, 32)
	fmt.Println(minFloatStr)

}
