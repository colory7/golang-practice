package timeutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuite(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		{1, "20220810", false},
		{1, "2022-08-11", false},
		{1, "2022/08/10", false},
		{1, "20170623113939", false},
		{1, "2017-06-23 11:39:39", false},
		{1, "2017/06/23 11:39:39", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, _, err := ParseTime2(test.ch)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println(tm)
			}
		})
	}
}

func ToChar() {

}

func TestToChar(t *testing.T) {

}
