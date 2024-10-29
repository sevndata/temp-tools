package rsa

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaEncryptOrDecrypt(t *testing.T) {
	ciphertext, err := EncryptOAEP("Hello, this is a secret message!")
	if err != nil {
		panic(err)
	}
	fmt.Println("Ciphertext:", ciphertext)
	encryptedBase64 := base64.StdEncoding.EncodeToString([]byte(ciphertext))
	fmt.Println("Ciphertext (Base64):", encryptedBase64)
	plaintext, err := DecryptOAEP(ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Println("Plaintext:", string(plaintext))
}
func TestGenerateKey(t *testing.T) {
	GenerateKey(2048)
}
