package repositories

import (
	"database/sql"

	"github.com/HT0323/go_api/models"
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
