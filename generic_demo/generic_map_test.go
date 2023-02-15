package generic_demo

import (
	"fmt"
	"testing"
)

type M[K string, V any] map[K]V

func TestMap(t *testing.T) {
	m1 := M[string, int]{"key": 1}
	m1["key"] = 2

	m2 := M[string, string]{"key": "value"}
	m2["key"] = "dudu"
	fmt.Println(m1, m2)
}
