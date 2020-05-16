package config

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

	return cfg
}

func (c *Config) generateDestinations() {
	for d, s := range c.Data.Services {
		c.Destinations[d] = s.Location
	}
}
