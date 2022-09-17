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

// 記事についているコメント一覧を取得
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// sqlの取得結果を構造体に保存
	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
