package database

import (
	"fmt"
	"log"

	"github.com/jaksonkallio/coinbake/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var handle *gorm.DB

func Connect() error {
	log.Printf("Connecting to database at %s", config.CurrentConfig.Database.Host)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.CurrentConfig.Database.Username,
		config.CurrentConfig.Database.Password,
		config.CurrentConfig.Database.Host,
		config.CurrentConfig.Database.Name,
	)
	newHandle, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	handle = newHandle

	return nil
}

func Handle() *gorm.DB {
	return handle
}
