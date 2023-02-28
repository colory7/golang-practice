package byte_demo

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestByte(t *testing.T) {
	str := "hello"
	println([]byte(str))
}

func TestByte2(t *testing.T) {
	type Bytes []byte
	s := "这是一个字符串垚"
	b := Bytes(s)
	fmt.Println(b)
}
func TestByte3(t *testing.T) {
	s := "这是一个字符串垚"
	b := []byte(s)
	fmt.Println(b)
}

func TestByte4(t *testing.T) {
	b := []byte{97, 232, 191, 153, 230, 152, 175, 228, 184, 128, 228, 184, 170, 229, 173, 151, 231, 172, 166, 228, 184, 178}
	s := string(b)
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[len(b)-1])
	fmt.Println(s)

	fmt.Println(s)

	s2 := string(s[len(b)-1])
	fmt.Println(s2)

}

func TestByte5(t *testing.T) {
	bytes := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	/// 小端模式
	fmt.Println(binary.LittleEndian.Uint16(bytes))
	fmt.Println(binary.LittleEndian.Uint32(bytes))
	fmt.Println(binary.LittleEndian.Uint64(bytes))
	fmt.Println("====")
	/// 大端模式
	fmt.Println(binary.BigEndian.Uint16(bytes))
	fmt.Println(binary.BigEndian.Uint32(bytes))
	fmt.Println(binary.BigEndian.Uint64(bytes))
}

func TestAscii(t *testing.T) {
	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii:%c %d\n", theme[i], theme[i])
	}
	fmt.Println("====")
	theme2 := "狙击 start"
	for _, v := range theme2 {
		fmt.Printf("unicode: %c %d\n", v, v)
	}

	fmt.Println("====")
	for k, v := range theme2 {
		fmt.Printf("unicode: %d:%d\n", k, v)
	}

}

func TestBoolByte(t *testing.T) {
	fmt.Println(bool(false))
}
