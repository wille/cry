package main

import (
	"os"
)

func GetHomeDir() string {
	home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")

	if home == "" {
		home = os.Getenv("USERPROFILE")
	}

	return home
}
