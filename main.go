package main

import (
	"flag"
	"sync"

	"github.com/accuknox/observability/src/cmd"
	"github.com/accuknox/observability/utils/database"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var configFilePath *string
var wg sync.WaitGroup

func main() {
	configFilePath = flag.String("config-path", "config/", "config/")
	flag.Parse()

	loadConfig()
	setDatabase()

	// wg.Add(1)
	// go kubearmor.GetWatchLogs()
	// go hubble.GetWatchLogs()

	// wg.Wait()
	cmd.Execute()

}

// loadConfig - Load the config parameters
func loadConfig() {
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

func setDatabase() {
	database.ConnectDB()
}
