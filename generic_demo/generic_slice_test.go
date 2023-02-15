package generic_demo

import "testing"

func TestSlice(t *testing.T) {
	v1 := vestor[int]{58, 1881}
	printslice(v1)
	v2 := vestor[string]{"dudu", "yiyi", "8号"}
	printslice(v2)
}
