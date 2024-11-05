package md5

import (
	"fmt"
	"strings"
	"testing"
)

func TestGt(t *testing.T) {
	srcString := "xxx"
	ciphertext := strings.ToUpper(MD5Encryption(srcString))
	fmt.Println("Plaintext:", ciphertext)
}
