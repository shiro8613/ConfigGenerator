package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var config = Config{}

func ImportConfig(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return config
}
