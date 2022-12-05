package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"origomicrosservices.com/consumer-contratos/domain/models"
	"os"
)

var (
	DB *gorm.DB
)

func Connect() {
	dsn := os.Getenv("DB_USER") +
		":" +
		os.Getenv("DB_PASS") +
		"@tcp(" +
		os.Getenv("DB_HOST") +
		":" +
		os.Getenv("DB_PORT") +
		")/" +
		os.Getenv("DB_NAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic(err)
	}

	// Create from map

	fmt.Println("Executing migrations ...")
	db.AutoMigrate(&models.ContractJSON{})
	fmt.Println("Sucess")

	DB = db
}
