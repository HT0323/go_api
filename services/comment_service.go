package services

import (
	"github.com/HT0323/go_api/apperrors"
	"github.com/HT0323/go_api/models"
	"github.com/HT0323/go_api/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
