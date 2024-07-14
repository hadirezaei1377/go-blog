package controllers

import (
	"net/http"

	"go-blog/models"
	"go-blog/tools"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	db databases.Database
}

func NewCommentController(db databases.Database) *commentController {
	return &commentController{
		db: db,
	}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	var comment models.Comment
	user_id, _ := tools.ExtractUserID(ctx)
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse comment"})
	}
	comment.UserID = user_id
	if err := cc.db.AddComment(&comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add comment"})
	}

	ctx.JSON(http.StatusOK, gin.H{"comment": comment})
}
