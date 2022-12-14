package make_demo

import (
	"fmt"
	"testing"
)

func TestMake(t *testing.T) {
	c1 := make(chan string)
	m1 := make(map[string]string)
	s1 := make([]string, 100)

	fmt.Println(c1)
	fmt.Println(m1["aa"])
	fmt.Println(s1[0])
}
