// +build !windows

package common

import (
	"os"
)

func GetHomeDir() string {
	return os.Getenv("HOME")
}
