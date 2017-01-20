package main

import (
	"crypto/rsa"
	"io"
	"net/http"
	"net/url"
)

// PostKey sends the private key to the remote serrver
func PostKey(priv *rsa.PrivateKey, id string) error {
	key := Stringify(priv)

	_, err := http.PostForm(UploadEndpoint, url.Values{
		"k": {key},
		"i": {id},
	})

	return err
}

func GetKey(id string) (*rsa.PrivateKey, error) {
	req, err := http.PostForm(RetrieveEndpoint, url.Values{
		"i": {id},
	})

	if err != nil {
		return nil, err
	}

	key := make([]byte, req.ContentLength)

	io.ReadFull(req.Body, key)

	priv, err := DecodeKey(key)

	return priv, err
}
