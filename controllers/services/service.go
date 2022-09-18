package services

import "github.com/HT0323/go_api/models"

type MyAppServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
	PostCommentService(comment models.Comment) (models.Comment, error)
}
