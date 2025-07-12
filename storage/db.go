package storage

import (
	"auth-service/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(dsn string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to db:", err)
    }

    if err := db.AutoMigrate(&models.User{}); err != nil {
        log.Fatal("auto migration failed:", err)
    }

    return db
}