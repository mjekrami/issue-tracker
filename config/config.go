package config

type EnvConfig struct {
	AppHost          string `envconfig:"APP_HOST" required:"true" default:"localhost"`
	AppPort          string `envconfig:"APP_PORT" required:"true" default:"8083"`
	Environment      string `envconfig:"ENVIRONMENT" required:"true"`
	DatabaseHost     string `envconfig:"DATABASE_HOST" required:"true"`
	DatabasePort     string `envconfig:"DATABASE_PORT" required:"true"`
	DatabaseUsername string `envconfig:"DATABASE_USERNAME" required:"true"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" required:"true"`
}
