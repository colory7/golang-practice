package generic_demo

import (
	"fmt"
	"testing"
)

type C[T any] chan T

func TestChannel(t *testing.T) {
	a := make(C[int], 10)
	a <- 1
	fmt.Println(<-a)
	b := make(C[string], 10)
	b <- "dudu"
	fmt.Println(<-b)
}
