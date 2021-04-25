package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/hpazk/go-ticketing/database/model"
	"gorm.io/gorm"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1619317987",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.User{})
			},
		},
	})
}
