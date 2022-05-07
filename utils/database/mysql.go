package database

import (
	"errors"
	"fmt"
	"go-blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func generateDSN() string {
	host := config.Get("database.host")
	port := config.Get("database.port")
	database := config.Get("database.database")
	username := config.Get("database.username")
	password := config.Get("database.password")

	if host == "" || port == "" || database == "" || username == "" || password == "" {
		return ""
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database,
	)
}

func Connect() (*gorm.DB, error) {
	dsn := generateDSN()
	if dsn == "" {
		return nil, errors.New("can not connect to mySQL, dsn error")
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func Close(db *gorm.DB) {
	database, _ := db.DB()
	database.Close()
}
