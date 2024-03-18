package config

import "net"

type server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s *server) Adders() string {
	return net.JoinHostPort(s.Host, s.Port)
}
