package task03

import (
	"GoTask/task03/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func Transfer(db *gorm.DB, fromAccountId uint, toAccountId uint, amount decimal.Decimal) {
	// 开启事务
	tx := db.Begin()
	// 获取 fromAccount
	var fromAccount models.Account
	//tx.First(&fromAccount, fromAccountId)
	tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&fromAccount, fromAccountId)
	if fromAccount.Balance.LessThan(amount) {
		tx.Rollback()
		return
	}
	fromAccount.Balance = fromAccount.Balance.Sub(amount)
	tx.Save(&fromAccount)
	// 获取 toAccount
	//var toAccount models.Account
	//tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&toAccount, toAccountId)
	//toAccount.Balance = toAccount.Balance.Add(amount)
	//tx.Save(&toAccount)

	tx.Exec("update accounts set balance = balance + ? where id = ?", amount, toAccountId)
	// 创建一个 transaction 实例
	transaction := models.Transaction{
		FromAccountId: fromAccountId,
		ToAccountId:   toAccountId,
		Amount:        amount,
	}
	tx.Create(&transaction)

	tx.Commit()
}
