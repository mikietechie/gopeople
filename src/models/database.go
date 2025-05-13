package models

import (
	"github.com/mikietechie/gopeople/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func PgConnect() {
	var err error
	Db, err = gorm.Open(postgres.Open(config.DATABASE_CONNECTION), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}
	err = Db.AutoMigrate(
		&User{},
	)
	if err != nil {
		panic(err)
	}
}

func PgDisconnect() {
	db, err := Db.DB()
	if err != nil {
		panic(err)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
}
