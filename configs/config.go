package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API    APIConfig
	ExtAPI ExternalConfig
}

type APIConfig struct {
	Port string
}

type ExternalConfig struct {
	URL string
}

func init() {
	viper.SetDefault("api.port", "2000")
	viper.SetDefault("external.url", "https://hp-api.onrender.com/api")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.ExtAPI = ExternalConfig{
		URL: viper.GetString("external.url"),
	}

	return nil
}

func GetBaseURL() string {
	return cfg.ExtAPI.URL
}

func GetServerPort() string {
	return cfg.API.Port
}
