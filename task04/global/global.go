package global

import (
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	Trans ut.Translator
)
