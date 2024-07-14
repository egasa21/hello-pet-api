package config

import (
	"github.com/egasa21/hello-pet-api/infra/logger"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

func SetupConfig() error {
	var configuration *Configuration
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error to reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Error("error to decode, %v", err)
	}

	return nil
}
