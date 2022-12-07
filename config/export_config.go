package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ExportConfig() error {
	export_config := Config{
		Generate: map[string]Generate_config{
			"dir-example": {
				Common:     "/dir/common.yml",
				AnyfileDir: "/dir/any",
				ExportDir:  "/dir/export",
			},
		},

		Stream: Stream_config{
			Enable:      true,
			Bind:        "0.0.0.0:8000",
			Stream_Conf: []string{"dir-example"},
		},
	}

	b, err := yaml.Marshal(export_config)
	if err != nil {
		return err
	}

	file, err := os.Create("./example.yml")
	if err != nil {
		return err
	}

	defer file.Close()

	file.Write(b)

	return nil
}
