package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

const (
	// Bits Keypair bit size (higher = exponentially slower)
	Bits int = 1024
)

// Generate new RSA keypair
func Generate() *rsa.PrivateKey {
	priv, err := rsa.GenerateKey(rand.Reader, Bits)

	if err != nil {
		panic(err)
	}

	return priv
}

// Stringify private key
func Stringify(priv *rsa.PrivateKey) string {
	privateKeyDer := x509.MarshalPKCS1PrivateKey(priv)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}

	return string(pem.EncodeToMemory(&privateKeyBlock))
}

// DecodeKey
func DecodeKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(key))

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	return priv, err
}
