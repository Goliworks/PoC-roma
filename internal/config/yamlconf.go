package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type YamlFile struct {
	Http struct {
		Port uint16
		TLS  struct {
			Port         uint16
			Certificates []struct {
				Cert string
				Key  string
			}
		}
	}
	Services map[string]struct {
		Location string
	}
}

var YamlConf YamlFile

func GetYamlConf() *YamlFile {
	yf, _ := ioutil.ReadFile("conf.yml")
	err := yaml.Unmarshal(yf, &YamlConf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(YamlConf)
	return &YamlConf
}
