package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func Load() (*ServerConfig, error) {
	conf := &ServerConfig{}

	configPath := os.Getenv("CONFIG_PATH")
	_, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadConfig(configPath, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
