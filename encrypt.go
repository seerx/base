package base

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

//MD5 MD5 编码
func MD5(val string) string {
	data := []byte(val)
	// has := md5.Sum(data)
	return MD5It(data)
}

// MD5It MD5 编码
func MD5It(data []byte) string {
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// MD5s MD5 加密多个字符串，以 : 分割
func MD5s(strs ...string) string {
	// str := ""
	str := strings.Join(strs, ":")
	// for _, val := range strs {
	// 	if str == "" {
	// 		str = val
	// 	} else {
	// 		str += ":" + val
	// 	}
	// }
	return MD5(str)
}

// MD5f 格式化后再进行 MD5 编码
func MD5f(format string, args ...interface{}) string {
	return MD5(fmt.Sprintf(format, args...))
}

// AESEncrypt AES 加密
// k 32 个长度的字符串
// n 24 个长度的字符串
func AESEncrypt(src, k, n string) (string, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	plaintext := []byte(src)
	nonce, err := hex.DecodeString(n)
	if err != nil {
		return "", err
	}

	data, err := AESEncryptData(plaintext, key, nonce)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", data), nil
	//block, err := aes.NewCipher(key)
	//if err != nil {
	//	return "", err
	//}
	//
	//aesGcm, err := cipher.NewGCM(block)
	//if err != nil {
	//	return "", err
	//}
	//
	//cipherText := aesGcm.Seal(nil, nonce, plaintext, nil)

	//return fmt.Sprintf("%x", cipherText), nil
}

func AESEncryptData(src, key, nonce []byte) ([]byte, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	//key := []byte(k)
	plaintext := src //[]byte(src)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//nonce, _ := hex.DecodeString(n)

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	cipherText := aesGcm.Seal(nil, nonce, plaintext, nil)

	return cipherText, nil
}

// AESDecrypt AES 解密
func AESDecrypt(src, k, n string) (string, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	cipherText, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	nonce, err := hex.DecodeString(n)
	if err != nil {
		return "", err
	}

	data, err := AESDecryptData(cipherText, key, nonce)
	if err != nil {
		return "", err
	}
	return string(data), nil
	//block, err := aes.NewCipher(key)
	//if err != nil {
	//	return "", err
	//}
	//
	//aesGcm, err := cipher.NewGCM(block)
	//if err != nil {
	//	return "", err
	//}
	//
	//plainText, err := aesGcm.Open(nil, nonce, cipherText, nil)
	//if err != nil {
	//	return "", err
	//}

	//return string(plainText), nil
}

// AESDecrypt AES 解密
func AESDecryptData(src, key, nonce []byte) ([]byte, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	//key := []byte(k)
	//cipherText, _ := hex.DecodeString(src)
	//cipherText := src
	//nonce, _ := hex.DecodeString(n)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainText, err := aesGcm.Open(nil, nonce, src, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
