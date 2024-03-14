package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server     server     `yaml:"server"`
	Postgres   postgres   `yaml:"postgres"`
	Clickhouse clickhouse `yaml:"clickhouse"`
}

var singleton *Config

func New() *Config {
	if singleton == nil {
		file, err := os.Open("config.yml")
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)

		var cfg Config
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(cfg)
		if err != nil {
			log.Fatal(err)
		}

		singleton = &cfg
	}
	return singleton
}
