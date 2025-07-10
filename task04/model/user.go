package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:user_name;unique;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"-"`
	Email    string `gorm:"column:email;unique;not null" json:"email"`
	Posts    []Post
}
