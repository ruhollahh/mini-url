package httpserver

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ruhollahh/mini-url/delivery/httpserver/urlhandler"
	urlsvc "github.com/ruhollahh/mini-url/service/url"
	"sync"
	"time"
)

type Config struct {
	Port int
}

type Server struct {
	cfg        Config
	router     *echo.Echo
	urlHandler *urlhandler.Handler
}

func New(cfg Config, urlSvc *urlsvc.Service) Server {
	return Server{
		cfg:        cfg,
		router:     echo.New(),
		urlHandler: urlhandler.New(urlSvc),
	}
}

func (s Server) Serve() {
	s.registerRoutes()

	s.router.Logger.Fatal(s.router.Start(fmt.Sprintf(":%d", s.cfg.Port)))
}

func (s Server) registerRoutes() {
	router := s.router
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())

	router.GET("/", home)

	router.POST("/urls/create", s.urlHandler.Create)
	router.GET("/:postfix", s.urlHandler.Visit)
}

func (h Server) Shutdown(wg *sync.WaitGroup) {
	defer wg.Done()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := h.router.Shutdown(ctx)
	if err != nil {
		fmt.Println("could not properly shutdown http server: ", err)
	}
}
