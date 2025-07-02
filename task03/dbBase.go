package task03

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	//db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	dsn := "root:baoqi0411@tcp(192.168.1.20:3306)/test_3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
