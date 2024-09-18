package mysql

import (
	"github.com/kafka-embracetheday/goResourceWatcher/config"
	"github.com/kafka-embracetheday/goResourceWatcher/internal/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitMysql() {
	cfg := config.GetConfig()
	db, err = gorm.Open(mysql.Open(cfg.Dsn()), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		logger.Logger.Panicf("failed to connect database: %v", err)
		return
	}
	logger.Logger.Infof("connected to mysql database")

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(2 * time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	cfg := config.GetConfig()
	db, err = gorm.Open(mysql.Open(cfg.Dsn()), &gorm.Config{})
	if err != nil {
		logger.Logger.Panicf("failed to connect database: %v", err)
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(2 * time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	return db
}
