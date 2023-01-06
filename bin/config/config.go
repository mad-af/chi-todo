package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var Env struct {
	Host string `json:"HOST"`
	Port string `json:"PORT"`
	// MYSQL
	MysqlUsername string `json:"MYSQL_USERNAME"`
	MysqlPassword string `json:"MYSQL_PASSWORD"`
	MysqlHost     string `json:"MYSQL_HOST"`
	MysqlPort     string `json:"MYSQL_PORT"`
	MysqlDbName   string `json:"MYSQL_DBNAME"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

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
