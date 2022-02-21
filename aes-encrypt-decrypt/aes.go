package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// plaintext is data to be encrypted
	// plaintext := "Companion API Encryption"

	// key or encryption key is used to encrypt data
	// 16 bytes for 128 bit, 32 bytes for 256 bit
	sessionKey := make([]byte, 32)

	_, err := rand.Read(sessionKey)
	if err != nil {
		panic(err)
	}

	hexKey := hex.EncodeToString(sessionKey)
	base64key := base64.StdEncoding.EncodeToString(sessionKey)
	base64URLkey := base64.URLEncoding.EncodeToString(sessionKey)

	fmt.Println(string(sessionKey))
	fmt.Println(hexKey)
	fmt.Println(base64key)
	fmt.Println(base64URLkey)

	rsaKeyString := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnel6eJl6Z85SicWrMbMF6amXz3NjbmF7VllNBGl44HuLGrAw8puVQQbF/cRzKfijaE2sNNt5CbkgbXC4wlgis6AAw4ZwrmBIhAh7UYdvT7vd+wrap+W3QQJtRJo56VDHZ1rwm9SyjNojlKTY+I61TEfipSO/5ko7X4goD2xSy6XO7VWCqvR5jpjKZhg9aNNLW0p6k5xLBRDf2RHNFEH2QPvgtDzqfzjkKxVpruxcHJceZGhayvhs9xXuJHO4h/VH/YOSBzBaOK/Dww7THBzNBwuvr/HQcJq7KFMODA0z9k/ozvmQN8AXs4Pnyx/F7Q+FqNIkefQb1tCF/snRjg9PsQIDAQAB"
	rsaPubKey, err := base64.StdEncoding.DecodeString(rsaKeyString)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", rsaPubKey)

	block, err := aes.NewCipher(rsaPubKey)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	fmt.Println(aesGCM.NonceSize())
}
