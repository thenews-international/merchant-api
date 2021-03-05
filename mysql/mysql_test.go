package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	c "merchant/config"
	"merchant/util/logutil"
)

var DB *gorm.DB

func init() {
	var err error
	if DB, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		sqlDB, err := DB.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v", err)
		}
	}
}

func OpenTestConnection() (db *gorm.DB, err error) {
	logger, cleanup, err := logutil.NewLogger("*", "light-console", "stdout")
	if err != nil {
		panic(err)
	}
	defer cleanup()

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	var cfg c.Config

	if err := viper.ReadInConfig(); err != nil {
		logger.Warn(fmt.Sprintf("Error reading config file, %s", err))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		logger.Warn(fmt.Sprintf("Unable to decode into struct, %v", err))
	}

	db, err = New(&cfg.Database, cfg.Debug)
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	return
}
