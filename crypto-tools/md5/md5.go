package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encryption(srcString string) string {
	hasher := md5.New()
	hasher.Write([]byte(srcString))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
