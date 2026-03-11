package database

import (
	"fmt"
	"os"
	"todo-api-golang/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// VARIABLE GLOBAL DB
var DB *gorm.DB

func ConnectDatabase() {
	// LOAD ENV
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// CONFIG
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// SET CONFIG TO DATABASE
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// TRY CONNECT DATABASE
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// IF ERROR -> DISCONNECT AND THROW ERROR
	if err != nil {
		panic("failed to connect database")
	}

	// DATABASE CONNECT
	DB = database

	// AUTO MIGRATION
	DB.AutoMigrate(&models.Todo{})
}
