package config

import "fmt"

const (
	defaultPort    = ":80"
	defaultTLSPort = ":443"
)

type Destinations map[string]string

type Config struct {
	Port    string
	PortTLS string
	Destinations
}

func NewConfig() *Config {
	cfg := new(Config)
	cfg.Destinations = make(Destinations)

	yc := NewYamlConf()

	if yc.Http.Port != 0 {
		cfg.Port = fmt.Sprintf(":%v", yc.Http.Port)
	} else {
		cfg.Port = defaultPort
	}
	if yc.Http.TLS.Port != 0 {
		cfg.PortTLS = fmt.Sprintf(":%v", yc.Http.TLS.Port)
	} else {
		cfg.PortTLS = defaultTLSPort
	}

	cfg.generateDestinations(yc)

	return cfg
}

func (c *Config) generateDestinations(yc *YamlFile) {
	for d, s := range yc.Services {
		c.Destinations[d] = s.Location
	}
}
