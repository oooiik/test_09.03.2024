package config

import (
	"fmt"
	"net/url"
)

const DefaultClickhousePort = 8123

type clickhouse struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
}

func (c *clickhouse) fillDefault() {
	if c.Port == 0 {
		c.Port = DefaultClickhousePort
	}
}

func (c *clickhouse) Driver() (string, string) {
	q := make(url.Values)
	q.Set("username", c.Username)
	q.Set("password", c.Password)
	q.Set("database", c.Database)
	dsn := (&url.URL{
		Scheme:   "tcp",
		Host:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		RawQuery: q.Encode(),
	}).String()
	return "clickhouse", dsn
}
