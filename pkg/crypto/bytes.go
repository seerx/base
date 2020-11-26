package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"time"

	"github.com/seerx/base"
)

// EncyptBytesRandom 随机 nonce 加密
func EncyptBytesRandom(buf []byte, key []byte) ([]byte, error) {
	// 生成 nonce
	nonce, err := randomNonce()
	if err != nil {
		return nil, err
	}
	return EncyptBytes(buf, key, nonce, true)
}

// EncyptBytes 加密字节数组
func EncyptBytes(buf []byte, key []byte, nonce []byte, appendNonce bool) ([]byte, error) {
	// 生成 nonce
	// nonce, err := randomNonce()
	// if err != nil {
	// 	return nil, err
	// }

	data, err := base.AESEncryptData(buf, key, nonce)
	if err != nil {
		return nil, err
	}
	if !appendNonce {
		return data, nil
	}
	body := append(data, nonce...)

	return body, nil
}

// EncyptBytesToString 加密为字符串
func EncyptBytesToString(buf []byte, key []byte, nonce []byte, appendNonce bool) (string, error) {
	data, err := EncyptBytes(buf, key, nonce, appendNonce)
	if err != nil {
		// log.WithError(err).Error("AES 加密错误")
		return "", err
	}
	return base64.URLEncoding.EncodeToString(data), nil
	// return base64.StdEncoding.EncodeToString(data), nil
}

// EncyptBytesToStringRandom 使用随机 nonce 加密为字符串
func EncyptBytesToStringRandom(buf []byte, key []byte) (string, error) {
	data, err := EncyptBytesRandom(buf, key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

// DecryptBytes 解密令牌
func DecryptBytes(buf []byte, key []byte, nonce []byte) ([]byte, error) {
	// ln := len(buf)
	// if ln <= 12 {
	// 	return nil, errors.New("Invalid token data")
	// }
	// data := buf[:ln-12]
	// nData := buf[ln-12:]
	decryptedData, err := base.AESDecryptData(buf, key, nonce)
	if err != nil {
		// log.WithError(err).Error("AES 解密错误")
		return nil, err
	}
	return decryptedData, nil
	// var tk Token
	// decoder := gob.NewDecoder(bytes.NewBuffer(tkData))
	// if err := decoder.Decode(token); err != nil {
	// 	log.WithError(err).Error("数据转换为令牌错误")
	// 	return err
	// }
	// return nil
}

// DecryptBytesFromString 从字符串解密
func DecryptBytesFromString(val string, key []byte, nonce []byte) ([]byte, error) {
	buf, err := base64.URLEncoding.DecodeString(val)
	// data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, err
	}
	var data []byte
	if nonce == nil {
		// 有提供 nonce
		ln := len(buf)
		if ln <= 12 {
			return nil, errors.New("Invalid token data")
		}
		data = buf[:ln-12]
		nonce = buf[ln-12:]
	} else {
		data = buf
	}

	return DecryptBytes(data, key, nonce)
}

// randomNonce 新建 Nonce
func randomNonce() ([]byte, error) {
	var err error
	nonce, err := hex.DecodeString(base.MD5s(base.UUID(), "AES-Nonce", time.Now().Format(base.TFDatetimeMilli))[:24])
	// NONCE = 12
	if err != nil {
		return nil, err
	}
	return nonce, nil
}
