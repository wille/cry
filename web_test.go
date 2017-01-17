package main

import (
	"fmt"
	"testing"
)

func TestWeb(t *testing.T) {
	fmt.Println("Generating key...")

	priv := Generate()
	str := Stringify(priv)
	fmt.Println(str)

	fmt.Println("Uploading...")
	err := PostKey(priv)

	if err != nil {
		panic(err)
	}

	fmt.Println("Key uploaded")
}
