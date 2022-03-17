package libs

import (
	"flag"

	logger "github.com/accuknox/observability/src/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var log *zerolog.Logger
var configFilePath *string

func init() {
	log = logger.GetInstance()
	configFilePath = flag.String("config-path", "config/", "config/")
	flag.Parse()
}

// LoadConfig - Load the config parameters
func LoadConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic().Msgf("No config file found at %s\n", *configFilePath)
		} else {
			log.Panic().Msgf("Error reading config file: %s\n", readErr)
		}
	}
}
