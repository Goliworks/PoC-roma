package config

const (
	defaultHttpPort = 80
)

type Config struct {
	Port uint16
}

func NewConfig() *Config {
	return &Config{defaultHttpPort}
}