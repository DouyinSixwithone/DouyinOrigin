package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Mysql struct {
	User      string `yaml:"user"`
	Pass      string `yaml:"pass"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Dbname    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Yaml2Go struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

const confFile = "config/config.yaml"

var Conf Yaml2Go

func Init() error {
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, &Conf); err != nil {
		return err
	}
	return nil
}