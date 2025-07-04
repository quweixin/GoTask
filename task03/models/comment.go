package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string
	PostId  uint
	Post    Post
}

//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var postId = c.PostId
	var commentsCount int
	tx.Model(&Comment{}).
		Where("post_id = ? ", postId).
		Select("COUNT(id)").
		Scan(&commentsCount)
	fmt.Println(commentsCount)
	if commentsCount == 0 {
		tx.Model(&Post{}).Where("id=?", postId).Update("comment_status", "无评论")
	}
	return err
}
