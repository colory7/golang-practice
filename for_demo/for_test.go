package postgresql_demo

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {

	i := 0
	for ; i < 1; i += 1 {
		if i == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(i)
}

func TestFor2(t *testing.T) {

	i := 0
	for i < 1 {
		if i == 0 {
			continue
		}
		fmt.Println(i)
		i += 1
	}
	fmt.Println(i)
}
