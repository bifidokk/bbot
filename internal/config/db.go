package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(GetEnv("db_dsn")), &gorm.Config{})

	if err != nil || db == nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
