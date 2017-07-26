package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

type databaseConfig struct {
	username string
	password string
	host     string
	port     int
	database string
}

var database *databaseConfig

func init() {
	database = &databaseConfig{}
	database.username = gonv.GetStringEnv("DBUSER", "nrk")
	database.password = gonv.GetStringEnv("DBPASSWORD", "nrk227")
	database.host = gonv.GetStringEnv("DBHOST", "13.59.36.111")
	database.port = gonv.GetIntEnv("DBPORT", 3306)
	database.database = gonv.GetStringEnv("DB", "restaurante")
}

func (this *databaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", this.username, this.password, this.host, this.port, this.database)
}

func GetUrlDatabase() string {
	return database.url()
}
