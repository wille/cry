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

type Pair struct {
	Id  string
	Key string
}

var Keys = []Pair{}

func main() {
	http.HandleFunc(UploadRoute, handleUpload)
	http.HandleFunc(RetrievalRoute, handleRetrieve)

	log.Fatal(http.ListenAndServe(":1312", nil))

}

func reject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rejecting ", r.RemoteAddr)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, http.StatusText(http.StatusNotFound))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("i")
	key := r.PostFormValue("k")

	if r.Method != "POST" || key == "" || id == "" {
		reject(w, r)
		return
	}

	fmt.Println(key)

	pair := Pair{Id: id, Key: key}
	Keys = append(Keys, pair)
}

func handleRetrieve(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("i")

	if r.Method != "POST" || id == "" {
		reject(w, r)
		return
	}

	for _, pair := range Keys {
		if pair.Id == id {
			fmt.Fprint(w, pair.Key)
			return
		}
	}

	reject(w, r)
}
