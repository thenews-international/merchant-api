package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"merchant/config"
)

func New(conf *config.DatabaseConfig, isDebugModeOn bool) (*gorm.DB, error) {
	dsn := conf.ToDsnString()

	var logLevel logger.LogLevel
	if isDebugModeOn {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
}
