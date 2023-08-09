package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// This used to connecting to the database.
func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=admin dbname=go-jwt-mux port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	db.AutoMigrate(&User{})

	DB = db
}
