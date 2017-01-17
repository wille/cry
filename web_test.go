package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

const Endpoint = "http://localhost:1312/upload"

func TestWeb(t *testing.T) {
	priv := Generate()
	key := Stringify(priv)

	fmt.Println(key)

	_, err := http.PostForm(Endpoint, url.Values{"k": {key}})

	if err != nil {
		panic(err)
	}

	fmt.Println("Key uploaded")
}
