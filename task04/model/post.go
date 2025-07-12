package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    sql.NullString
	Content  sql.NullString
	User     *User `json:",omitempty"` // 当 User == nil 时不显示
	UserId   uint
	Comments []*Comment
}
