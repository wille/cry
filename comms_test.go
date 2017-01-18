package main

import (
	"fmt"
	"testing"
)

func TestComms(t *testing.T) {
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

	fmt.Println("Retrieving key...")
	priv, err = GetKey()

	if err != nil {
		panic(err)
	}

	fmt.Println(Stringify(priv))
}
