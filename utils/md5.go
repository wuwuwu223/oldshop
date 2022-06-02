package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Compare this snippet from utils/md5.go:
// package utils
//md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
