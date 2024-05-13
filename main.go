package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mjekrami/issue-tracker/cmd/auth"
	"github.com/mjekrami/issue-tracker/cmd/database"
	"github.com/mjekrami/issue-tracker/cmd/routes"
	"github.com/mjekrami/issue-tracker/config"
)

var cfg config.EnvConfig

func main() {
	// Read .env file
	log.Info("loading environment variables")
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// Populate cfg variable with exposed environment variable
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}
	log.Info(fmt.Sprintf("connecting to the database: %s:%s", cfg.DatabaseHost, cfg.DatabasePort))
	if err = database.New(cfg.DatabaseHost, cfg.DatabasePort, auth.DBAuth{Username: cfg.DatabaseUsername, Password: cfg.DatabasePassword}); err != nil {
		log.Error(cfg.DatabaseHost, cfg.DatabasePort)
		panic(err)
	}
	log.Info("successfully conncted to the database")
	e := echo.New()
	// Register routes
	routes.HealthCheckRoute(e)
	routes.UsersRoute(e)
	routes.IssuesRoute(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)))
}
