package logic

import (
	"crypto/md5"
	"fmt"
)

// MD5Encrypt MD5加密
func MD5Encrypt(inStr string) (outStr string) {

	data := []byte(inStr + "_")
	md5 := md5.New()
	hash := md5.Sum(data)
	outStr = fmt.Sprintf("%x", hash)
	return
}
