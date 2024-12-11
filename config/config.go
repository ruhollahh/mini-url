package config

import (
	"github.com/ruhollahh/mini-url/delivery/httpserver"
	"github.com/ruhollahh/mini-url/repository/mysql"
	urlsvc "github.com/ruhollahh/mini-url/service/url"
)

type Config struct {
	Mysql      mysql.Config
	UrlSvc     urlsvc.Config
	HTTPServer httpserver.Config
}
