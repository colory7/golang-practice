package type_cast_demo

import (
	"fmt"
	"strconv"
	"testing"
)

// https://www.cnblogs.com/hello-/articles/15498398.html
func Test(t *testing.T) {
	input := 20220228101010.233232354
	s := strconv.FormatFloat(input, 'f', 6, 64)
	fmt.Println(s)
}

func TestCastFloat(t *testing.T) {
	val := 22.35678434
	fmt.Println(int64(val * 1000))
}
