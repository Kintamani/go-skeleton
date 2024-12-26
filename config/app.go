package config

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	ENV  *Config
	once sync.Once
)

type Config struct {
	App struct {
		Name  string `env:"APP_NAME" envDefault:"skeleton"`
		Port  string `env:"APP_PORT" envDefault:"8181"`
		Host  string `env:"APP_HOST" envDefault:"127.0.0.1"`
		Stage string `env:"APP_STAGE" envDefault:"development"`
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

// option adalah fungsi yang digunakan untuk mengatur konfigurasi aplikasi
type Option = func(c *Configure) error

// configure mengatur data struct
type Configure struct {
	path     string
	filename string
}

// ConfigEnv mengatur konfigurasi aplikasi
func ConfigEnv(opt ...Option) *Configure {
	c := &Configure{}

	for _, o := range opt {
		err := o(c)
		if err != nil {
			panic(err)
		}
	}
	return c
}

// initialize akan membuat konfigurasi aplikasi
func (c *Configure) Initialize() {
	logger := logrus.New()
	once.Do(func() {
		ENV = &Config{}
		if err := LoadEnv(Options{
			Config:    ENV,
			Paths:     []string{c.path},
			Filenames: []string{c.filename},
		}); err != nil {
			logger.Warn("Warning: env file does not exist")
		}
	})
}

// withPath akan menambahkan path untuk mengatur konfigurasi aplikasi
func WithPath(path string) Option {
	return func(c *Configure) error {
		c.path = path
		return nil
	}
}

// WithFilename akan menambahkan filename untuk mengatur konfigurasi aplikasi
func WithFilename(filename string) Option {
	return func(c *Configure) error {
		c.filename = filename
		return nil
	}
}
