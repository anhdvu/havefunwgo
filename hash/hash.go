package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	plain := "hash this one for me please. Let's try a longer string so that I know if there is a cut off if a string is way too long."

	h := sha256.New()
	_, err := h.Write([]byte(plain))
	if err != nil {
		fmt.Println(err)
	}

	hashed1 := h.Sum(nil)
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(hashed1))

	hashed2 := sha256.Sum256([]byte(plain))
	hashed3 := make([]byte, 32)
	copy(hashed3, hashed2[:])
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString([]byte(hashed3)))
}
