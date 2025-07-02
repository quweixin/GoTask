package task03

import (
	"GoTask/task03/models"
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

func CreateTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Student{})
	if err != nil {
		panic("failed to create table")
	}
}

func CreateStudent(db *gorm.DB, student *models.Student) *models.Student {
	result := db.Create(student)
	if result.Error != nil {
		return nil
	}
	return student
}

func GetStudent(db *gorm.DB, age int) []models.Student {
	var students []models.Student
	db.Debug().Where("age >= ?", age).Find(&students)
	return students
}

func UpdateStudent(db *gorm.DB, student *models.Student) *models.Student {
	result := db.Debug().Model(&models.Student{}).Where("name = ?", student.Name).Update("grade", student.Grade)
	if result.Error != nil {
		return nil
	}
	return student
}

func DeleteStudent(db *gorm.DB, age int) {
	db.Debug().Where("age >= ?", age).Delete(&models.Student{})
}
