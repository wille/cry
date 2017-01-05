// +build !windows

package main

import (
	"os"
)

func getHomeDir() string {
	return os.Getenv("HOME")
}
