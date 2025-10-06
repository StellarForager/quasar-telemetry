package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	fmt.Println("Private key:", base64.StdEncoding.EncodeToString(priv))
	fmt.Println("Public key:", base64.StdEncoding.EncodeToString(pub))
	data, _ := os.ReadFile("config.json")
	sig := ed25519.Sign(ed25519.PrivateKey(priv), data)
	fmt.Println("Signature:", base64.StdEncoding.EncodeToString(sig))
}
