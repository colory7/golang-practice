package grammer_demo

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	d := 3
	ref(&d)

	fmt.Println(d)
}

func ref(d *int) interface{} {
	*d = 2
	return nil
}
