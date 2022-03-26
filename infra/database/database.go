package database

import (
	"auth-api/infra/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost user=postgres dbname=auth sslmode=disable password=docker port=5432"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error", err)
	} else {
		log.Println("Database connected")
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetGatabase() *gorm.DB {
	return db
}
