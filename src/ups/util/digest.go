package util

import (
	"crypto/sha256"
	"fmt"
)

func Md5(content string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(content)))
}
