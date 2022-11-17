package tp_demo

import (
	"fmt"
	"golang_practice/tp_demo/timeutil"
	"testing"
	"time"
)

func TestParseTime(xxx *testing.T) {
	dt1 := "202101021504"
	//dt1 := "20210102150405.999999"
	//dt1 := "05"
	//dt1 := "05.9999"
	//dt1 := "0102150405.999999"
	t, err := time.Parse(timeutil.TimestampNumWithoutTZFormat, dt1)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)

}

func TestParseTime2(xxx *testing.T) {
	dt1 := "2022-08-11"
	dt2 := "20:30:40"
	t1, format1, err := timeutil.ParseTime(dt1)
	if err != nil {
		panic(err)
	}

	t2, format2, err := timeutil.ParseTime(dt2)
	if err != nil {
		panic(err)
	}

	fmt.Println(format1)
	fmt.Println(format2)
	fmt.Println(t1)
	fmt.Println(t2)

}

func TestParseTime3(xxx *testing.T) {
	dt1 := "2022-08-11"
	t, err := time.Parse(timeutil.DateFormat, dt1)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.UnixMilli())
	fmt.Println(t.UnixMicro())
	fmt.Println(t.UnixNano())
}

func TestDiffDate(t *testing.T) {
	date1 := "2022-08-11"
	date2 := "2022-08-10"
	t1, _, err := timeutil.ParseTime(date1)
	if err != nil {
		panic(err)
	}
	t2, _, err := timeutil.ParseTime(date2)
	if err != nil {
		panic(err)
	}

	n := int(t1.Sub(t2).Hours() / 24)

	fmt.Println(n)
}

func TestDiffDate2(t *testing.T) {
	date1 := "2022-08-11"
	date2 := "2022-08-11"
	t1, _, err := timeutil.ParseTime(date1)
	if err != nil {
		panic(err)
	}
	t2, _, err := timeutil.ParseTime(date2)
	if err != nil {
		panic(err)
	}

	n := int(t1.Sub(t2).Hours() / 24)

	fmt.Println(n)
}

func TestDiffDate3(t *testing.T) {
	date1 := "2022-08-11"
	date2 := "1891-03-12"
	t1, _, err := timeutil.ParseTime(date1)
	if err != nil {
		panic(err)
	}
	t2, _, err := timeutil.ParseTime(date2)
	if err != nil {
		panic(err)
	}

	n := int(t1.Sub(t2).Hours() / 24)

	fmt.Println(n)
}

func TestDiffInterval(t *testing.T) {
	date1 := "2022-08-11"
	t1, _, err := timeutil.ParseTime(date1)
	if err != nil {
		panic(err)
	}

	fmt.Println(t1.Day())
	fmt.Println(t1.Hour())

}
