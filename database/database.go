package database

import {
	"fmt"
	"log"
	"os"

	"github.com/leonardobaranelli/rest-api-backend-app/models"
	"gorm.io/dirver/postgres"
	"gorm.io/gorm"	
	"gorm.io/gorm/logger"
}

type DBinstance struct {
	Db gorm.DB
}

var DB DBinstance

func ConnectDb() {
	dns := fmt.Sprintf{
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shangai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	}
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.info),
	})

	if err != nil {
		log.Fatal("Falied to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.Fact())

	DB = DBinstance{
		Db: db,
	}
}
