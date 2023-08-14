package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

type (
	Config struct {
		MySQL MySQL
		HTTP  HTTP
	}

	MySQL struct {
		User     string `env:"MYSQL_USER" envDefault:"root"`
		Password string `env:"MYSQL_PASSWORD" envDefault:"root"`
		Database string `env:"MYSQL_DATABASE" envDefault:"library"`
		Port     string `env:"MYSQL_PORT" envDefault:"3306"`
	}

	HTTP struct {
		Address string `env:"SERVER_ADDRESS" envDefault:":8080"`
	}
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func New() *Config {

	err := godotenv.Load(filepath.Join(basepath, "../../.env"))
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables for MYSQL: %e", err)
	}

	return cfg
}
