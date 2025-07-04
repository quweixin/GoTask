package main

import (
	"GoTask/task03"
	"GoTask/task03/models"
)

//func main() {
//	db := task03.GetDb()
//	task03.CreateTable(db)
//
//	//创建数据
//	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
//	for i := 1; i <= 10; i++ {
//		age := rand.Intn(11) + 20
//		student := &models.Student{Name: "LiSi" + strconv.Itoa(i), Age: age, Grade: "四年级", Sex: 1}
//		student = task03.CreateStudent(db, student)
//		fmt.Printf("创建成功: %+v\n", student.Name)
//	}
//	//查询数据
//	students := task03.GetStudent(db, 20)
//	for _, student := range students {
//		fmt.Println(student.Name)
//	}
//
//	//更新数据
//	task03.UpdateStudent(db, &models.Student{Name: "张三", Grade: "四年级"})
//
//	//删除数据
//	task03.DeleteStudent(db, 25)
//}

func main() {
	db := task03.GetDb()
	//task03.InitTable(db)
	//task03.InitData(db)

	// 测试转账
	//task03.Transfer(db, 2, 1, decimal.NewFromFloat(100))

	// gorm 进阶 ，表创建
	//task03.InitTable02(db)
	//初始化 用户
	//task03.CreateUser(db)

	var user models.User
	db.Debug().Where("name = ?", "张三").Find(&user)
	task03.CreatePost(db, user)

	//var posts []models.Post
	//db.Where("id  in ?", []uint{9, 10, 11}).Find(&posts)
	//task03.CreateComment(db, posts)

	//使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//var posts []models.Post
	//posts = task03.GetUserPostAndComments(db, user)
	//for _, post := range posts {
	//	fmt.Println(post)
	//}

	//	编写Go代码，使用Gorm查询评论数量最多的文章信息。
	//var postMax models.Post
	//postMax = task03.GetMaxCommentCountPost(db)
	//fmt.Println(postMax.ID, postMax.Title, postMax.CommentsCount)

}
