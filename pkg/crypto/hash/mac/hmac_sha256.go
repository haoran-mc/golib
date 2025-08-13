package mac

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
)

// Sign 使用 HMAC-SHA256 计算消息的 MAC
func Sign(message, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return hex.EncodeToString(mac.Sum(nil))
}

// Verify 使用 HMAC-SHA256 验证消息的 MAC
func Verify(message, signature, key []byte) bool {
	expectedMAC := Sign(message, key)
	return subtle.ConstantTimeCompare([]byte(signature), []byte(expectedMAC)) == 1
}