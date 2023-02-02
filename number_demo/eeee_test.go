package number_demo

import (
	"fmt"
	"strconv"
	"testing"
)

func TestEEEE(t *testing.T) {
	f, err := strconv.ParseFloat("-13e-2", 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}
