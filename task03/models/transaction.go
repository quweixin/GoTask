package models

import "github.com/shopspring/decimal"

type Transaction struct {
	Id            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        decimal.Decimal `gorm:"type:decimal(20,8)"`
}
