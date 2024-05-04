package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mjekrami/issue-tracker/config"
)

var cfg config.EnvConfig

func main() {
	// Read .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// Populate cfg variable with exposed environment variable
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)))
	//TODO gracefully shutdown the server
}
