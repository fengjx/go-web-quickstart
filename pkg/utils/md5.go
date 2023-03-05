package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5SumString(input string) string {
	hash := md5.New()
	return fmt.Sprintf("%x", hash.Sum([]byte(input)))
}
