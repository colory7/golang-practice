package hello

import (
	"fmt"
	"testing"
)

func TestPrefix(t *testing.T) {
	//NUM_F_DOLLAR = 1 << 2
	//NUM_F_B      = 1 << 3
	//NUM_F_C      = 1 << 4
	//NUM_F_L      = 1 << 5
	//NUM_F_U      = 1 << 6

	fmt.Println(1<<4 < 1<<6)
	fmt.Println(1<<4 > 1<<2)
}
