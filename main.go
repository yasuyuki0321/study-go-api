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

	articleListHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}

	articleDetailHandler := func(w http.ResponseWriter, r *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No %d\n", articleID)
		io.WriteString(w, resString)
	}

	postNiceHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	postCommentHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
