package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/mjekrami/issue-tracker/cmd/auth"
	database "github.com/mjekrami/issue-tracker/cmd/database"
	"github.com/mjekrami/issue-tracker/cmd/routes"
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
	if err = database.New(cfg.DatabaseHost, cfg.DatabasePort, auth.DBAuth{Username: cfg.DatabaseUsername, Password: cfg.DatabasePassword}); err != nil {
		panic(err)
	}
	e := echo.New()
	// Add all routes
	routes.HealthCheck(e)
	routes.UsersRoute(e)
	routes.IssuesRoute(e)
	//	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)))
}
