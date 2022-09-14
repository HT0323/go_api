package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test1",
		CreatedAT: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test2",
		CreatedAT: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "test article",
		Contents:    "TEST",
		UserName:    "saki",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAT:   time.Now(),
	}

	Article2 = Article{
		ID:        2,
		Title:     "second article",
		Contents:  "TEST",
		UserName:  "yamada",
		NiceNum:   2,
		CreatedAT: time.Now(),
	}
)
