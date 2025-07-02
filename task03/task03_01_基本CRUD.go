package task03

import (
	"GoTask/task03/models"
	"gorm.io/gorm"
)

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
