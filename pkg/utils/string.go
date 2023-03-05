package utils

import (
	"github.com/samber/lo"
)

func RandomString(length int) string {
	return lo.RandomString(length, lo.LettersCharset)
}
