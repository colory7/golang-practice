package mysql_func_demo

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrcmp(t *testing.T) {
	tests := []struct {
		i        int
		ch1      string
		ch2      string
		expected int
	}{
		{1, "", "", 0},
		{2, "", "a", -1},
		{3, "a", "", 1},
		{5, "0", "0", 0},
		{6, "0", "1", -1},
		{7, "0", "-1", 1},
		{8, "ab", "abc", -1},
		{8, "abcd", "abc", 1},
		{8, "abce", "abcd", 1},
		{8, "abce", "abcda", 1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual := strings.Compare(test.ch1, test.ch2)
			fmt.Println("actual:   ", actual)
			if actual != test.expected {
				fmt.Println("expected: ", test.expected)
				t.Fail()
			}
		})
	}
}
