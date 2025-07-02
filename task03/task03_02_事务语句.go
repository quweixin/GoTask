package task03

import (
	"GoTask/task03/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Account{}, &models.Transaction{})
	if err != nil {
		return
	}
}

func InitData(db *gorm.DB) {
	// 创建一个 account 实例
	balanceA, _ := decimal.NewFromString("200.50")
	accountA := models.Account{
		Balance: balanceA,
	}
	db.Create(&accountA)
	balanceB, _ := decimal.NewFromString("100")
	accountB := models.Account{
		Balance: balanceB,
	}
	db.Create(&accountB)
}
