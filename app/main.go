package main

import (
	"github.com/smpeotn/go-debug-sample/app/extension"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello World\n")
	})
	http.HandleFunc("/foobarbaz/", func(w http.ResponseWriter, _ *http.Request) {
		extension.FooBarBaz(w)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
