package util

import (
	"crypto/sha256"
	"fmt"
)

// 生成md5 32位字符串
func Md5String32(content []byte) string {
	encodedBytes := sha256.Sum256(content)
	return fmt.Sprintf("%v", encodedBytes)
}
