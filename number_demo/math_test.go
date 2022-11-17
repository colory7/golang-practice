package number_demo

import (
	"fmt"
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	fmt.Println(math.Round(2.4))
	fmt.Println(math.Round(2.5))
	fmt.Println(math.Round(2.9))

	fmt.Println(math.Round(2.45))

	fmt.Println(math.Trunc(math.Round(2.45)))

	fmt.Println(math.Round(889898.999666))

}
