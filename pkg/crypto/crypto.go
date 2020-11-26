package crypto

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"

	"github.com/seerx/base/pkg/log"
)

// Encypt 加密对象
func Encypt(obj interface{}, key []byte) (string, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(obj); err != nil {
		return "", err
	}

	data, err := EncyptBytesRandom(buf.Bytes(), key)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// // EncyptRandom 加密
// func EncyptRandom(token interface{}, key []byte) (string, error) {
// 	var buf bytes.Buffer
// 	encoder := gob.NewEncoder(&buf)
// 	if err := encoder.Encode(token); err != nil {
// 		return "", err
// 	}

// 	data, err := EncyptBytes(buf.Bytes(), key)
// 	if err != nil {
// 		return "", err
// 	}

// 	return base64.StdEncoding.EncodeToString(data), nil
// }

// Decrypt 解密对象
func Decrypt(data string, key []byte, obj interface{}) error {
	buf, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.WithError(err).Error("Base64 解码错误")
		return err
	}
	ln := len(buf)
	if ln <= 12 {
		return errors.New("Invalid token data")
	}
	body := buf[:ln-12]
	nonce := buf[ln-12:]
	res, err := DecryptBytes(body, key, nonce)
	if err != nil {
		log.WithError(err).Error("AES 解密错误")
		return err
	}
	// var tk Token
	decoder := gob.NewDecoder(bytes.NewBuffer(res))
	if err := decoder.Decode(obj); err != nil {
		log.WithError(err).Error("数据转换为令牌错误")
		return err
	}
	return nil
}
