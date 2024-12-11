package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ruhollahh/mini-url/config"
	"github.com/ruhollahh/mini-url/delivery/httpserver"
	"github.com/ruhollahh/mini-url/repository/mysql"
	urlrepo "github.com/ruhollahh/mini-url/repository/mysql/url"
	urlsvc "github.com/ruhollahh/mini-url/service/url"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	shortDomainURL := os.Getenv("MINIURL_URLSVC__SHORT_DOMAIN_URL")
	parsedDomainURL, err := url.Parse(shortDomainURL)
	if err != nil {
		log.Fatalln("could not parse server domain url: ", err.Error())
	}

	serverPort, err := strconv.Atoi(os.Getenv("MINIURL_SERVER__PORT"))
	if err != nil {
		log.Fatalln("could not convert server port: ", err.Error())
	}

	mysqlPort, err := strconv.Atoi(os.Getenv("MINIURL_MYSQL__PORT"))
	if err != nil {
		log.Fatalln("could not convert mysql port: ", err.Error())
	}

	cfg := config.Config{
		Mysql: mysql.Config{
			User:     os.Getenv("MINIURL_MYSQL__USER"),
			Password: os.Getenv("MINIURL_MYSQL__PASSWORD"),
			Host:     os.Getenv("MINIURL_MYSQL__HOST"),
			Port:     uint(mysqlPort),
			Database: os.Getenv("MINIURL_MYSQL__DATABASE"),
		},
		HTTPServer: httpserver.Config{
			Port: serverPort,
		},
		UrlSvc: urlsvc.Config{
			ShortDomainURL: parsedDomainURL,
			MaxRetries:     5,
			PostfixLen:     5,
		},
	}

	mysqlDB, err := mysql.New(cfg.Mysql)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func() {
		cErr := mysqlDB.Close()
		if cErr != nil {
			log.Fatalln(cErr.Error())
		}
	}()

	urlRepo := urlrepo.New(mysqlDB)

	urlSvc := urlsvc.New(cfg.UrlSvc, urlRepo)

	httpServer := httpserver.New(cfg.HTTPServer, urlSvc)
	log.Println("starting server", "port", cfg.HTTPServer.Port)
	go httpServer.Serve()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	<-sig
	fmt.Println("Shutting down gracefully...")

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go httpServer.Shutdown(wg)

	wg.Wait()
}
