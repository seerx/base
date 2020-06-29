package tokencrypto

import (
	"encoding/hex"
	"time"

	"github.com/seerx/base"
)

const (
	sign   = "20200629"
	host   = "xval.cn"
	answer = "It's right!"
)

var key []byte = []byte(base.MD5s(sign, host, answer))

// SetAESKey 设置秘钥,建议使用长度为 32 byte 数组
func SetAESKey(aesKey []byte) {
	key = aesKey
}

// newNonce 新建 Nonce
func newNonce() ([]byte, error) {
	var err error
	nonce, err := hex.DecodeString(base.MD5s(base.UUID(), host, time.Now().Format(base.TFDatetimeMilli))[:24])
	// NONCE = 12
	if err != nil {
		return nil, err
	}
	return nonce, nil
}
