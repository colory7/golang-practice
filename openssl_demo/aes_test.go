package openssl_demo

import (
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
	"testing"
)

func TestAes(t *testing.T) {
	src := []byte("123456")
	key := []byte("1234567890123456")
	//key := []byte("123456789012345")
	dst, _ := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	fmt.Printf(base64.StdEncoding.EncodeToString(dst)) // yXVUkR45PFz0UfpbDB8/ew==
	fmt.Println()
	dst, _ = openssl.AesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	fmt.Println(string(dst)) // 123456
}

func TestAes2(t *testing.T) {
	src := []byte("123456")
	key := []byte("1234567890123456")
	iv := []byte("1234567890123456")
	dst, _ := openssl.AesCBCEncrypt(src, key, iv, openssl.PKCS7_PADDING)
	fmt.Println(base64.StdEncoding.EncodeToString(dst)) // 1jdzWuniG6UMtoa3T6uNLA==

	dst, _ = openssl.AesCBCDecrypt(dst, key, iv, openssl.PKCS7_PADDING)
	fmt.Println(string(dst)) // 123456
}
