package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/segmentio/ksuid"
)

type key struct {
	public  string // used to generate api key
	private string // used to generate api secret
}

func (k key) generateKey() string {
	return ksuid.New().String()
}

func (k key) generateSecret(p string) string {
	h := hmac.New(sha256.New, []byte(k.private))
	h.Write([]byte(p))
	hashed := h.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(hashed)
}

func main() {
	cfg := key{
		private: "khongthebatmi",
	}

	public := cfg.generateKey()
	private := cfg.generateSecret(public)

	fmt.Printf("API Key: %s\n", public)
	fmt.Printf("API Secret: %s\n", private)
}
