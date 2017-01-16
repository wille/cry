package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

const Endpoint = "http://localhost/upload"

func TestWeb(t *testing.T) {
	priv := Generate()
	key := Stringify(priv)

	_, err := http.PostForm(Endpoint, url.Values{"k": {key}})

	if err != nil {
		panic(err)
	}

	fmt.Println("Key uploaded")
}
