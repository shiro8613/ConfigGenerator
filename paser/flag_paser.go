package paser

import "flag"

type FlagData struct {
	Config_example bool
	Config_path    string
}

func FlagPaser() FlagData {
	var (
		config_example = flag.Bool("config_example", false, "export Config Example")
		config_path    = flag.String("config", "./config.yml", "Config File path")
	)

	flag.Parse()

	return FlagData{
		Config_example: *config_example,
		Config_path:    *config_path,
	}
}
