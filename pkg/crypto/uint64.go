package crypto

import "encoding/binary"

// EncryptUInt64 加密 uint64
func EncryptUInt64(val uint64, key []byte, nonce []byte, appendNonce bool) (string, error) {
	var buf = make([]byte, 8)
	// binary.BigEndian
	binary.BigEndian.PutUint64(buf, val)
	return EncyptBytesToString(buf, key, nonce, appendNonce)
}

// EncryptUInt64Random 加密 uint64，使用随机 nonce
func EncryptUInt64Random(val uint64, key []byte) (string, error) {
	var buf = make([]byte, 8)
	// binary.BigEndian
	binary.BigEndian.PutUint64(buf, val)
	return EncyptBytesToStringRandom(buf, key)
}

// DecryptUInt64 解密 uint64
func DecryptUInt64(val string, key []byte, nonce []byte) (*uint64, error) {
	data, err := DecryptBytesFromString(val, key, nonce)
	if err != nil {
		return nil, err
	}
	n := binary.BigEndian.Uint64(data)
	return &n, nil
}
