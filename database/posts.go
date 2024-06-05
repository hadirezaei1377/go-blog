package database

import "go-blog/models"

func CreatePost(post *models.Post) (uint, error) {
	err := DB.Create(&post).Error
	return post.ID, err
}
