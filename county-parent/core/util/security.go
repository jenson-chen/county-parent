package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptHex 使用MD5加密 并将其转化为十六进制
func EncryptHex(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
