package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	model "graphql/cmd/app/domain/dao"
)

type DBinstance struct {
	DB *gorm.DB
}

var DB *DBinstance

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	log.Printf("PASSWORD: %s", os.Getenv("DB_PASSWORD"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connected")
	db.Logger = db.Logger.LogMode(logger.Info)

	log.Println("running migration")
	db.AutoMigrate(&model.User{}, &model.Post{})
	log.Println("migration completed")

	DB = &DBinstance{DB: db}
}
