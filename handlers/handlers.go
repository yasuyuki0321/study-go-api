package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yasuyuki0321/study-go-api.git/models"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!!\n")
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article

	err := json.NewDecoder(r.Body).Decode(&reqArticle)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()
	// クエリパラメータpageを取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	log.Println(page)

	articles := []models.Article{
		models.Article1,
		models.Article2,
	}

	json.NewEncoder(w).Encode(articles)
}

func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	}

	log.Println(articleID)

	article := models.Article1

	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	article := models.Article2

	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(article)

}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment1

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(comment)
}
