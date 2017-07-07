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
	database.username = gonv.GetStringEnv("DBUSER", "root")
	database.password = gonv.GetStringEnv("DBPASSWORD", "")
	database.host = gonv.GetStringEnv("DBHOST", "localhost")
	database.port = gonv.GetIntEnv("DBPORT", 3306)
	database.database = gonv.GetStringEnv("DB", "goweb")
}

func (this *databaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", this.username, this.password, this.host, this.port, this.database)
}

func GetUrlDatabase() string {
	return database.url()
}
