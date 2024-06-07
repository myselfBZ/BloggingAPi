package models

import (
	"time"

	"github.com/myselfBZ/BloggingAPI/pkg/config"
)

type Blog struct {
	UserID    uint      `gorm:"not null"`
	ID        uint      `gorm:"primaryKey;not null" json:"id"`
	Content   string    `json:"content"`
	Title     string    `gorm:"type:varchar(50)" json:"title"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Likes     []Like    `json:"likes"`
}

func (b *Blog) Create() error {
	result := config.DB.Create(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Blog) Delete(userID, id uint) error {
	result := config.DB.Delete(&Blog{}, userID, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *Blog) GetAll() ([]Blog, error) {
	var blogs []Blog
	result := config.DB.Preload("Likes").Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

func (b *Blog) Update(id uint, newB *Blog, userId uint) (*Blog, error) {
	var oldBlog Blog
	result := config.DB.Where("user_id = ? AND id = ?", userId, id).First(&oldBlog)
	if result.Error != nil {
		return nil, result.Error
	}
	if result := config.DB.Model(&oldBlog).Updates(&Blog{
		Content: newB.Content,
		Title:   newB.Title,
	}); result.Error != nil {
		return nil, result.Error
	}
	b.Content = newB.Content
	b.Title = newB.Title
	return b, nil

}

func (b Blog) GetBlog(id uint) (*Blog, error) {
	result := config.DB.Preload("Likes").First(&b, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &b, nil
}
