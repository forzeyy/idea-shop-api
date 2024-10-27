package database

import (
	"log"
	"os"

	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("sdelivery.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database. error:\n", err.Error())
		os.Exit(1)
	}

	log.Println("successfully connected to database!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations...")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.Category{})
	log.Println("done!")
}
