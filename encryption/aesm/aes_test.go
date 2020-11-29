package aesm

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	// 加解密Key
	key := []byte("A123456789012345")
	// 明文
	plaintext := []byte("%7B%22Name%22%3A%22Test%22%2C%22ID%22%3A%22A123456789%22%7D")
	// 初始向量
	iv := []byte("B123456789012345")

	// 加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}

	//列印加密base64後密碼
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
}

func Test_Decrypt(t *testing.T) {
	// 加解密Key
	key := []byte("A123456789012345")
	// 明文
	plaintext := []byte("%7B%22Name%22%3A%22Test%22%2C%22ID%22%3A%22A123456789%22%7D")
	// 初始向量
	iv := []byte("B123456789012345")

	// 加密
	ciphertext, err := AesEncrypt(plaintext, key, iv)
	if err != nil {
		panic(err)
	}

	// 解密
	res, err := AesDecrypt(ciphertext, key, iv)
	if err != nil {
		panic(err)
	}

	//列印解密明文
	fmt.Println(string(res))
}
