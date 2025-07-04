package task03

import (
	"GoTask/task03/models"
	"gorm.io/gorm"
	"strconv"
)

func InitTable02(db *gorm.DB) {
	_ = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}

func CreateUser(db *gorm.DB) {
	users := []*models.User{
		{Name: "张三"},
		{Name: "李四"},
		{Name: "王五"},
		{Name: "赵十"},
		{Name: "钱一"},
	}
	db.Create(users)
}

func CreatePost(db *gorm.DB, user models.User) {

	posts := []*models.Post{
		{Title: "张三的博客", Content: "张三的博客内容", User: user},
		{Title: "张三的博客02", Content: "张三的博客内容02", User: user},
		{Title: "张三的博客03", Content: "张三的博客内容03", User: user},
		{Title: "张三的博客04", Content: "张三的博客内容04", User: user},
	}
	db.Create(posts)
}

func CreateComment(db *gorm.DB, post []models.Post) {
	comments := make([]models.Comment, 0, 10)
	for i := 0; i < len(post); i++ {
		var post = post[i]
		if i == 1 {
			for i := 0; i < 5; i++ {
				comments = append(comments, models.Comment{Content: "评论" + strconv.Itoa(i), Post: post})
			}
		} else {
			for i := 0; i < 2; i++ {
				comments = append(comments, models.Comment{Content: "评论" + strconv.Itoa(i), Post: post})

			}
		}
	}

	db.Create(comments)
}

func GetUserPostAndComments(db *gorm.DB, user models.User) []models.Post {
	var posts []models.Post
	db.Preload("Comments").Where("user_id = ?", user.ID).Find(&posts)
	return posts
}
