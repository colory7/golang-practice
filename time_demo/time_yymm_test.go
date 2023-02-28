package time_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestYYMM(tx *testing.T) {
	t, err := time.Parse("0601", "7001")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}

func TestSimple(t *testing.T) {

}
