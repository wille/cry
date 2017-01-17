package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", handleUpload)

	log.Fatal(http.ListenAndServe(":1312", nil))

}

func reject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rejecting ", r.RemoteAddr)
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, http.StatusText(http.StatusNotFound))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	key := r.PostFormValue("k")

	if r.Method != "POST" || key == "" {
		reject(w, r)
		return
	}

	fmt.Println(key)
}
