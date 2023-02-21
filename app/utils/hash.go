package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetShortHash(text string, length int) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])[:length]
}
