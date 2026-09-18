package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	ENV  *Config
	once sync.Once
)

type Config struct {
	App struct {
		Name        string `env:"APP_NAME" env-default:"skeleton"`
		PortHTTP    string `env:"APP_PORT_HTTP" env-default:"8181"`
		PortGRPC    string `env:"APP_PORT_GRPC" env-default:"50051"`
		Host        string `env:"APP_HOST" env-default:"127.0.0.1"`
		Environment string `env:"APP_ENV" env-default:"development"`
	}

	DB struct {
		Host              string `env:"DB_HOST" env-default:"localhost"`
		Port              int    `env:"DB_PORT" env-default:"5432"`
		Username          string `env:"DB_USERNAME" env-default:"skeleton"`
		Password          string `env:"DB_PASSWORD" env-default:"skeleton"`
		Database          string `env:"DB_NAME" env-default:"skeleton_db"`
		SSLMode           string `env:"DB_SSL_MODE" env-default:"disable"`
		ConnectionTimeout int    `env:"DB_CONN_TIMEOUT" env-default:"30" env-description:"database timeout in seconds"`
		MaxOpenCons       int    `env:"DB_MAX_OPEN_CONS" env-default:"20" env-description:"database max open conn in seconds"`
		MaxIdleCons       int    `env:"DB_MAX_IDLE_CONS" env-default:"20" env-description:"database max idle conn in seconds"`
		ConnMaxLifetime   int    `env:"DB_CONN_MAX_LIFETIME" env-default:"0" env-description:"database conn max lifetime in seconds"`
	}
}

// Load loads configuration from optional env file path or environment variables.
func Load(path ...string) *Config {
	once.Do(func() {
		ENV = &Config{}
		filePath := ".env"
		if len(path) > 0 && path[0] != "" {
			filePath = path[0]
		}
		if err := cleanenv.ReadConfig(filePath, ENV); err != nil {
			_ = cleanenv.ReadEnv(ENV)
		}
	})
	return ENV
}
