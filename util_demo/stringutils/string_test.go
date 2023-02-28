package stringutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubstringByCount(t *testing.T) {
	tests := []struct {
		i         int
		src       string
		pos       int
		count     int
		expected  string
		exception bool
	}{
		{0, "ty在手w机hello上阅world读所有教程", 0, 999, "ty在手w机hello上阅world读所有教程", false},
		{0, "ty在手w机hello上阅world读所有教程", 2, 1, "在", false},
		{0, "ty在手w机hello上阅world读所有教程", 2, 2, "在手", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 1, "t", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 22, "ty在手w机hello上阅world读所有教", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 23, "ty在手w机hello上阅world读所有教程", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 24, "ty在手w机hello上阅world读所有教程", false},
		{0, "ty在手w机hello上阅world读所有教程", 1, 1, "y", false},
		{0, "ty在手w机hello上阅world读所有教程", 1, 2, "y在", false},
		{0, "ty在手w机hello上阅world读所有教程", 1, 3, "y在手", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 1, "t", false},
		{0, "ty在手w机hello上阅world读所有教程", -1, 1, "程", false},
		{0, "ty在手w机hello上阅world读所有教程", -1, 2, "程", false},
		{0, "ty在手w机hello上阅world读所有教程", -6, 6, "d读所有教程", false},
		{0, "ty在手w机hello上阅world读所有教程", -6, 5, "d读所有教", false},
		{0, "ty在手w机hello上阅world读所有教程", -6, 7, "d读所有教程", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := SubstringByCount(test.src, test.pos, test.count)
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

func TestSubstring(t *testing.T) {
	tests := []struct {
		i         int
		src       string
		start     int
		end       int
		expected  string
		exception bool
	}{
		{0, "ty在手w机hello上阅world读所有教程", 2, 1, "t", true},
		{0, "ty在手w机hello上阅world读所有教程", 0, 0, "t", true},
		{0, "ty在手w机hello上阅world读所有教程", -2, 1, "t", true},
		{0, "ty在手w机hello上阅world读所有教程", -1, 1, "t", true},
		{0, "ty在手w机hello上阅world读所有教程", 1, 1, "", true},
		{0, "ty在手w机hello上阅world读所有教程", 22, 22, "", true},
		{0, "ty在手w机hello上阅world读所有教程", 22, 23, "程", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 1, "t", false},
		{0, "ty在手w机hello上阅world读所有教程", 2, 5, "在手w", false},
		{0, "ty在手w机hello上阅world读所有教程", 0, 23, "ty在手w机hello上阅world读所有教程", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := Substring(test.src, test.start, test.end)
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
