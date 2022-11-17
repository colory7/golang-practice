package number_demo

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	var n1 = fmt.Sprintf("%x", 0x4D7953514C)
	fmt.Println(n1)

	var n2 = fmt.Sprintf("%b", 0x4D7953514C)
	fmt.Println(n2)
}
