package config

import (
	"gopkg.in/yaml.v3"
	log2 "log"
	"os"
)

const defaultConfFile = "./config.yml"

type Config struct {
	Server     server     `yaml:"server"`
	Postgres   postgres   `yaml:"postgres"`
	Clickhouse clickhouse `yaml:"clickhouse"`
	Log        log        `yaml:"log"`
}

var singleton *Config

func Load() *Config {
	if singleton == nil {
		file, err := os.Open(defaultConfFile)
		if err != nil {
			log2.Fatal(err)
		}
		defer file.Close()

		var cfg Config
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&cfg)
		if err != nil {
			log2.Fatal(err)
		}

		cfg.Postgres.fillDefault()
		cfg.Log.fillDefault()

		singleton = &cfg
	}
	return singleton
}
