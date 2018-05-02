package main

import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, UNH!")
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":80", nil))
}
