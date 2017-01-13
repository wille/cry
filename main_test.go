package main

import (
	"testing"

	"github.com/redpois0n/cry/common"
)

func TestCrypto(t *testing.T) {
	file := "test.docx"
	priv := common.Generate()

	encrypt(file, priv)
	decrypt(file+common.LockedExtension, priv)
}
