package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// Address server listening address
	Address = ":1312"

	// UploadRoute key uploading path
	UploadRoute = "/upload"

	// RetrievalRoute
	RetrievalRoute = "/retrieve"
)

// Pair private key and computer id
type Pair struct {
	Id  string
	Key string
}

// Keys stored in memory
var Keys = []Pair{}

func main() {
	http.HandleFunc(UploadRoute, handleUpload)
	http.HandleFunc(RetrievalRoute, handleRetrieve)

	fmt.Println("Listening on", Address)
	log.Fatal(http.ListenAndServe(Address, nil))
}

func reject(w http.ResponseWriter, r *http.Request, reason string) {
	fmt.Println("Rejecting ", r.RemoteAddr+":", reason)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, http.StatusText(http.StatusNotFound))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("i")
	key := r.PostFormValue("k")

	if r.Method != "POST" {
		reject(w, r, "HTTP method is not POST, got "+r.Method)
		return
	}

	if id == "" {
		reject(w, r, "id parameter i not set or empty")
		return
	}

	if key == "" {
		reject(w, r, "key parameter k is not set or empty")
	}

	for _, pair := range Keys {
		if pair.Id == id {
			reject(w, r, "key already exists")
			return
		}
	}

	pair := Pair{Id: id, Key: key}
	Keys = append(Keys, pair)
}

func handleRetrieve(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("i")

	if r.Method != "POST" {
		reject(w, r, "HTTP method is not POST, got "+r.Method)
		return
	}

	if id == "" {
		reject(w, r, "id parameter i is not set")
		return
	}

	for _, pair := range Keys {
		if pair.Id == id {
			fmt.Fprint(w, pair.Key)
			return
		}
	}

	reject(w, r, "no key found for id "+id)
}
