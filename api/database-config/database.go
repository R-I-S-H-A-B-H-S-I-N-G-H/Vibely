package databaseconfig

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct{}

var MY_SQL_DB *gorm.DB

func (d *Database) Init() (*gorm.DB, error) {
	dsn := "root:lemein..23@tcp(127.0.0.1:3306)/vibely?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	MY_SQL_DB = db
	return db, err
}

func GetMySqlDB() *gorm.DB {
	return MY_SQL_DB
}
