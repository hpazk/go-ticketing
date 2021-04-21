package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func DbConfig() DbConnection {
	dbConfig := DbConnection{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_DATABASE"),
		User:     os.Getenv("DB_USER"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	return dbConfig
}

func GetDbInstance() *gorm.DB {
	dbConfig := DbConfig()

	// dsn := "host=localhost user=postgres password= dbname=jersey_dev2 port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Database,
		dbConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connection failed")
	}

	return db
}
