package config

import "fmt"

const (
	defaultHttpPort = ":80"
)

type Destinations map[string]string

type Config struct {
	Port string
	Destinations
}

func NewConfig() *Config {
	cfg := new(Config)
	cfg.Destinations = make(Destinations)

	yc := NewYamlConf()
	if yc.Http.Port != 0 {
		cfg.Port = fmt.Sprintf(":%v", yc.Http.Port)
	} else {
		cfg.Port = defaultHttpPort
	}
	cfg.generateDestinations(yc)

	return cfg
}

func (c *Config) generateDestinations(yc *YamlFile) {
	for d, s := range yc.Services {
		c.Destinations[d] = s.Location
	}
}
