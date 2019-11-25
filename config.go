package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type (
	Config struct {
		Countries []ConfigCountry `yaml:"countries"`
	}

	ConfigCountry struct {
		Country string `yaml:"country"`
		Timezone string `yaml:"timezone"`
	}

)

func NewAppConfig(path string) (config Config) {
	var filecontent []byte

	config = Config{}

	Logger.Infof("reading configuration from file %v", path)
	if data, err := ioutil.ReadFile(path); err == nil {
		filecontent = data
	} else {
		panic(err)
	}

	Logger.Info("parsing configuration")
	if err := yaml.Unmarshal(filecontent, &config); err != nil {
		panic(err)
	}

	return
}
