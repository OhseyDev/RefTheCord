package lib

import (
	"log"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	DBType string `env:"DB_TYPE" envDefault:"sqlite3"`
	Token  string `env:"TOKEN"`
}

func NewConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%v\n", err)
	}
	return cfg
}
