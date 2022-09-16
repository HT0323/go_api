package repositories

import (
	"database/sql"

	"github.com/HT0323/go_api/models"
)

const (
	// 1ページあたりの記事の表示数
	articleNum = 5
)

// 引数で受け取ったArticle構造体の情報でレコードを作成
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0 now());
	`

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	var newArticle models.Article
	newArticle.Title = article.Title
	newArticle.Contents = article.Contents
	newArticle.UserName = article.UserName

	if err != nil {
		return models.Article{}, err
	}
	id, _ := result.LastInsertId()
	newArticle.ID = int(id)

	return newArticle, nil
}

// 記事一覧の取得
func SelectArticleList(db *sql.DB, page int) ([]models.Article, err) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, articleNum, (page-1)*articleNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// sqlの取得結果を構造体に保存
	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}
