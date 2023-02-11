package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World!!\n")
		fmt.Printf("%v\n", r.Header.Get("foo"))
	}

	postArticleHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
