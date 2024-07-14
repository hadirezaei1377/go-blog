package database

import "go=blog/models"

func AddComment(comment *models.Comment) error {
	return DB.Create(comment).Error
}

func GetComment(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := DB.Where("id = ?", id).First(&comment).Error
	return &comment, err
}
