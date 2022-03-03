package server

import (
	"context"
	"sync"

	"github.com/accuknox/observability/src/feeds/consumer"
	logger "github.com/accuknox/observability/src/logger"
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

	//Register gRPC Server
	cpb.RegisterConsumerServer(s, consumerServer)

	//Start Consumer automatically
	consumer.StartConsumer()

	return s
}
