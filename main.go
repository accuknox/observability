package main

import (
	"net"

	libs "github.com/accuknox/observability/src/libs"
	logger "github.com/accuknox/observability/src/logger"
	grpcserver "github.com/accuknox/observability/src/server"
	"github.com/accuknox/observability/utils/database"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// var wg sync.WaitGroup
var log *zerolog.Logger = logger.GetInstance()

func init() {
	libs.LoadConfig()
	logger.SetLogLevel(viper.GetString("logging.level"))
	// log = logger.GetInstance()
	database.ConnectDB()
}

func main() {

	// wg.Add(1)
	// go kubearmor.GetWatchLogs()
	// go hubble.GetWatchLogs()

	// wg.Wait()
	// cmd.Execute()

	//Create Server
	listen, err := net.Listen("tcp", ":"+grpcserver.PortNumber)
	if err != nil {
		log.Error().Msgf("gRPC server failed to listen : %v", err)
	}

	server := grpcserver.GetNewServer()
	//Start service
	log.Info().Msgf("gRPC server on %s port started", grpcserver.PortNumber)
	if err := server.Serve(listen); err != nil {
		log.Error().Msgf("Failed to serve: %v", err)
	}
}
