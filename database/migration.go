package database

import (
	"fmt"
	"time"

	"github.com/hpazk/go-booklib/apps/user"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// {
		// 	ID: "2020080201",
		// 	Migrate: func(tx *gorm.DB) error {
		// 		if err := tx.AutoMigrate(&user.User{}).Error; err != nil {
		// 			return err
		// 		}
		// 		if err := tx.AutoMigrate(&book.Book{}).Error; err != nil {
		// 			return err
		// 		}
		// 		return nil
		// 	},
		// Rollback: func(tx *gorm.DB) error {
		// 	if err := tx.DropTable("blogs").Error; err != nil {
		// 		return nil
		// 	}
		// 	if err := tx.DropTable("users").Error; err != nil {
		// 		return nil
		// 	}
		// 	return nil
		// },
		// },
		{
			ID: fmt.Sprintf("%d", time.Now().Unix()),
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(&user.User{}).Error; err != nil {
					return err
				}
				return nil
			},
		},
	})
}

// 1618796640
