package main

import (
	"os"
)

func getHomeDir() string {
	home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")

	if home == "" {
		home = os.Getenv("USERPROFILE")
	}

	return home
}
