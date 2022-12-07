package config

type Generate_config struct {
	Common     string `yaml:"Common"`
	AnyfileDir string `yaml:"AnyfileDir"`
	ExportDir  string `yaml:"ExportDir"`
}

type Stream_config struct {
	Enable      bool     `yaml:"Enable"`
	Bind        string   `yaml:"Bind"`
	Stream_Conf []string `yaml:"Stream_conf"`
}

type Config struct {
	Generate map[string]Generate_config `yaml:"Generate"`
	Stream   Stream_config              `yaml:"Stream"`
}
