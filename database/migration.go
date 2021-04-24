package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/hpazk/go-booklib/database/model"
	"gorm.io/gorm"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1619281958",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.Transaction{})
			},
		},
	})
}
