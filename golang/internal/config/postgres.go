package config

type postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
}
