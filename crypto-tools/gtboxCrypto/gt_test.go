package gtboxCrypto

import (
	"fmt"
	"testing"
)

func TestGt(t *testing.T) {
	keyString := "snlongpark"
	srcString := "00391-80000-00001-AA304Windows 10 Pro for Workstations{AA9E919A-B650-4F8C-BD3C-7C35141D0F83}"
	ciphertext := GtEncryption(srcString, keyString)
	fmt.Println("Plaintext:", ciphertext)
	plaintext := GtDecryption(ciphertext, keyString)
	fmt.Println("Plaintext:", plaintext)
}
