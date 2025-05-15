package models

import (
	"strconv"

	"github.com/mikietechie/gopeople/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func PgConnect() {
	logLevel, err := strconv.Atoi(config.DB_LOGGING_LEVEL)
	if err != nil {
		panic(err)
	}
	Db, err = gorm.Open(postgres.Open(config.DATABASE_CONNECTION), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
	})
	if err != nil {
		panic(err)
	}

	err = Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
				CREATE TYPE gender_enum AS ENUM ('male', 'female', 'other');
			END IF;
		END$$;
		`).Error; err != nil {
			return err
		}
		if err := tx.AutoMigrate(
			&User{},
		); err != nil {
			return err
		}
		return nil
	})
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
