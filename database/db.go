package database

import (
	"fmt"
	"log"

	"github.com/kanugi/rest-api-assignment2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "user123"
	dbname   = "postgres"
	port     = 5432
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		host, port, user, dbname, password)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	fmt.Println("Successfully connected to database")

	db.AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
