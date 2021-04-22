package database

import (
	"errors"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/hpazk/go-booklib/apps/event"
	"github.com/hpazk/go-booklib/apps/transaction"
	"github.com/hpazk/go-booklib/apps/user"
	"gorm.io/gorm"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// ID: fmt.Sprintf("%d", time.Now().Unix()),
			ID: "1618983623",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(&user.User{}).Error; err != nil {
					return errors.New("Migration failed")
				}
				if err := tx.AutoMigrate(&event.Event{}).Error; err != nil {
					return errors.New("Migration failed")
				}
				if err := tx.AutoMigrate(&transaction.Transaction{}).Error; err != nil {
					return errors.New("Migration failed")
				}
				return nil
			},
		},
	})
}

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
