package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"go-blog/database"
	"go-blog/models"
	"go-blog/session"
	"go-blog/tools"

	"github.com/gin-gonic/gin"
)

type UserRegisterRequest struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	FisrtName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
}

type UserLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func UserRegister(ctx *gin.Context) {
	var user UserRegisterRequest
	if ctx.BindJSON(&user) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
		return
	}
	if !database.CheckUserExists(user.Username) {
		new_user := models.User{
			Username:  user.Username,
			Email:     user.Email,
			FisrtName: user.FisrtName,
			LastName:  user.LastName,
		}
		new_user.SetPassword(user.Password)
		database.CreateUser(&new_user)
		ctx.JSON(http.StatusOK, gin.H{"status": "user created"})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"status": "user already exists"})
	}
}

func CheckUsername(ctx *gin.Context) {
	var username string
	if ctx.BindJSON(&username) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
	} else if database.CheckUserExists(username) {
		ctx.JSON(http.StatusConflict, gin.H{"status": "username has already taken"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "username available"})
	}
}

func UserLogin(ctx *gin.Context) {
	// Decode the body of request
	var user UserLoginRequest
	if ctx.BindJSON(&user) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "invalid request"})
		return
	}

	// Check if user exists
	db_user, err := database.GetUserByUsername(user.Username)
	if err != nil {
		if db_user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error getting user"})
		return
	}

	// Check if password is correct
	if db_user.ComparePasswords(user.Password) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "wrong password"})
		return
	}

	// Store username in the session
	sn := session.Create(db_user.ID)

	// Generate access token and refresh token
	access_token, _ := tools.GenerateToken(strconv.Itoa(int(db_user.ID)), time.Hour*1, os.Getenv("JWT_SECRET"))
	refresh_token, _ := tools.GenerateToken(sn.SessionID, time.Hour*24*7, os.Getenv("REFRESH_TOKEN_SECRET"))
	ctx.SetCookie("access_token", access_token, 3600, "/", "", false, true)
	ctx.SetCookie("refresh_token", refresh_token, 3600*24*7, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "login success", "session": sn})
}

func UserLogout(ctx *gin.Context) {
	ctx.SetCookie("refresh_token", "", -1, "/", "", false, true)
	ctx.SetCookie("access_token", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"status": "logout success"})
}

func UserID(ctx *gin.Context) {
	value, _ := ctx.Get("user_id")
	ctx.JSON(200, gin.H{"user_id": value})
}
