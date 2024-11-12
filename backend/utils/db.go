package utils

import (
    "GeoDataApp/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func InitDB() {
    dsn := "host=localhost user=postgres password=Pranav@5jan dbname=geodataapp port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB = db

    // Auto-migrate models
    db.AutoMigrate(&models.Credentials{})
}
