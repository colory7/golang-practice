package string_demo

import (
	"fmt"
	"strconv"
	"strings"
)

func ReverseBytes(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}

func Utf8Index2(str, substr string) (int, int) {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return -1, asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos, asciiPos
		}
	}
	return pos, asciiPos
}

func Index(str, substr string, count int) int {
	return 0

}

func Utf8IndexOf(str, substr string, count int) int {
	if count <= 0 {
		return -1
	}
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	found := false
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			fmt.Println("pos" + strconv.Itoa(pos))
			count--
			if count == 0 {
				found = true
				break
			}
		}
	}
	if found {
		return pos
	} else {
		return -1
	}
}
