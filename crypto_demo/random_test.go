package crypto_demo

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(1)
	n := 4
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}
