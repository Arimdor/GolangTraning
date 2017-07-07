package orm

import (
	"../config"
	_ "github.com/go-sql-driver/mysql" //Mysql Driver
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func CreateConnection() {
	url := config.GetUrlDatabase()
	if connection, err := gorm.Open("mysql", url); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

func CloseConnection() {
	db.Close()
}

func CreateTables() {
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})
}

func DropTables() {
	db.DropTableIfExists(&User{})
}

func Ping() error {
	return db.DB().Ping()
}
