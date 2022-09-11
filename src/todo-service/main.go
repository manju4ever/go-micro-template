package main

import (
	natsb "github.com/go-micro/plugins/v4/broker/nats"
	natsr "github.com/go-micro/plugins/v4/registry/nats"
	natst "github.com/go-micro/plugins/v4/transport/nats"

	"todo-service/handler"
	pb "todo-service/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "todo-service"
	version = "latest"
)

func main() {
	// Create service
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Broker(natsb.NewBroker()),
		micro.Registry(natsr.NewRegistry()),
		micro.Transport(natst.NewTransport()),
	)
	srv.Init()

	// Register handler
	pb.RegisterTodoServiceHandler(srv.Server(), new(handler.TodoService))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
