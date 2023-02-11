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

	http.HandleFunc("/hello", helloHandler)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
