package number_demo

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	r1 := romanToInt("MI")
	fmt.Println(r1)

	r2 := romanToInt("MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM")
	fmt.Println(r2)
}

func TestRomanToInt2(t *testing.T) {
	r1 := romanToInt("MI")
	fmt.Println(r1)

	r2 := intToRoman2(1001)
	fmt.Println(r2.String())
}

func TestIntToRoman(t *testing.T) {
	r1 := intToRoman(32)
	fmt.Println(r1)
}

func romanToInt(s string) int {
	var sum int
	preNum := getValue(s[0])
	for i := 1; i < len(s); i++ {
		num := getValue(s[i])
		if preNum < num {
			sum -= preNum
		} else {
			sum += preNum
		}
		preNum = num
	}
	sum += preNum
	return sum
}

func getValue(ch byte) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func intToRoman(num int) string {
	romes := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var rm string
	for i := 0; i < len(numbers); i++ {
		for num >= numbers[i] {
			num -= numbers[i]
			rm += romes[i]
		}
	}
	return rm
}

func intToRoman2(num int) bytes.Buffer {
	romes := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	rm := bytes.Buffer{}
	for i := 0; i < len(numbers); i++ {
		for num >= numbers[i] {
			num -= numbers[i]
			rm.WriteString(romes[i])
		}
	}
	return rm
}
