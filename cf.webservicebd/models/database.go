package models

import "database/sql"
import "fmt"
import _ "github.com/go-sql-driver/mysql" //Probar bd
import "log"

var db *sql.DB

const username string = "root"
const password string = ""
const host string = "localhost"
const port int = 3306
const database string = "goweb"

func CreateConnection() {
	if connection, err := sql.Open("mysql", generateURL()); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

func CreateTables() {
	createTable("user", userSchema)
}

func createTable(tableName string, schema string) {
	if !existTable(tableName) {
		_, err := db.Exec(schema)
		if err != nil {
			log.Println(err)
		}
	}
}

func existTable(tableName string) bool {
	sql := fmt.Sprintf("Show Tables Like '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}
	return rows.Next()
}

func CloseConnection() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, database)
}
