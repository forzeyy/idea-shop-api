package database

import (
	"log"
	"os"

	"github.com/forzeyy/idea-shop-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_FILENAME")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database. error:\n", err.Error())
		os.Exit(1)
	}
	log.Println("successfully connected to database!")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations...")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.CartItem{}, &models.Category{}, &models.Order{}, &models.Comment{}, &models.Admin{})
	log.Println("done!")

	return db
}
