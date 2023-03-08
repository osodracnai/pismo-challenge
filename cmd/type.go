package cmd

type (
	ConfigServerCmd struct {
		Listen string `mapstructure:"listen"`
		Debug  bool   `mapstructure:"debug"`
	}
	Jaeger struct {
		Name     string
		LogError bool `mapstructure:"log-error"`
		LogInfo  bool `mapstructure:"log-info"`
	}
)
