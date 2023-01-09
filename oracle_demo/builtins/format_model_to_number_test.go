package builtins

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuiteToNumber(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		format    string
		exception bool
	}{
		{1, "34,50", "999,99", false},
		{1, "12,4,548", "99G9G99B9", false},
		{1, "3450", "99999,", true},
		{1, "3450,", "99999,", false},
		{1, "12,4,548.689", "99G9G99B9.999", false},
		{1, "12,4,548.689", "9G9G99B9.999", true},
		{1, "12,4,548.689", "999G9G99B9D999", false},
		{1, "12,4,548.689", "RN", true},
		{1, "12,4,548", "XXXXXX", false},
		{1, "1ABC3EDEF", "XXXXXXxXXXX", false},
		{1, "11", "XXXXXX", false},
		{1, "124548", "XXXXXX", false},
		{1, ",48", "XXXXXX", false},
		{1, "48,", "XXXXXX", false},
		{1, "12,4,548.689", "XXXXXXXXXXXXXXXX", true},
		{1, "12,4,", "XXXXXXXXXXXXXXXX", false},
		{1, "124548", "XXXXXXA", true},
		{1, "1a", "x", true},
		{1, "12,4,548.689", "TM", true},
		{1, "12,4,548.689", "TM0", true},
		{1, "12,4,548.327", "99G9G999.999", false},
		{1, "12,4,548.689", "aa", true},
		{1, "12,4,548.689", "90", true},
		{1, ".48", ".99", false},
		{1, "48.", "99", false},
		{1, "48,", "99,", false},
		{1, ",48", ",99", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := ToNumber(test.ch, test.format)
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
