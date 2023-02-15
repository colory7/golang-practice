package mysql_func_demo

import (
	"fmt"
	"testing"
)

func TestOrdString(t *testing.T) {
	tests := []struct {
		i        int
		ch       string
		expected int
	}{
		{1, "", 0},
		{2, "0", 48},
		{3, "-1", 45},
		{5, "-1.3", 45},
		{6, "null", 110},
		{7, "A", 65},
		{8, "0", 48},
		{8, "1.2", 49},
		{8, "hello,world", 104},
		{8, "你", 14990752},
		{8, "你好fsfasfasfasfsbvaf", 14990752},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual := Ord(test.ch)
			fmt.Println("actual:   ", actual)
			if actual != test.expected {
				fmt.Println("expected: ", test.expected)
				t.Fail()
			}
		})
	}
}
