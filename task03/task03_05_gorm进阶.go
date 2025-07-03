package task03

import (
	"GoTask/task03/models"
	"gorm.io/gorm"
)

func InitTable02(db *gorm.DB) {
	_ = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
