package main

import (
	"GoTask/task03"
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
	task03.InitTable02(db)
}
