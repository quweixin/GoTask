package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title         string
	Content       string
	Comments      []Comment
	CommentsCount int64 `gorm:"column:commentsCount"`
	UserId        uint
	User          User
}
