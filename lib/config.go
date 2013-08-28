package lib

import (
	"io/ioutil"
	"launchpad.net/goyaml"
)

type ConfigValue struct {
	Listen   string
	Mongodb  string
	Dbname   string
	Loglevel int
}

var Config ConfigValue

func init() {
	LoadConfig("./config.yml")
}

func LoadConfig(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	goyaml.Unmarshal(data, &Config)
}
