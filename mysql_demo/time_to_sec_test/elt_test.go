package mysql_func_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElt(t *testing.T) {
	tests := []struct {
		i         int
		index     int
		strSet    []string
		expected  string
		exception bool
	}{
		{0, 2, []string{"这ww", "是", "一", "个", "测", "试"}, "是", false},
		{0, 2, []string{"这ww", "是f", "一", "个", "测", "试"}, "是f", false},
		{0, 2, []string{"这ww", "是我们", "一", "个", "测", "试"}, "是我们", false},
		{0, -1, []string{"这ww", "是我们", "一", "个", "测", "试"}, "nice", true},
		{0, 0, []string{"women", "是我们", "一", "个", "测", ""}, "nice", true},
		{0, 6, []string{"women", "是我们", "一", "个", "测", ""}, "", false},
		{0, 7, []string{"women", "是我们", "一", "个", "测", "试"}, "nice", true},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := Elt(test.index, test.strSet)
			if test.exception {
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
