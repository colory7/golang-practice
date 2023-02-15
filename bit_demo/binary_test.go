package bit_demo

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
)

func TestBuma(t *testing.T) {
	var a = []byte{0x15, 0x4B}
	var b = []byte{0xE0, 0x15}
	// fmt.Println(fmt.Sprintf("%.16b", 0x154B))
	fmt.Println(GetYuanMa(a), GetYuanMa(b))
}

// 当data为原码时，输出补码
func GetBuMa(data []byte) uint16 {
	var ym uint16
	bm := binary.BigEndian.Uint16(data)
	var bitNum = len(data) * 8
	f := "%." + strconv.Itoa(bitNum) + "b"
	bmStr := fmt.Sprintf(f, bm)
	if string(bmStr[0]) == "1" {
		ym = ^bm + 1
	} else {
		ym = bm
	}
	return ym
}

// 当data为补码时，输出原码。
// 原理: 补码的补码为原码
func GetYuanMa(data []byte) uint16 {
	return GetBuMa(data)
}
