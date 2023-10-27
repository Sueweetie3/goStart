package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		lissajous(writer)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
