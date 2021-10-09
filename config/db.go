package config

import (
	"fmt"
	"test-start/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var e error

func InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		GetEnvVariable("DB_HOST"),
		GetEnvVariable("DB_USERNAME"),
		GetEnvVariable("DB_PASSWORD"),
		GetEnvVariable("DB_NAME"),
		GetEnvVariable("DB_PORT"))

	DB, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	// // nyalakan komentar ini untuk auto migrate
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(
		&migrations.Students{})
}
