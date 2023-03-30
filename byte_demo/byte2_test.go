package byte_demo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	binary.Write(bytesBuffer, binary.BigEndian, x)
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

func TestIntToBytes(t *testing.T) {
	b1 := IntToBytes(334367843356575672)
	fmt.Println(b1)
	fmt.Printf("%X\n", b1)
}

func IntToBytes2(a int) []byte {
	buf := make([]byte, 4)
	for i := 0; i < 4; i++ {
		buf[i] = uint8(a & 0xff)
		a = a >> 8
	}
	return buf
}
