package mysql_func_demo

import (
	"fmt"
	"testing"
)

func TestFindInSet(t *testing.T) {
	tests := []struct {
		i        int
		s        string
		strSet   string
		expected int
	}{
		{0, "b", "abcd", 0},
		{0, "b", "abcd,c,e", 0},
		{0, "b", "abcd,c,b", 3},
		{0, "我们", "abcd,c,b", 0},
		{0, "我们", "abcd,知识,c,b", 0},
		{0, "我们", "abcd,知识,c,我们b", 0},
		{0, "我们", "abcd,知识,c,我们b,我们", 5},
		{0, "我们", "abcd,知识,c,我们b,我们,我们", 5},
		{0, "我们", "aa", 0},
		{0, "我们", "我", 0},
		{0, "我们", "我们", 1},
		{0, "我们", "我们的", 0},
		{0, "我们的", "我们", 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual := FindInSet(test.s, test.strSet)

			fmt.Println("actual:   ", actual)
			if actual != test.expected {
				fmt.Println("expected: ", test.expected)
				t.Fail()
			}
		})
	}
}
