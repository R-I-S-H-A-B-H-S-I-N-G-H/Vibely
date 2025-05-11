package databaseconfig

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct{}

var MY_SQL_DB *gorm.DB

func (d *Database) Init() (*gorm.DB, error) {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	MY_SQL_DB = db
	return db, err
}

func GetMySqlDB() *gorm.DB {
	return MY_SQL_DB
}
