package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Posts     []Post
	PostCount int64 `gorm:"default:0"`
}
