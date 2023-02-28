package mysql_func_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang_practice/util_demo/stringutils"
	"testing"
)

func TestMid(t *testing.T) {
	tests := []struct {
		i         int
		src       string
		pos       int
		count     int
		expected  string
		exception bool
	}{
		{0, "ty在手w机hello上阅world读所有教程", 1, 999, "ty在手w机hello上阅world读所有教程", false},
		{0, "ty在手w机hello上阅world读所有教程", 3, 1, "在", false},
		{0, "ty在手w机hello上阅world读所有教程", 3, 2, "在手", false},
		{0, "ty在手w机hello上阅world读所有教程", 1, 1, "t", false},
		{0, "ty在手w机hello上阅world读所有教程", 1, 22, "ty在手w机hello上阅world读所有教", false},
		{0, "ty在手w机hello上阅world读所有教程", 1, 23, "ty在手w机hello上阅world读所有教程", false},
		{0, "ty在手w机hello上阅world读所有教程", 1, 24, "ty在手w机hello上阅world读所有教程", false},
		{0, "ty在手w机hello上阅world读所有教程", 2, 1, "y", false},
		{0, "ty在手w机hello上阅world读所有教程", 2, 2, "y在", false},
		{0, "ty在手w机hello上阅world读所有教程", 2, 3, "y在手", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 0, "", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 1, "", false},
		{0, "ty在手w机hello上阅world读所有教程", 1, 1, "t", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 1, "", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 2, "", false},
		{0, "ty在手w机hello上阅world读所有教程", -5, 6, "d读所有教程", false},
		{0, "ty在手w机hello上阅world读所有教程", -5, 5, "d读所有教", false},
		{0, "ty在手w机hello上阅world读所有教程", -5, 7, "d读所有教程", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := Mid(test.src, test.pos, test.count)
			if test.exception {
				fmt.Println("expected exception: ", err)
				assert.Error(t, err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println("actual:   ", actual)
				if actual != test.expected {
					fmt.Println("expected: ", test.expected)
					t.Fail()
				}
			}
		})
	}
}

// Mid 字符串截取
// @Param str
// @Param pos 数值范围从-1或1开始
// @Param len 可选,如果不指定，则提取到原字符串的结尾
// @Return 子字符串
func Mid(str string, pos, count int) (string, error) {
	if pos == 0 {
		return "", nil
	} else {
		return stringutils.SubstringByCount(str, pos-1, count)
	}
}
