package controllers

import (
	"net/http"

	"go-blog/database"
	"go-blog/models"
	"go-blog/tools"

	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var comment models.Comment
	user_id, _ := tools.ExtractUserID(ctx)
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse comment"})
	}
	comment.UserID = user_id
	if err := database.AddComment(&comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add comment"})
	}

	ctx.JSON(http.StatusOK, gin.H{"comment": comment})
}
