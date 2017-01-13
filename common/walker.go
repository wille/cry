package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Walk recursively walks the input directory and applies all rules (extensions, limits etc)
func Walk(startPath string, callback func(filePath string, fileInfo os.FileInfo)) {
	var count int
	filepath.Walk(startPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("error", err.Error())
			return nil
		}

		var proceed bool
		for _, v := range Extensions {
			if strings.HasSuffix(filePath, v) {
				proceed = true
				break
			}
		}

		for _, v := range IgnoreDirs {
			if strings.Contains(filepath.Dir(filePath), v) {
				proceed = false
				break
			}
		}

		if proceed && count < ProcessMax {
			count++

			callback(filePath, fileInfo)
		}

		return nil
	})
}
