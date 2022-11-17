package tmp

import (
	"fmt"
	"runtime"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(runtime.GOMAXPROCS(0))
}
