package config

const (
	defaultHttpPort = 80
)

type Destinations map[string]string

type Config struct {
	Port uint16
	Destinations
}

func NewConfig() *Config {
	cfg := new(Config)
	cfg.Port = defaultHttpPort
	cfg.Destinations = make(Destinations)

	// temporary test
	cfg.Destinations["dev1.test"] = "localhost:3500"
	cfg.Destinations["dev2.test"] = "localhost:3000"

	return cfg
}