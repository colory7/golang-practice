package mysql_func_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOctString(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		expected  string
		exception bool
	}{
		{1, "-2", "", true},
		{2, "-1", "", true},
		{3, "-1.2", "", true},
		{4, "0", "0", false},
		{4, "1", "1", false},
		{4, "1.2", "1", false},
		{4, "12", "14", false},
		{4, "64", "100", false},
		{4, "10", "12", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := Oct(test.ch)
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

func TestOctFloat32(t *testing.T) {
	tests := []struct {
		i         int
		f         float32
		expected  string
		exception bool
	}{
		{1, -2, "", true},
		{2, -1, "", true},
		{3, -1.2, "", true},
		{4, 0, "0", false},
		{4, 1, "1", false},
		{4, 1.2, "1", false},
		{4, 12, "14", false},
		{4, 64, "100", false},
		{4, 10, "12", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := Oct(test.f)
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

func TestOctFloat64(t *testing.T) {
	tests := []struct {
		i         int
		f         float32
		expected  string
		exception bool
	}{
		{1, -2, "", true},
		{2, -1, "", true},
		{3, -1.2, "", true},
		{4, 0, "0", false},
		{4, 1, "1", false},
		{4, 1.2, "1", false},
		{4, 12, "14", false},
		{4, 64, "100", false},
		{4, 10, "12", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := Oct(test.f)
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

func TestOctInt32(t *testing.T) {
	tests := []struct {
		i         int
		f         int
		expected  string
		exception bool
	}{
		{1, -2, "", true},
		{2, -1, "", true},
		{4, 0, "0", false},
		{4, 1, "1", false},
		{4, 12, "14", false},
		{4, 64, "100", false},
		{4, 10, "12", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := Oct(test.f)
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

func TestOctInt64(t *testing.T) {
	tests := []struct {
		i         int
		f         int64
		expected  string
		exception bool
	}{
		{1, -2, "", true},
		{2, -1, "", true},
		{4, 0, "0", false},
		{4, 1, "1", false},
		{4, 12, "14", false},
		{4, 64, "100", false},
		{4, 10, "12", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			actual, err := Oct(test.f)
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
