package tokencrypto

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"

	"github.com/seerx/base"
)

// Encypt 加密令牌
func Encypt(token interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(token); err != nil {
		return "", err
	}

	// 生成 nonce
	nonce, err := newNonce()
	if err != nil {
		return "", err
	}

	data, err := base.AESEncryptData(buf.Bytes(), key, nonce)
	if err != nil {
		return "", err
	}

	// body := fmt.Sprintf("%s%s",
	// 	base64.StdEncoding.EncodeToString(data),
	// 	base64.StdEncoding.EncodeToString(nonce))
	body := append(data, nonce...)

	return base64.StdEncoding.EncodeToString(body), nil
}

// Decrypt 解密令牌
func Decrypt(data string, token interface{}) error {
	body, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	ln := len(body)
	if ln <= 12 {
		return errors.New("Invalid token data")
	}
	tkData := body[:ln-12]
	nData := body[ln-12:]

	tkData, err = base.AESDecryptData(tkData, key, nData)
	if err != nil {
		return err
	}
	// var tk Token
	decoder := gob.NewDecoder(bytes.NewBuffer(tkData))
	if err := decoder.Decode(token); err != nil {
		return err
	}
	return nil
}
