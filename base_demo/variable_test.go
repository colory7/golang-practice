package base_demo

import (
	"fmt"
	"strconv"
	"testing"
)

func TestVariable(t *testing.T) {
	s := "3"
	var v1 int

	if len("accc") > 2 {
		var err error
		v1, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(v1)
}
