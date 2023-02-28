package crypto_demo

import (
	"encoding/base64"
	"fmt"
	"testing"
)

const (
	base64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
)

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src []byte) []byte { //编码
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) { //解码
	return coder.DecodeString(string(src))
}

func TestBase64(t *testing.T) {
	aa := "abc我们t3你"
	bb := Base64Encode([]byte(aa))
	fmt.Println(bb)

	fmt.Println(Base64Encode(bb))

}
