package config

import (
	"www.ivtlinfoview.com/infotax/infotax-backend/pkg/log"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB function used to make new mysql db connection
func NewDB(conf *Config) *gorm.DB {

	log.Logger.Println("Connecting to mysql database")
	db, err := gorm.Open(mysql.Open(conf.DatabaseSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Logger.Fatalf("Error, While connection mysql database: %v\n", err)
		return nil
	}

	return db

}
