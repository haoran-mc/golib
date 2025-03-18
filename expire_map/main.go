package main

import (
	"fmt"
	"time"
)

var mp = make(map[string]struct{})

func SetCache(key string) {
	mp[key] = struct{}{}

	expireTime := 5 * time.Second
	time.AfterFunc(expireTime, func() {
		delete(mp, key)
	})
}

func main() {
	SetCache("hello")

	for i := range 10 {
		_, found := mp["hello"]
		if found {
			fmt.Printf("%ds found\n", i)
		} else {
			fmt.Printf("%ds not found\n", i)
		}
		time.Sleep(1 * time.Second)
	}
}
