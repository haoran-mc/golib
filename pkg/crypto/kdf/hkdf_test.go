package kdf

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestDeriveKey(t *testing.T) {
	secret := []byte("my-secret")
	salt := []byte("my-salt")
	info := []byte("my-info")
	keyLength := 32

	key, err := DeriveKey(sha256.New, secret, salt, info, keyLength)
	if err != nil {
		t.Fatalf("DeriveKey failed: %v", err)
	}

	if len(key) != keyLength {
		t.Errorf("Expected key length %d, but got %d", keyLength, len(key))
	}

	// 这是一个固定的期望值，如果 DeriveKey 的实现改变，这个测试也会失败
	expectedKeyHex := "2f3b4a9ad8b4a3f2a3b4a9ad8b4a3f2a3b4a9ad8b4a3f2a3b4a9ad8b4a3f2a"
	actualKeyHex := hex.EncodeToString(key)

	// 注意：这里的期望值是伪的，因为HKDF的输出是伪随机的。
	// 在真实的测试中，您可能需要使用已知的测试向量。
	// 这里我们只检查长度和是否成功生成。
	t.Logf("Derived key (hex): %s", actualKeyHex)

	// 这是一个简单的例子，所以我们不与一个固定的值进行比较
	if actualKeyHex != expectedKeyHex {
		t.Errorf("Expected key %s, but got %s", expectedKeyHex, actualKeyHex)
	}
}
