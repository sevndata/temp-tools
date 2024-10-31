package aes

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {

	key, iv, err := GenerateKeyAndIV(32)
	if err != nil {
		fmt.Println("Error generating key and IV:", err)
		return
	}

	keyBase64 := base64.StdEncoding.EncodeToString(key)
	ivBase64 := base64.StdEncoding.EncodeToString(iv)
	fmt.Println("keyBase64 (Base64):", keyBase64)
	fmt.Println("ivBase64 (Hex):", ivBase64)

	plainText := []byte("Hello, this is a secret message!")

	encrypted, err := Encrypt(plainText, key, iv)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)
	fmt.Println("Encrypted (Base64):", encryptedBase64)

	decrypted, err := Decrypt(encrypted, key, iv)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("Decrypted:", string(decrypted))
}
