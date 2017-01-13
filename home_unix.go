// +build !windows

package main

import (
	"os"
)

func GetHomeDir() string {
	return os.Getenv("HOME")
}
