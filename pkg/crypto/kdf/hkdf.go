package kdf

import (
	"hash"
	"io"

	"golang.org/x/crypto/hkdf"
)

// DeriveKey 使用 HKDF 派生密钥
// hash: 哈希函数, e.g. sha256.New
// secret: 输入的密钥材料
// salt: 可选的盐
// info: 可选的上下文信息
// keyLength: 派生密钥的长度
func DeriveKey(hash func() hash.Hash, secret, salt, info []byte, keyLength int) ([]byte, error) {
	hkdf := hkdf.New(hash, secret, salt, info)
	key := make([]byte, keyLength)
	_, err := io.ReadFull(hkdf, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}