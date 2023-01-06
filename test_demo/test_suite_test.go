package hello

import (
	"fmt"
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
