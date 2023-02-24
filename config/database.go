package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	USER := GetInv("MYSQL_USER")
	PASS := GetInv("MYSQL_PASSWORD")
	HOST := GetInv("MYSQL_HOST")
	PORT := GetInv("MYSQL_PORT")
	DBNAME := GetInv("MYSQL_DBNAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func GetInv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
