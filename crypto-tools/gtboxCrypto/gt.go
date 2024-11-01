package gtboxCrypto

import "github.com/george012/gtbox/gtbox_encryption"

func GtEncryption(srcString, keyString string) string {
	return gtbox_encryption.GTEncryptionGo(srcString, keyString)
}

func GtDecryption(srcString, key string) string {
	return gtbox_encryption.GTDecryptionGo(srcString, key)
}
