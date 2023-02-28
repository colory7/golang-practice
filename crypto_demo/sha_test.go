package crypto_demo

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestSha(t *testing.T) {
	s := "sha1 this string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
