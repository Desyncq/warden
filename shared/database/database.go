package database

import (
	"github.com/Desyncq/warden/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func Connect() (*gorm.DB, error) {
	var err error

	once.Do(func() {

		dsn := "host=localhost user=Desyncq password=Desyncq dbname=desyncqdb port=5432 sslmode=disable"
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database: ", err)
		} else {
			log.Println("Connected to database")
		}

		// Create Tables
		dbInstance.AutoMigrate(&models.Admin{})
	})

	return dbInstance, err
}
