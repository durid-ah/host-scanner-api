package config

import (
	"log"

	"go-simpler.org/env"
)

type Config struct {
	NmapAPIHost string `env:"SCANNER_API_HOST" default:"0.0.0.0"`
	NmapAPIPort string `env:"SCANNER_API_PORT" default:"8080"`
	NmapCronTab string `env:"SCANNER_CRON_TAB" default:"*/5 * * * *"`
	NmapTarget  string `env:"SCANNER_TARGET" default:"192.168.1.*"`
}

func NewConfig() *Config {
	cfg := &Config{}
	if err := env.Load(cfg, nil); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return cfg
}
