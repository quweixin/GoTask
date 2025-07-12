package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    sql.NullString
	Content  sql.NullString
	User     *User
	UserId   uint
	Comments []*Comment
}
