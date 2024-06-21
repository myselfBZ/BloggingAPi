package models

import (
	"github.com/myselfBZ/BloggingAPI/pkg/config"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Blogs    []Blog `json:"blogs"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Likes    []Like `json:"likes"`
}

func (u *User) Create(user *User) ( error) {
	result := config.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) Delete(id uint) error {
	result := config.DB.Delete(u, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) Update(username string, id uint) (*User, error) {
	result := config.DB.First(u, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result := config.DB.Model(u).Updates(&User{
		Username: username,
	}); result.Error != nil {
		return nil, result.Error
	}

	u.Username = username
	return u, nil

}

func (u *User) Get(username string) (*User, error) {
	var user User
	if result := config.DB.Where("username = ?", username).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
