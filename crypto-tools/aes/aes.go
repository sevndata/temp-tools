package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func GenerateKeyAndIV(keySize int) ([]byte, []byte, error) {
	key := make([]byte, keySize)
	iv := make([]byte, aes.BlockSize)

	if _, err := rand.Read(key); err != nil {
		return nil, nil, err
	}

	if _, err := rand.Read(iv); err != nil {
		return nil, nil, err
	}

	return key, iv, nil
}

func Encrypt(plainText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plainText = pkcs7Padding(plainText, block.BlockSize())

	mode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	mode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

func Decrypt(cipherText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	return pkcs7UnPadding(plainText)
}

func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(padding)
	}
	return append(src, padText...)
}

func pkcs7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("invalid padding")
	}
	return src[:(length - unpadding)], nil
}
