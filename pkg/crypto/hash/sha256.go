package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// String 计算字符串的 SHA256 哈希值
func String(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// File 计算文件的 SHA256 哈希值
func File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
