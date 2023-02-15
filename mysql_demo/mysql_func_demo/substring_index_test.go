package mysql_func_demo

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestSubstringIndex(t *testing.T) {
	tests := []struct {
		i        int
		s        string
		sep      string
		count    int
		expected string
	}{
		{0, "www.mysql.com", ".", -4, "www.mysql.com"},
		{0, "www.mysql.com", ".", -3, "www.mysql.com"},
		{0, "www.mysql.com", ".", -2, "mysql.com"},
		{0, "www.mysql.com", ".", -1, "com"},
		{0, "www.mysql.com", ".", 0, ""},
		{0, "www.mysql.com", ".", 1, "www"},
		{0, "www.mysql.com", ".", 2, "www.mysql"},
		{0, "www.mysql.com", ".", 3, "www.mysql.com"},
		{0, "www.mysql.com", ".", 4, "www.mysql.com"},

		{0, "文章点知识点e相关知识d", "点", -4, "文章点知识点e相关知识d"},
		{0, "文章点知识点e相关知识d", "点", -3, "文章点知识点e相关知识d"},
		{0, "文章点知识点e相关知识d", "点", -2, "知识点e相关知识d"},
		{0, "文章点知识点e相关知识d", "点", -1, "e相关知识d"},
		{0, "文章点知识点e相关知识d", "点", 0, ""},
		{0, "文章点知识点e相关知识d", "点", 1, "文章"},
		{0, "文章点知识点e相关知识d", "点", 2, "文章点知识"},
		{0, "文章点知识点e相关知识d", "点", 3, "文章点知识点e相关知识d"},
		{0, "文章点知识点e相关知识d", "点", 4, "文章点知识点e相关知识d"},
		{0, "文章点知识点e相关知识d", "点", 2, "文章点知识"},
		{0, "文章点知识点e相关知识d", "知识", 2, "文章点知识点e相关"},
		{0, "文章点知识点e相关知识d", "知识", -2, "点e相关知识d"},
		{0, "文章点知识点e相关知识d", "我们", 1, "文章点知识点e相关知识d"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual := SubstringIndex(test.s, test.sep, test.count)
			fmt.Println("actual:   ", actual, "\nexpected: ", test.expected)
			if actual != test.expected {
				t.Fail()
			} else {
			}
		})
	}
}

func Test(t *testing.T) {
	s := "www.mysql.com"

	idx1 := strings.Index(s[0:], ".")

	idx2 := strings.Index(s[idx1:], ".")
	fmt.Println(idx2)
	fmt.Println(s[idx2:])
}

func Test2(t *testing.T) {
	s1 := "www.mysql.com"
	s2 := "."
	var bs1 = []byte(s1)
	var bs2 = []byte(s2)

	aa := bytes.Index(bs1, bs2)
	fmt.Println(aa)
}

func Test3(t *testing.T) {
	ids := "www.mysql.com"
	idList := strings.Split(ids, ".")
	for i := 0; i < len(idList); i++ {
		fmt.Println(idList[i])
	}
}

func Test4(t *testing.T) {
	ids := "文章点知识点e相关知识d"
	idList := strings.Split(ids, "知识")
	for i := 0; i < len(idList); i++ {
		fmt.Println(idList[i])
	}
}

func SubstringIndex2(str string, delim string, count int) string {
	strLen := len(str)
	delimLen := len(delim)
	if strLen == 0 || delimLen == 0 || count == 0 {
		return ""
	}
	if count > 0 {
		for i := 0; i < strLen && count > 0; i++ {
			if str[i] == delim[0] {
				if strLen-i >= delimLen && str[i:i+delimLen] == delim {
					count--
					if count == 0 {
						return str[:i]
					}
					i += delimLen - 1
				}
			}
		}
		return str
	} else {
		for i := strLen - 1; i >= 0 && count < 0; i-- {
			if str[i] == delim[0] {
				if i+delimLen <= strLen && str[i:i+delimLen] == delim {
					count++
					if count == 0 {
						return str[i+delimLen:]
					}
					i -= delimLen - 1
				}
			}
		}
		return str
	}
}
