package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var Env struct {
	ApiUrl string `json:"API_URL"`
	Host   string `json:"HOST"`
	Port   string `json:"PORT"`
	// MYSQL
	MysqlUsername string `json:"MYSQL_USERNAME"`
	MysqlPassword string `json:"MYSQL_PASSWORD"`
	MysqlHost     string `json:"MYSQL_HOST"`
	MysqlPort     string `json:"MYSQL_PORT"`
	MysqlDbName   string `json:"MYSQL_DBNAME"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		Env.ApiUrl = os.Getenv("API_URL")
		Env.Host = os.Getenv("HOST")
		Env.Port = os.Getenv("PORT")

		Env.MysqlUsername = os.Getenv("MYSQL_USERNAME")
		Env.MysqlPassword = os.Getenv("MYSQL_PASSWORD")
		Env.MysqlHost = os.Getenv("MYSQL_HOST")
		Env.MysqlPort = os.Getenv("MYSQL_PORT")
		Env.MysqlDbName = os.Getenv("MYSQL_DBNAME")
	} else {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		rootApp := strings.TrimSuffix(path, "/config")
		os.Setenv("APP_PATH", rootApp)

		var myEnv map[string]string
		myEnv, err = godotenv.Read()
		if err != nil {
			panic(err)
		}

		b, _ := json.Marshal(myEnv)
		json.Unmarshal(b, &Env)
	}

}
