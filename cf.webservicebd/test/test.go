package main

import "../models"

func main() {
	models.CreateConnection()
	models.CreateTables()
	//models.Ping()
	models.CloseConnection()
}
