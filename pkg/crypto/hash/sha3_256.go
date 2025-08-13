package hash

import (
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/sha3"
)

// StringSHA3_256 计算字符串的 SHA3-256 哈希值
func StringSHA3_256(s string) string {
	h := sha3.New256()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// FileSHA3_256 计算文件的 SHA3-256 哈希值
func FileSHA3_256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha3.New256()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}