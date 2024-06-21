package models

import (
	"github.com/myselfBZ/BloggingAPI/pkg/config"
)

type Like struct {
	UserID uint `gorm:"uniqueIndex:idx_user_blog" json:"user_id"`
	ID     uint `gorm:"primaryKey" json:"id"`
	BlogID uint `gorm:"uniqueIndex:idx_user_blog" json:"blog_id"`
}

func (l *Like) Like(userId, blogId uint) error {
	result := config.DB.Create(&Like{UserID: userId, BlogID: blogId})
	if result.Error != nil {
		config.DB.Where("user_id = ? AND blog_id = ?", userId, blogId).Delete(&Like{})
		return result.Error
	}
	return nil
}
