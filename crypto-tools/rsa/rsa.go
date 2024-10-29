package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func private_key() []byte {
	return []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQD36x3neV/vlUE7tcLcwpSlnErSuKu6fljo7mA2/8GBp18imSXW
5FtoQZAOQ6wbNgYlFdXJ1zvKNWBxykv6dfI8yeLWlzsKs4DlIkJ/Q3gM9s3K9y5q
JHsycPnu7CQTccSMvg0mt2i2kVSOTFkesCF8ZvN/g6beA3akvIAjJmDLHQIDAQAB
AoGACDV2xgVU0eAFSu7XxuuCdIFaUOPge8pXM09BTFDnnV1nXwPnJthk1mzlUjOX
bUc6qNuyOn6v8iAWU245Wc+x95SgZqKDIR+hSUn65BP9jEYqDkSLvtEIbUhD7lhu
hd251Xra9Ql8fz3s8+BlYS+w8OHhWwZg4p9kx7Xm0bD4FwECQQD/RfYViQhExS6a
b7jhrBIw457voTlGHz0Lx183hLGQPxNzf+hcFBJSslyAsRTW+GpVt2ekMn9WR0GX
a+KF9/OhAkEA+J/LosFjjVuTObRUyeR8k56eZEa3+Rw3YM6vR4OBd9qB8Jrw6iHT
2wq9yZ7iyoRQQoTktEmLtYN9QRgTWabl/QJBAJMIRPmaeDzEJlelyiOR7NhkAwOb
JVYSYCMMrvdXYr7m1dlM7xU6KmOfgFBPruUzKGw+u1+EBnw3hrwk09C+RIECQHyF
zEAgC/RowTxNjYnw3lQxRbODl/E0qKfb3P48Z0PYNOyFLdSWTL2Qi63H3l8AFhK6
7LE/hLHMwZcwr8BfTyECQCrEvWWz+gD2sP47peii1kmvXwY/0d1V3+tR2qtbwXRT
qMHSZ9FZ6DijwJ2NHQJakI6YJbfSfTigju/Y1VwAAXw=
-----END RSA PRIVATE KEY-----
`)
}

func public_key() []byte {
	return []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD36x3neV/vlUE7tcLcwpSlnErS
uKu6fljo7mA2/8GBp18imSXW5FtoQZAOQ6wbNgYlFdXJ1zvKNWBxykv6dfI8yeLW
lzsKs4DlIkJ/Q3gM9s3K9y5qJHsycPnu7CQTccSMvg0mt2i2kVSOTFkesCF8ZvN/
g6beA3akvIAjJmDLHQIDAQAB
-----END PUBLIC KEY-----
`)
}
func GenerateKey(bits int) (*rsa.PrivateKey, rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey
	return privateKey, publicKey

}
func EncryptOAEP(text string) (string, error) {
	rsaPublicKey, err := ParsePKIXPublicKey(public_key())
	if err != nil {
		return "", err
	}
	secretMessage := []byte(text)
	rng := rand.Reader
	cipherdata, err := rsa.EncryptOAEP(sha256.New(), rng, rsaPublicKey, secretMessage, nil)
	if err != nil {
		return "", nil
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
	return ciphertext, nil
}

func DecryptOAEP(ciphertext string) (string, error) {
	rsaPrivateKey, err := ParsePKCS1PrivateKey(private_key())
	if err != nil {
		return "", err
	}

	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, rsaPrivateKey, cipherdata, nil)
	if err != nil {
		return "", nil
	}

	return string(plaintext), nil
}

func ParsePKIXPublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubInterface.(*rsa.PublicKey), nil
}

func ParsePKCS1PrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	privateInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateInterface, nil
}
