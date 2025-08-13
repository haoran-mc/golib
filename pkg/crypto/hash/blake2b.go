package hash

import (
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/blake2b"
)

// StringBLAKE2b 计算字符串的 BLAKE2b-256 哈希值
func StringBLAKE2b(s string) string {
	h, _ := blake2b.New256(nil)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// FileBLAKE2b 计算文件的 BLAKE2b-256 哈希值
func FileBLAKE2b(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h, _ := blake2b.New256(nil)
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
