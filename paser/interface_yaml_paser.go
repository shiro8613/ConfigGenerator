package paser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type HogeYaml map[interface{}]interface{}

func InterfaceYamlPaser(path string) (HogeYaml, error) {
	hogeYaml := make(map[interface{}]interface{})

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, hogeYaml)
	if err != nil {
		return nil, err
	}

	return hogeYaml, nil
}
