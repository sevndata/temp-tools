package sm2

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"
)

func TestSm2GenerateKey(t *testing.T) {
	private, err := GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyHex := hex.EncodeToString(private.D.Bytes())

	publicXHex := hex.EncodeToString(private.PublicKey.X.Bytes())
	publicYHex := hex.EncodeToString(private.PublicKey.Y.Bytes())

	fmt.Printf("还原的私钥: %s\n", privateKeyHex)
	fmt.Printf("公钥 X: %s\n", publicXHex)
	fmt.Printf("公钥 Y: %s\n", publicYHex)

	privateTemp := StringToPrivateKey(privateKeyHex)
	publicTemp := StringToPublicKey(publicXHex, publicYHex)

	EncryptAndDecrypt(privateTemp, publicTemp)
}

func TestSm2(t *testing.T) {
	privateKeyHexString := "d94fb682fe5a40009e7df0e87142d380b5ec63fe7edbd651873ed74e80c0a371"
	publicXHexString := "ade411ef4a348bd6b9b5cff7736d9ca20d9f2bed9aa54f83aa5101f141aa4bc1"
	publicYHexString := "dd389c25a6fa127c3b3fb685e1ea6601454e4535a18e9c9ae1a5736508c6f272"
	privateTemp := StringToPrivateKey(privateKeyHexString)
	publicTemp := StringToPublicKey(publicXHexString, publicYHexString)

	EncryptAndDecrypt(privateTemp, publicTemp)
}

func StringToPrivateKey(privateKeyHex string) *PrivateKey {
	privaKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		fmt.Println("解析私钥失败:", err)
		return nil
	}
	private := new(PrivateKey)
	private.D = new(big.Int).SetBytes(privaKeyBytes)
	private.PublicKey.Curve = P256Sm2() // 设置曲线
	private.PublicKey.X, private.PublicKey.Y = private.PublicKey.Curve.ScalarBaseMult(private.D.Bytes())
	fmt.Printf("还原的私钥: %x\n", private.D)
	return private
}

func StringToPublicKey(publicXHex string, publicYHex string) *PublicKey {

	pubXBytes, err := hex.DecodeString(publicXHex)
	if err != nil {
		fmt.Println("解析公钥 X 失败:", err)
		return nil
	}

	pubYBytes, err := hex.DecodeString(publicYHex)
	if err != nil {
		fmt.Println("解析公钥 Y 失败:", err)
		return nil
	}

	// 创建公钥对象
	pub := new(PublicKey)
	pub.Curve = P256Sm2() // 设置曲线
	pub.X = new(big.Int).SetBytes(pubXBytes)
	pub.Y = new(big.Int).SetBytes(pubYBytes)

	// 输出验证
	fmt.Printf("还原的公钥 X: %x\n", pub.X)
	fmt.Printf("还原的公钥 Y: %x\n", pub.Y)
	return pub
}

func EncryptAndDecrypt(priv *PrivateKey, pub *PublicKey) {
	msg := []byte("{\"name\":\"sevndata\",\"name\":\"sevndata\",\"name\":\"sevndata\",\"name\":\"sevndata\",\"name\":\"sevndata\",\"name\":\"sevndata\",\"name\":\"sevndata\",\"name\":\"sevndata\",\"name\":\"sevndata\",\"name\":\"sevndata\"}")
	fmt.Printf("原文:%s\n", string(msg))
	ciphertxt, err := pub.EncryptAsn1(msg, rand.Reader) //sm2加密
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("加密结果:%x\n", ciphertxt)
	plaintxt, err := priv.DecryptAsn1(ciphertxt) //sm2解密
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("解密结果:%s\n", string(plaintxt))
	if !bytes.Equal(msg, plaintxt) {
		log.Fatal("原文不匹配")
	}
	sign, err := priv.Sign(rand.Reader, msg, nil) //sm2签名
	if err != nil {
		log.Fatal(err)
	}
	isok := pub.Verify(msg, sign) //sm2验签
	fmt.Printf("Verified: %v\n", isok)
}
