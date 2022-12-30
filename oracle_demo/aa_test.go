package oracle_demo

import (
	"fmt"
	"testing"
)

type AA int

const CC AA = iota
const DD AA = iota + 1
const EE AA = iota + 1

func TestAA(t *testing.T) {
	fmt.Println(CC)
	fmt.Println(DD)
	fmt.Println(EE)
}
