package postgresql_demo

import (
	"bytes"
	"fmt"
	"testing"
)

func TestASCII(t *testing.T) {
	result := bytes.Buffer{}
	result.WriteByte(97)
	fmt.Println(result.String())

	result = bytes.Buffer{}
	result.WriteByte(1)
	fmt.Println(result.String())
	fmt.Println(result.String() == "A")
	fmt.Println(rune('0'))

	fmt.Println(byte(uint(0)))

	fmt.Println(string(byte(47)))
	fmt.Println(string(byte(48)))
	fmt.Println(string(byte(97)))
	fmt.Println(string(byte(0)))
	fmt.Println(string(byte(27)))
}

func TestAscii2(t *testing.T) {
	fmt.Println(string(byte(47)))
	fmt.Println(string(byte(48)))
	fmt.Println(string(byte(97)))
	fmt.Println(string(byte(0)))
	fmt.Println(string(byte(27)))
}
