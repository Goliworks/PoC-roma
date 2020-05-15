package config

import (
	"fmt"
)

const (
	defaultHttpPort = 80
)

type Destinations map[string]string

type Config struct {
	Port uint16
	Destinations
	*YamlConf
}

func NewConfig() *Config {
	cfg := new(Config)
	cfg.Port = defaultHttpPort
	cfg.Destinations = make(Destinations)

	cfg.YamlConf = NewYamlConf()
	cfg.generateDestinations()

	// temporary test
	cfg.Destinations["dev1.test"] = "localhost:3500"
	cfg.Destinations["dev2.test"] = "localhost:3000"

	return cfg
}

func (c *Config) generateDestinations(){
	for i, s := range c.Data.Services {
		fmt.Printf("%v & %v\n", i, s)
	}
}
