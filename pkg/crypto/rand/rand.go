package rand

import (
	"crypto/rand"
	"math/big"
)

// String 生成一个指定长度的随机字符串
func String(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// Int 生成一个在 [min, max] 范围内的随机整数
func Int(min, max int) (int, error) {
	if min > max {
		min, max = max, min
	}

	result, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}

	return int(result.Int64()) + min, nil
}
