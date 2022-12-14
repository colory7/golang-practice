package base_demo

import (
	"strconv"
	"testing"
)

func TestNumberConvertStr(t *testing.T) {
	str := strconv.Itoa(12)
	println("str + '23' = ", str+"23")
}

func TestStrConvertNumber(t *testing.T) {
	n, err := strconv.Atoi("123")
	if err != nil {
		panic(err)
	}
	println("n = ", n)
}

func TestStrConvertNumber2(t *testing.T) {
	n, err := strconv.ParseFloat("2.3", 64)
	if err != nil {
		panic(err)
	}
	println(n)
}

func TestStrConvertNumber3(t *testing.T) {
	n, err := strconv.ParseFloat("1e5", 64)
	if err != nil {
		panic(err)
	}
	println(n)
}
