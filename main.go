package main

import (
	"log"
	"os"

	"github.com/shiro8613/ConfigGenerator/config"
	"github.com/shiro8613/ConfigGenerator/logic"
	"github.com/shiro8613/ConfigGenerator/paser"
)

func main() {
	flag_data := paser.FlagPaser()
	if flag_data.Config_example {
		log.Fatal(config.ExportConfig())
		os.Exit(0)
	}

	err := config.ImportConfig(flag_data.Config_path)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(logic.MainLogic())

}
