package config

import (
	"fmt"

	"github.com/adafatya/micro-services/inventory-service/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dbName := util.GetEnv("DB_NAME", "")
	dbHost := util.GetEnv("DB_HOST", "")
	dbPort := util.GetEnv("DB_PORT", "")
	dbUser := util.GetEnv("DB_USER", "")
	dbPass := util.GetEnv("DB_PASS", "")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
