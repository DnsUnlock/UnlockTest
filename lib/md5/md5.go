package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Sum(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
