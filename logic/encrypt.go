package logic

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Encrypt MD5加密
func MD5Encrypt(inStr string) (outStr string) {

	m5 := md5.New()
	_, err := m5.Write([]byte(inStr))
	if err != nil {
		panic(err)
	}
	outStr = hex.EncodeToString(m5.Sum(nil))
	return
}
