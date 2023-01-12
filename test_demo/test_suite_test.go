package hello

import (
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
		i         int
		ch        string
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

func aa(ch string) (string, error) {
	return "", nil
}
