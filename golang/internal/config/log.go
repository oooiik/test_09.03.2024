package config

import (
	"io"
	"os"
)

const (
	DefaultLogLevel      = "info"    // error, warning, info, debug
	DefaultLogWriterType = "console" // console, file
	DefaultLogPath       = "./log/app.log"
)

type log struct {
	Level      string `yaml:"level"`
	WriterType string `yaml:"writer"`
	Path       string `yaml:"path"`
}

func (c *log) fillDefault() {
	if c.Level == "" {
		c.Level = DefaultLogLevel
	}
	if c.WriterType == "" {
		c.WriterType = DefaultLogWriterType
	}
	if c.Path == "" {
		if c.WriterType == "file" {
			c.Path = DefaultLogPath
		}
	}
}

func (c *log) Writer() io.Writer {
	switch c.WriterType {
	case "file":
		file, err := os.OpenFile(c.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {

		}
		defer file.Close()

		return file
	case "console":
	}
	return os.Stdout
}
