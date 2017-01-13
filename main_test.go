package main

import (
	"testing"
)

func TestCrypto(t *testing.T) {
	file := "test.docx"
	priv := Generate()

	encrypt(file, priv)
	decrypt(file+LockedExtension, priv)
}
