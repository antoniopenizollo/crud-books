package database

import (
	"fmt"
	"github.com/antoniopenizollo/crud-books/configs"
	"github.com/antoniopenizollo/crud-books/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB() {

	conf := configs.GetDB()

	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Database, conf.Pass)

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatalln("error", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDataBase() *gorm.DB {
	return db
}
