package func_demo

import (
	"fmt"
	"testing"
)

func TestDefaultReturn(t *testing.T) {
	fmt.Println(defaultReturn())
}

func defaultReturn() string {
	var aa string
	return aa
}
