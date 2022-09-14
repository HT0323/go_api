package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/HT0323/go_api/models"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// queryMap := req.URL.Query()

	// var page int
	// if p, ok := queryMap["page"]; ok && len(p) > 0 {
	// 	var err error
	// 	// pageパラメータに複数値が入ってる場合は初めのものを採用する
	// 	page, err = strconv.Atoi(p[0])
	// 	// クエリパラが数字以外であればエラーを出力
	// 	if err != nil {
	// 		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	// 		return
	// 	}
	// } else {
	// 	// クエリパラが付与されていない場合は、page=1が付与されているとみなす
	// 	page = 1
	// }

	article1 := models.Article1
	article2 := models.Article2
	articleList := []models.Article{article1, article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	// articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	// if err != nil {
	// 	http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	// 	return
	// }
	// resString := fmt.Sprintf("Article No.%d\n", articleId)
	// io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article Comment...\n")
}
