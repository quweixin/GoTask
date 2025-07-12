package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content sql.NullString
	Post    *Post
	PostId  uint
	User    *User
	UserId  uint
}
