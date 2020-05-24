package config

import (
	"crypto/tls"
	"fmt"
	"github.com/Goliworks/Roma/internal/utils"
)

const (
	DefaultPort    = ":80"
	DefaultTLSPort = ":443"
)

type DestConf struct {
	Location string
	AutoTLS  bool
}

type Config struct {
	Port         string
	PortTLS      string
	Destinations map[string]*DestConf
	TLSConf      *tls.Config
	AutoTLS      bool
}

func NewConfig() *Config {
	cfg := new(Config)
	cfg.Destinations = make(map[string]*DestConf)
	yc := GetYamlConf()

	if yc.Http.Port != 0 {
		cfg.Port = fmt.Sprintf(":%v", yc.Http.Port)
	} else {
		cfg.Port = DefaultPort
	}
	if yc.Http.TLS.Port != 0 {
		cfg.PortTLS = fmt.Sprintf(":%v", yc.Http.TLS.Port)
	} else {
		cfg.PortTLS = DefaultTLSPort
	}

	cfg.AutoTLS = yc.Http.TLS.Auto
	cfg.generateDestinations(yc)
	cfg.generateCertificates(yc)

	return cfg
}

func (c *Config) generateDestinations(yc *YamlFile) {
	for d, s := range yc.Services {
		dc := &DestConf{Location: s.Location}
		if s.AutoTLS == nil {
			dc.AutoTLS = true
		} else {
			dc.AutoTLS = *s.AutoTLS
		}
		c.Destinations[d] = dc
	}
}

func (c *Config) generateCertificates(yc *YamlFile) {
	c.TLSConf = &tls.Config{}
	for _, i := range yc.Http.TLS.Certificates {
		crt := utils.AbsPath(i.Cert)
		key := utils.AbsPath(i.Key)
		kp, _ := tls.LoadX509KeyPair(crt, key)
		c.TLSConf.Certificates = append(c.TLSConf.Certificates, kp)
	}
}
