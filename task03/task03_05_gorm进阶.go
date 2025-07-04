package task03

import (
	"GoTask/task03/models"
	"gorm.io/gorm"
	"strconv"
)

func InitTable02(db *gorm.DB) {
	//_ = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	_ = db.AutoMigrate(&models.Post{})
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
		{Title: "张三的博客5", Content: "张三的博客内容", User: user},
		{Title: "张三的博客6", Content: "张三的博客内容02", User: user},
		{Title: "张三的博客7", Content: "张三的博客内容03", User: user},
		{Title: "张三的博客8", Content: "张三的博客内容04", User: user},
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

func GetMaxCommentCountPost(db *gorm.DB) models.Post {
	//var maxCountPost models.Post
	//var sql = "SELECT    p.id,    p.title,   p.content,    COUNT(c.id) AS commentsCount " +
	//	"FROM posts p " +
	//	"LEFT JOIN    comments c ON p.id = c.post_id AND c.deleted_at IS NULL " +
	//	"WHERE    p.deleted_at IS NULL " +
	//	"GROUP BY     p.id " +
	//	"HAVING     commentsCount >= " +
	//	"All (        SELECT          COUNT(c2.id)     " +
	//	" FROM             posts p2         " +
	//	"LEFT JOIN            comments c2 ON p2.id = c2.post_id AND c2.deleted_at IS NULL" +
	//	" WHERE             p2.deleted_at IS NULL       GROUP BY             p2.id    )"
	//db.Raw(sql).Scan(&maxCountPost)
	var maxCountPost2 models.Post
	//maxCommentSubQuery := db.Table("comments").
	maxCommentSubQuery := db.Model(&models.Comment{}).
		Select("COUNT(comments.id)").
		Joins("JOIN posts ON posts.id = comments.post_id AND posts.deleted_at IS NULL").
		Where("comments.deleted_at IS NULL").
		Group("posts.id")

	db.Model(&models.Post{}).
		Select("posts.id, posts.title, posts.content, COUNT(comments.id) AS commentsCount").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id AND comments.deleted_at IS NULL").
		Where("posts.deleted_at IS NULL").
		Group("posts.id").
		Having("COUNT(comments.id) >=ALL  (?)", maxCommentSubQuery).
		Scan(&maxCountPost2)
	return maxCountPost2
}

func DeleteComment(db *gorm.DB, ids []int) {
	//tx := db.Begin()
	for _, id := range ids {
		var comment models.Comment
		db.Where("id = ?", id).Find(&comment)
		db.Delete(&comment)
	}
	//tx.Commit()
}
