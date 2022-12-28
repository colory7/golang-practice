package oracle_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestToDate(txx *testing.T) {
	dch := "2018-05-06"
	format := "YYYY-MM-DD"
	tm, err := ToDate(dch, format)
	if err != nil {
		panic(err)
	}

	fmt.Println(*tm)
}

func TestToDate2(txx *testing.T) {
	dch := "20180506"
	format := "YYYYMMDD"
	tm, err := ToDate(dch, format)
	if err != nil {
		panic(err)
	}

	fmt.Println(*tm)
}

func TestToChar(t *testing.T) {
	tm := time.Date(2017, 02, 27, 20, 20, 20, 20, time.Local)
	fmt.Println(tm)
	format := "YYYY-MM-DD"

	str, err := ToDatetimeChar(tm, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}

func TestToNumber(t *testing.T) {
	numParam := "123456"
	format := "999999EEEE"

	numResult, err := ToNumberByStr(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}
