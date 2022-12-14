package postgresql_demo

import (
	"fmt"
	"testing"
)

func TestChar(t *testing.T) {
	var c = "I love 中国"
	fmt.Println(c[2])         //108
	fmt.Println(string(c[2])) // l
	fmt.Println(string(c[7])) // 乱码 三个字节才能造成中文
	fmt.Println([]byte(c))    // [73 32 108 111 118 101 32 228 184 173 229 155 189]

	// 转化成rune类型数组
	c2 := []rune(c)
	fmt.Println(c2)            // [73 32 108 111 118 101 32 20013 22269]
	fmt.Println(string(c2[8])) // 中
	fmt.Println(string(c2[8])) // 国
}

func TestChar2(t *testing.T) {

	v4 := `床前明月光,
		疑似地上霜.
			举着望明月,
		低头思故乡.
	`

	v6 := []rune(v4)
	v7 := "故"
	for k, v := range v6 {
		if string(v) == v7 {
			fmt.Printf("找到字符---\"%s\",\n其索引为%d\n", v7, k)
			fmt.Printf("%d--%c--%T\n", k, v, v)
		}
	}

}
