package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!!\n")
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Listing Articles...\n")
}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No %d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
