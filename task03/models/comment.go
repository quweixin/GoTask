package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	PostId  uint
	Post    Post
}
