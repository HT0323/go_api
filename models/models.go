package models

import "time"

type Comment struct {
	CommentID int       `json:"comment_id"` //コメントのID
	ArticleID int       `json:"article_id"` //コメントを投稿する記事のID
	Message   string    `json:"message"`    // コメント内容
	CreatedAT time.Time `json:"created_at"` // 投稿日時
}

type Article struct {
	ID          int       `json:"article_id"` //記事のID
	Title       string    `json:"title"`      //記事のタイトル
	Contents    string    `json:"contents"`   //記事の内容
	UserName    string    `json:"user_name"`  //投稿者名
	NiceNum     int       `json:"nice"`       //記事へのいいね数
	CommentList []Comment `json:"comments"`   //記事へついたコメントの一覧
	CreatedAT   time.Time `json:"created_at"` // 投稿日時
}
