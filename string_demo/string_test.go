package string_demo

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	s := " aa cc  "
	s2 := strings.TrimSpace(s)
	fmt.Println(len(s))
	fmt.Println(len(s2))

}
