package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Get(filePath string) (Config, error) {
	var config Config

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		return config, err
	}

	return config, err
}
