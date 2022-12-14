package slice_demo

import (
	"fmt"
	"testing"
)

func TestSliceByte(t *testing.T) {
	bs := []byte{}
	bs = append(bs, 97)
	bs = append(bs, 98)
	bs = append(bs, 100)

	fmt.Println(string(bs))
}
