package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HT0323/go_api/apperrors"
	"github.com/HT0323/go_api/controllers/services"
	"github.com/HT0323/go_api/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonを構造体にデコード
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.RrqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	// 構造体をjsonにエンコード
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
