package helperlib

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

type app struct {
	db *Store
	c  *appConfig
	e  *echo.Echo
}

type appConfig struct {
	Addr   string `json:"addr"`
	DBURL  string `json:"dbUrl"`
	DBUser string `json:"dbUser"`
	DBPass string `json:"dbPass"`
	DBName string `json:"dbName"`
}

func newApp(config *appConfig) *app {
	if config == nil {
		return nil
	}

	config.Addr = ":8080"

	s := &Store{
		Address:  config.DBURL,
		Database: config.DBName,
		Password: config.DBPass,
		Username: config.DBUser,
	}

	return &app{
		c:  config,
		db: s,
		e:  echo.New(),
	}
}

func (a *app) initApp() error {
	log.Println("init start")

	if err := a.db.Connect(); err != nil {
		return err
	}

	a.e.HideBanner = true
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())
	a.e.Use(middleware.CORS())

	// a.e.Static("/images", "images")

	// a.e.POST("/login", a.login)
	// a.e.POST("/register", a.register)

	// api := a.e.Group("/api")

	// api.Use(middleware.JWT([]byte(a.c.JwtSecret)))

	// product := api.Group("/product")

	// Start server
	go func() {
		if err := a.e.Start(a.c.Addr); err != nil {
			a.e.Logger.Info("shutting down the server")
		}
	}()

	log.Println("init complete")

	return nil
}

func (a *app) closeApp() {
	log.Println("Closing echo server")
	if a.e != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		if err := a.e.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		cancel()
	}
	log.Println("Closed echo server")

	log.Println("Closing db connection")
	if a.db != nil {
		a.db.Close()
	}

	log.Println("Closed db connection")

}
