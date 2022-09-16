package repositories

import (
	"database/sql"

	"github.com/HT0323/go_api/models"
)

// 引数で受け取ったComment構造体の情報でレコードを作成
func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)

	var newComment models.Comment
	newComment.ArticleID = comment.ArticleID
	newComment.Message = comment.Message

	if err != nil {
		return models.Comment{}, err
	}
	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}
