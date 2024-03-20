package config

import (
	"fmt"
)

const DefaultPostgresPort = 5432

type postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
}

func (p *postgres) fillDefault() {
	if p.Port == 0 {
		p.Port = DefaultPostgresPort
	}
}

func (p *postgres) Driver() (string, string) {
	return "postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.Database)
}
