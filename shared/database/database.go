package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=Desyncq password=Desyncq dbname=desyncqdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to database")

	return db, nil
}


