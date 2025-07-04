package models

import (
	"database/sql"
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
	CommentStatus sql.NullString
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("id = ?", p.User.ID).Update("post_count", gorm.Expr("post_count + ?", 1))
	return err
}
