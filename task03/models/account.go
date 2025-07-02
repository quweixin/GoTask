package models

import "github.com/shopspring/decimal"

type Account struct {
	Id      uint
	Balance decimal.Decimal `gorm:"type:decimal(20,8)"`
}
