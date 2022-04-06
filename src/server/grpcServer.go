package server

import (
	"context"
	"sync"

	"github.com/accuknox/observability/src/aggregator"
	"github.com/accuknox/observability/src/feeds/consumer"
	logger "github.com/accuknox/observability/src/logger"
	agg "github.com/accuknox/observability/src/proto/aggregator"
	cpb "github.com/accuknox/observability/src/proto/consumer"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const PortNumber = "9089"

var log *zerolog.Logger = logger.GetInstance()
var wg sync.WaitGroup

// ======================= //
// == Consumer Service == //
// ===================== //

type consumerServer struct {
	cpb.ConsumerServer
}

func (c *consumerServer) Start(ctx context.Context, in *cpb.ConsumerRequest) (*cpb.ConsumerResponse, error) {
	log.Info().Msg("Start Consumer Called")
	consumer.StartConsumer()

	return &cpb.ConsumerResponse{Result: "Ok"}, nil
}

func (c *consumerServer) Stop(ctx context.Context, in *cpb.ConsumerRequest) (*cpb.ConsumerResponse, error) {
	log.Info().Msg("Stop Consumer Called")
	consumer.StopConsumer()

	return &cpb.ConsumerResponse{Result: "Ok"}, nil
}

// ======================= //
// == Aggregator Service == //
// ===================== //

type aggregatorServer struct {
	agg.AggregatorServer
}

//FetchNetworkLogs -  Service to fetch network logs or count
func (a *aggregatorServer) FetchNetworkLogs(in *agg.NetworkLogsRequest, stream agg.Aggregator_FetchNetworkLogsServer) error {

	if err := aggregator.GetNetworkLogs(in, stream); err != nil {
		return err
	}
	return nil
}

//FetchSystemLogs -  Service to system network logs or count
func (a *aggregatorServer) FetchSystemLogs(in *agg.SystemLogsRequest, stream agg.Aggregator_FetchSystemLogsServer) error {

	if err := aggregator.GetSystemLogs(in, stream); err != nil {
		return err
	}
	return nil
}

// ================= //
// == gRPC Server == //
// ================= //

//GetNewServer - gRPC Server
func GetNewServer() *grpc.Server {

	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	reflection.Register(s)

	//Create Server Instance
	consumerServer := &consumerServer{}
	aggregatorServer := &aggregatorServer{}

	//Register gRPC Server
	cpb.RegisterConsumerServer(s, consumerServer)
	agg.RegisterAggregatorServer(s, aggregatorServer)

	//Start Consumer automatically
	consumer.StartConsumer()

	return s
}
