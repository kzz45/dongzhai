package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Name  string `yaml:"name"`
	Debug bool   `yaml:"debug"`
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var GlobalConfig Config

func init() {
	yaml_file, err := ioutil.ReadFile("./server.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}
	if err := yaml.Unmarshal(yaml_file, &GlobalConfig); err != nil {
		log.Fatalln(err)
		return
	}
}
