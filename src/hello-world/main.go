package main

import (
	"context"
	"fmt"

	micro "go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	natsb "github.com/go-micro/plugins/v4/broker/nats"
	natsr "github.com/go-micro/plugins/v4/registry/nats"
	natst "github.com/go-micro/plugins/v4/transport/nats"

	"hello-world/handler"
	pb "hello-world/proto"
)

var (
	service = "hello-world"
	version = "latest"
)

func main() {

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Broker(natsb.NewBroker()),
		micro.Registry(natsr.NewRegistry()),
		micro.Transport(natst.NewTransport()),
	)
	srv.Init()

	todoClient := srv.Client()
	req := todoClient.NewRequest("todo-service", "TodoService.Call", &pb.CallRequest{Name: "Manju"})
	res := &pb.CallResponse{}
	ctx := context.Background()
	if err := todoClient.Call(ctx, req, res); err != nil {
		fmt.Println("Something is wrong:", err, res)
		return
	}
	fmt.Println(res)

	// Register handler
	pb.RegisterHelloWorldHandler(srv.Server(), new(handler.HelloWorld))
	// Run service
	if err := srv.Run(); err != nil {
		log.Debugf("[main] Something went wrong")
		log.Fatal(err)
	}
}
