package paser

import (
	"os"

	"gopkg.in/yaml.v3"
)

func InterfaceYamlExport(data HogeYaml, path string) error {
	b, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	file.Write(b)

	return nil
}
