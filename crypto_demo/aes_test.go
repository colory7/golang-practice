package crypto_demo

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	orig := "hello world"
	key := "123456781234567812345678"
	fmt.Println("原文：", orig)

	encryptCode := AesEncryptString(orig, key)
	fmt.Println("密文：", encryptCode)

	decryptCode := AesDecryptString(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}

func TestAesBytes(t *testing.T) {
	orig := "hello world"
	key := "123456781234567812345678"
	fmt.Println("原文：", orig)

	encryptCode, err := AesEncrypt(orig, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文：", encryptCode)

	decryptCode, err := AesDecrypt(encryptCode, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密结果：", decryptCode)
}

func TestAesBytes2(t *testing.T) {
	orig := "123456"
	key := "salt22"
	fmt.Println("原文：", orig)

	encryptCode, err := AesEncrypt(orig, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("密文：", encryptCode)

	decryptCode, err := AesDecrypt(encryptCode, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密结果：", decryptCode)
}

func AesEncrypt(orig string, key string) ([]byte, error) {
	origData := []byte(orig)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, err
}

func AesDecrypt(cryted []byte, key string) (string, error) {
	k := []byte(key)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cryted))
	// 解密
	blockMode.CryptBlocks(orig, cryted)
	return string(PKCS7UnPadding(orig)), nil
}

func AesEncryptString(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	return base64.RawURLEncoding.EncodeToString(cryted)

}

func AesDecryptString(cryted string, key string) string {
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	crytedByte, _ := base64.RawURLEncoding.DecodeString(cryted)
	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

// 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
