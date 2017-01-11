package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Walk recursively walks the input directory and applies all rules (extensions, limits etc)
func Walk(startPath string, callback func(filePath string, fileInfo os.FileInfo)) {
	filepath.Walk(startPath, func(filepath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("error", err.Error())
			return nil
		}

		var proceed bool
		for _, v := range Extensions {
			if strings.HasSuffix(filepath, v) {
				proceed = true
				break
			}
		}

		if proceed {
			callback(filepath, fileInfo)
		}

		return nil
	})
}
