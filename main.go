package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"aki/api"
	"aki/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/sys/unix"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	version = "?"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("reading config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(conf.PgDSN), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		log.Fatalf("opening db conection: %v", err)
	}

	e := echo.New()
	e.HidePort = true
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RemoveTrailingSlash(),
	)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     conf.AllowOrigins,
		AllowMethods:     []string{echo.GET, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	apiHandler := api.New(db)

	privateGroup := e.Group(
		conf.PrivatePrefix(),
		apiHandler.AuthMiddleware(db, conf.LoginURL),
	)
	fillGroup(privateGroup, apiHandler.PrivateURLs())

	publicGroup := e.Group(conf.PublicPrefix())
	fillGroup(publicGroup, apiHandler.PublicURLs())

	txGroup := e.Group(
		conf.PrivatePrefix(),
		apiHandler.AuthMiddleware(db, conf.LoginURL),
		DBTransactionMiddleware(db),
	)

	fillGroup(txGroup, apiHandler.TxURLs())

	s := startServer(e, conf.HTTPAddr())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, unix.SIGINT, unix.SIGQUIT, unix.SIGTERM)
	for {
		<-sig

		s.Shutdown(context.Background())
	}
}

func startServer(e *echo.Echo, addr string) *http.Server {
	s := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	go func() {
		log.Printf("starting server on %s... version: %s", s.Addr, version)
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	return s
}

func fillGroup(g *echo.Group, urls []api.ApiURL) {
	for _, u := range urls {
		g.Add(u.Method, u.Path, u.Handler)
	}
}
