package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type yamlFile struct {
	Services map[string]struct {
		Location string
	}
}

type YamlConf struct {
	Data yamlFile
}

func NewYamlConf() *YamlConf {
	yc := new(YamlConf)
	yc.Data = yamlFile{}
	yf, _ := ioutil.ReadFile("conf.yml")
	err := yaml.Unmarshal(yf, &yc.Data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(yc.Data)
	return yc
}
