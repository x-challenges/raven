package gorm

import (
	"gorm.io/gorm"
)

func NewDriver(dialector gorm.Dialector) (*gorm.DB, error) {
	return gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
}
