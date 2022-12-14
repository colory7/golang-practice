package make_demo

import (
	"fmt"
	"testing"
)

func TestMapCreate(t *testing.T) {
	//1. 声明
	var m1 map[string]string
	fmt.Printf("类型：%T, 长度：%d, 值：%#v, 地址:%p \n", m1, len(m1), m1, &m1)
	// 类型：map[string]string, 长度：0, 值：map[string]string(nil), 地址:0xc000006028

	//2. 使用map[keyType]valueType
	var m2 map[string]string = map[string]string{}
	fmt.Printf("类型：%T, 长度：%d, 值：%#v, 地址:%p \n", m2, len(m2), m2, &m2)
	// 类型：map[string]string, 长度：0, 值：map[string]string{}, 地址:0xc000006038

	//3. 使用make
	var m3 map[string]string = make(map[string]string)
	fmt.Printf("类型：%T, 长度：%d, 值：%#v, 地址:%p \n", m3, len(m3), m3, &m3)
	// 类型：map[string]string, 长度：0, 值：map[string]string{}, 地址:0xc000006040

}
