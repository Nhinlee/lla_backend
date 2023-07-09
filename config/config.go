package config

type Config struct {
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

func LoadConfig() (*Config, error) {
	return &Config{
		HTTPServerAddress: "0.0.0.0:8080",
	}, nil
}
