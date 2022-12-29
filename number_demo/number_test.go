package number_demo

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

func TestFormat(t *testing.T) {
	var n1 = fmt.Sprintf("%x", 0x4D7953514C)
	fmt.Println(n1)

	var n2 = fmt.Sprintf("%b", 0x4D7953514C)
	fmt.Println(n2)
}

func TestHexToDecimal(t *testing.T) {
	h := "123a"
	num, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	fmt.Println(BytesToString(num))
}

func TestHexToDecimal2(t *testing.T) {
	h := "-123a"
	num, err := Hex2Dec(h)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}

func Hex2Dec(val string) (int, error) {
	n, err := strconv.ParseInt(val, 16, 64)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func TestDecimalToHex(t *testing.T) {
	var w int64 = -4666
	s := strconv.FormatInt(w, 16)
	println(s)
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func TestTmp(t *testing.T) {
	d := 198
	d /= 10
	fmt.Println(d)
	d /= 10
	fmt.Println(d)
	d /= 10
	fmt.Println(d)
	d /= 10
	fmt.Println(d)
}

func TestNum(t *testing.T) {
	var num float64 = 3.1415926
	str := strconv.FormatFloat(num, 'E', -1, 64)
	fmt.Printf("type:%T,value:%v\n ", str, str)

	str2 := strconv.FormatFloat(num, 'f', -1, 64)
	fmt.Printf("type:%T,value:%v\n ", str2, str2)
}

func TestNum2(t *testing.T) {
	var num = "3.1415926"

	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		panic(err)
	}
	s := strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Println(s)
}
