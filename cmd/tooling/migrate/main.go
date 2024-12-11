package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ruhollahh/mini-url/repository/mysql"
	"github.com/ruhollahh/mini-url/repository/mysql/migrator"
	"log"
	"os"
	"strconv"
)

func main() {
	mysqlPort, err := strconv.Atoi(os.Getenv("MINIURL_MYSQL__PORT"))
	if err != nil {
		log.Fatalln("could not convert mysql port: ", err.Error())
	}

	config := mysql.Config{
		User:     os.Getenv("MINIURL_MYSQL__USER"),
		Password: os.Getenv("MINIURL_MYSQL__PASSWORD"),
		Host:     os.Getenv("MINIURL_MYSQL__HOST"),
		Port:     uint(mysqlPort),
		Database: os.Getenv("MINIURL_MYSQL__DATABASE"),
	}

	err = migrator.Migrate(config)
	if err != nil {
		log.Fatalln("could not migrate mysql: ", err.Error())
	}
}
