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
