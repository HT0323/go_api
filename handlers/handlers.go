package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/HT0323/go_api/models"
	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// jsonを構造体にデコード
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// 構造体をjsonにエンコード
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		// pageパラメータに複数値が入ってる場合は初めのものを採用する
		page, err = strconv.Atoi(p[0])
		// クエリパラが数字以外であればエラーを出力
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		// クエリパラが付与されていない場合は、page=1が付与されているとみなす
		page = 1
	}
	log.Println(page)

	article1 := models.Article1
	article2 := models.Article2
	articleList := []models.Article{article1, article2}

	// 構造体をjsonにエンコード
	json.NewEncoder(w).Encode(articleList)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {

	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	log.Println(articleId)

	article := models.Article1

	// 構造体をjsonにエンコード
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// jsonを構造体にデコード
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// 構造体をjsonにエンコード
	article := reqArticle
	json.NewEncoder(w).Encode(article)
	// io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonを構造体にデコード
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// 構造体をjsonにエンコード
	comment := reqComment
	json.NewEncoder(w).Encode(comment)
	// io.WriteString(w, "Article Comment...\n")
}
