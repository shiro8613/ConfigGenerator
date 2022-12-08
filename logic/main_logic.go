package logic

import (
	"fmt"
	"path/filepath"

	"github.com/shiro8613/ConfigGenerator/config"
	"github.com/shiro8613/ConfigGenerator/fileutil"
	"github.com/shiro8613/ConfigGenerator/paser"
)

func MainLogic() error {
	config := config.GetConfig()
	stream_conf_key := make([]string, 0)

	for k, v := range config.Generate {
		stream_conf_key = append(stream_conf_key, k)
		commonY, err := paser.InterfaceYamlPaser(v.Common)
		if err != nil {
			return err
		}

		files, err := fileutil.SearchFiles(v.AnyfileDir)
		if err != nil {
			return err
		}

		for _, file := range files {
			anyFile, _ := paser.InterfaceYamlPaser(file)

			exportName := filepath.Base(file)
			exportPath := v.ExportDir + "/" + exportName

			generatedY := YamlAppender(commonY, anyFile)

			paser.InterfaceYamlExport(generatedY, exportPath)
		}

		fmt.Printf("%s Generated!\n", k)
	}

	if config.Stream.Enable {
		err := StreamServer(stream_conf_key)
		if err != nil {
			return nil
		}
	}

	return nil
}
