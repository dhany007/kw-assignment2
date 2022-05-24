package databases

import (
	"assignment2/models"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = ""
	DB_NAME = "postgres"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = migrate(db)
	if err != nil {
		panic(err)
	}

	log.Default().Println("connection db success")

	return db
}

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Order{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Item{}); err != nil {
		return err
	}
	return nil
}
