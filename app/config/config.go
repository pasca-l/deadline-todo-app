package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"

	"app/utils"
)

type Server struct {
	Port    string
	LogFile string
	Static  string
}

type Database struct {
	Driver string
	Name   string
}

type CustomConfig struct {
	Srv Server
	Db  Database
}

var Config CustomConfig

func init() {
	LoadConfigYaml()
	utils.LoggingConfig(Config.Srv.LogFile)
}

func LoadConfigIni() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	cfg.MapTo(&Config)
	// or use `cfg.Section("").Key("").String()` notation,
	// to create a struct with designated fields
}

func LoadConfigYaml() {
	f, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	if err := yaml.Unmarshal(f, &Config); err != nil {
		log.Fatalln(err)
	}
}
