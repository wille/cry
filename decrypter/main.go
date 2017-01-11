package main

import (
	"fmt"
	"os"

	"github.com/redpois0n/cry/common"
)

func main() {
	startWalk := common.GetHomeDir()

	common.Walk(startWalk, func(filePath string, fileInfo os.FileInfo) {
		fmt.Println("decrypting", filePath)
	})
}
