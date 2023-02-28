package mysql_func_demo

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBin(t *testing.T) {
	fmt.Println(strconv.FormatInt(12, 2))
	fmt.Println(strconv.FormatInt(-1, 2))

	fmt.Println(strconv.FormatInt(-1, 2))
	fmt.Println(strconv.FormatInt(-2, 2))
	fmt.Println(strconv.FormatInt(-3, 2))

}
