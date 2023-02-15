package hello

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuite(t *testing.T) {
	tests := []struct {
		name string
	}{
		{`a1`},
		{name: "37"},
		{"56"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.name)
		})
	}
}

func TestSuite2(t *testing.T) {
	tests := []struct {
		i  int
		ch string

		exception bool
	}{
		{1, "", false},
		{1, "", true},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			bb, err := aa(test.ch)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println(bb)
			}
		})
	}
}

func TestSuite3(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		expected  string
		exception bool
	}{
		{1, "", "", false},
		{2, "e", "", true},
		{3, "2", "", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := aa(test.ch)

			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
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

func aa(ch string) (string, error) {
	if ch == "e" {
		return "", errors.New("test error")
	}
	return "", nil
}
