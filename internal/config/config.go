package config

import (
	"crypto/tls"
	"fmt"
	"path/filepath"
)

const (
	defaultPort    = ":80"
	defaultTLSPort = ":443"
)

type Destinations map[string]string

type Config struct {
	Port    string
	PortTLS string
	Destinations
	TLSConf *tls.Config
}

func NewConfig() *Config {
	cfg := new(Config)
	cfg.Destinations = make(Destinations)

	yc := GetYamlConf()

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
	cfg.generateCertificates(yc)

	return cfg
}

func (c *Config) generateDestinations(yc *YamlFile) {
	for d, s := range yc.Services {
		c.Destinations[d] = s.Location
	}
}

func (c *Config) generateCertificates(yc *YamlFile) {
	c.TLSConf = &tls.Config{}
	for _, i := range yc.Http.TLS.Certificates {
		crt, _ := filepath.Abs(i.Cert)
		key, _ := filepath.Abs(i.Key)
		kp, _ := tls.LoadX509KeyPair(crt, key)
		c.TLSConf.Certificates = append(c.TLSConf.Certificates, kp)
	}
}
