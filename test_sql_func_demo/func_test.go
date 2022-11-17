package test_sql_func_demo

import (
	"fmt"
	"math"
	"testing"
	"unicode"
)

//ABS、EXP、MOD、POWER、SIGN、SQRT、CEIL、FLOOR、ROUND、TRUNCATE

func TestFunc(t *testing.T) {
	fmt.Println(math.Abs(22))
	fmt.Println(math.Abs(-22))
	fmt.Println(math.Abs(2.2))
	fmt.Println(math.Abs(-2.3))
	fmt.Println(math.Abs(0))

	fmt.Println("==============================")
	fmt.Println(math.Exp(0))
	fmt.Println(math.Exp(1))
	fmt.Println(math.Exp(2))
	fmt.Println(math.Exp(3))

	fmt.Println("==============================")

	fmt.Println(math.Mod(3, 2))
	fmt.Println(math.Mod(2, 3))
	fmt.Println(math.Mod(4, 8))
	fmt.Println(math.Mod(4, 0))
	fmt.Println(math.Mod(0, 4))

	fmt.Println("==============================")
	fmt.Println(math.Pow(2, 3))

	fmt.Println("Sign ==============================")

	fmt.Println(Sign(0))
	fmt.Println(Sign(0.3))
	fmt.Println(Sign(-0))
	fmt.Println(Sign(-0.2))
	fmt.Println(Sign(-32))
	fmt.Println(Sign(6))

	fmt.Println(math.Signbit(-32))
	fmt.Println(math.Signbit(0))
	fmt.Println(math.Signbit(-0))

	fmt.Println("==============================")

	fmt.Println(math.Sqrt(3.0))
	fmt.Println(math.Sqrt(2.0))
	fmt.Println(math.Sqrt(5.0))
	fmt.Println(math.Sqrt(4.0))
	fmt.Println(math.Sqrt(1.1))

	fmt.Println("==============================")
	fmt.Println(math.Ceil(2.3))

	fmt.Println("==============================")
	fmt.Println(math.Floor(3.4))

	fmt.Println("==============================")
	fmt.Println(math.Round(3.6))

	fmt.Println("==============================")
	fmt.Println(math.Trunc(3.6))

}

func Sign(x float64) int8 {
	isNegative := math.Signbit(x)
	switch isNegative {
	case true:
		return -1
	case false:
		if x == 0 {
			return 0
		} else {
			return 1
		}
	}

	panic("number: " + fmt.Sprintf("%f", x) + " is not correct")
}

func TestIsNumber(t *testing.T) {
	s := "Hello 123 4 5 6"
	for _, r := range s {
		fmt.Printf("%c = %v\n", r, unicode.IsNumber(r))
	}

	//fmt.Println(unicode.IsNumber(2))
	//fmt.Println(unicode.IsNumber(2.3))
	//fmt.Println(unicode.IsNumber(2.3))
	//fmt.Println(unicode.IsNumber(0.0))
	//fmt.Println(unicode.IsNumber(0))
	//fmt.Println(unicode.IsNumber(-0))
	//fmt.Println(unicode.IsNumber(-2.3))
}
