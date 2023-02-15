package mysql_func_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	NULL = ""
	null = ""
)

func TestMakeSet(t *testing.T) {
	tests := []struct {
		i         int
		index     int
		strSet    []string
		expected  string
		exception bool
	}{
		{0, 2, []string{}, "", false},
		{0, 2, []string{"hello"}, "", false},
		{0, 2, []string{"hello", "nice"}, "nice", false},
		{0, 2, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g", "h", "i"}, "nice", false},
		{0, 10, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g", "h", "i"}, "nice,world", false},
		{0, 10, []string{"hello", "nice", "world", NULL, "a", "b", "c", "d", "e", "f", "g", "h", "i"}, "nice", false},
		{0, 10, []string{"hello", "nice", "", NULL, "a", "b", "c", "d", "e", "f", "g", "h", "i"}, "nice", false},
		{0, 2561, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g"}, "hello,f", false},
		{0, 2561, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g", "h"}, "hello,f,h", false},
		{0, 3087, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g"}, "hello,nice,world,g", false},
		{0, 2561, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g", "h", "i"}, "hello,f,h", false},
		{0, 3087, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g", "h", "i"}, "hello,nice,world,g,h", false},
		{0, 3086, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g", "h", "i"}, "nice,world,g,h", false},
		{0, 99999999, []string{"hello", "nice"}, "hello,nice", false},
		{0, 99999999, []string{"hello", "nice", NULL, "world", "a", "b", "c", "d", "e", "f", "g", "h", "i"}, "hello,nice,world,a,b,c,d", false},
		//{0, 3087,[]string{1,32,56,7,8}, "1,32,56,7", false},
		//{0, 3087,[]string{1,32,"bce","2",8}, "1,32,bce,2", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := MakeSet(test.index, test.strSet)
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
