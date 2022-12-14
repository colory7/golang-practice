package postgresql_demo

import (
	"bytes"
	"fmt"
	"testing"
)

func TestASCII(t *testing.T) {
	result := bytes.Buffer{}
	result.WriteByte(97)
	fmt.Println(result.String())
}
