package consumer

import (
	"fmt"
	"sync"

	"github.com/accuknox/observability/src/feeds/hubble"
	logger "github.com/accuknox/observability/src/logger"
	"github.com/rs/zerolog"
)

const ( // status
	STATUS_RUNNING = "running"
	STATUS_IDLE    = "idle"
)

var Status string
var wg sync.WaitGroup
var stopChan chan struct{}
var log *zerolog.Logger = logger.GetInstance()

func startConsumer() {

	defer wg.Done()

	log.Info().Msgf("Starting consumer")
	//Connect Hubble Relay client
	hubbleClient, err := hubble.GetWatchLogs()
	if err != nil {
		return
	}
	//Connect KubeArmor Relay client
	// kubearmorStream, err := kubearmor.GetWatchLogs()
	// if err != nil {
	// 	return
	// }

	run := true
	for run {
		select {
		case <-stopChan:
			log.Info().Msgf("Got a signal to terminate the consumer")
			run = false
		default:

			hubble.FetchLogs(hubbleClient)
			// kubearmor.FetchLogs(kubearmorStream)

		}

	}
	log.Info().Msgf("Closing consumer")

}

func StartConsumer() {
	fmt.Println("Status in Start Consumer : ", Status)
	if Status == STATUS_RUNNING {
		return
	}

	go startConsumer()
	wg.Add(1)

	stopChan = make(chan struct{})
	Status = STATUS_RUNNING

	log.Info().Msg("Knox observability consumer(s) started")
}

func StopConsumer() {
	log.Info().Msg("Inside Stop Consumer : " + Status)
	if Status != STATUS_RUNNING {
		log.Info().Msg("There is no running consumer(s)")
		return
	}
	Status = STATUS_IDLE
	close(stopChan)

	wg.Wait()
	log.Info().Msg("Knox observability consumer(s) stopped")
}
