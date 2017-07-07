package models

import "database/sql"
import "fmt"
import "log"

import "../config"
import _ "github.com/go-sql-driver/mysql" //MySQL Driver
var db *sql.DB

func CreateConnection() {
	url := config.GetUrlDatabase()
	fmt.Println(url)
	if connection, err := sql.Open("mysql", url); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

func CreateTables() {
	createTable("users", userSchema)
}

func createTable(tableName string, schema string) {
	if !existTable(tableName) {
		log.Println("Existe")
		Exec(schema)
	} else {
		truncateTable(tableName)
	}
}

func truncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

func existTable(tableName string) bool {
	sql := fmt.Sprintf("Show Tables Like '%s'", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return rows, err
}

func CloseConnection() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
