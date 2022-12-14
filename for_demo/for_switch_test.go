package postgresql_demo

import (
	"fmt"
	"testing"
)

func TestAA(t *testing.T) {
	bb := 3
	cc := "uiojkl"

	fmt.Println(string(cc[0]))
	fmt.Println(cc[0:1])
	for i, v := range cc {
		fmt.Println(i, string(v))
		switch bb {
		case 3:
			switch bb {
			case 3:
				fmt.Println(222)
				continue
			}
			fmt.Println(333)
			break
		case 2:
			break
		}

	}

}
