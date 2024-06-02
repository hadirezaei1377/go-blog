package database

import (
	"go-blog/models"
)

func CheckUserExists(username string) bool {
	results := DB.Take(&models.User{}, "username = ?", username)
	return results.RowsAffected > 0
}

func CreateUser(user *models.User) (uint, error) {
	err := DB.Create(&user).Error
	return user.ID, err
}
