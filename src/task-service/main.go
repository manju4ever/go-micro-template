package main

import (
	"context"
	"fmt"
	"task-service/handler"
	pb "task-service/proto"
	"time"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	natsb "github.com/go-micro/plugins/v4/broker/nats"
	natsr "github.com/go-micro/plugins/v4/registry/nats"
	natst "github.com/go-micro/plugins/v4/transport/nats"
)

var (
	service = "task-service"
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

	// Register handler
	pb.RegisterTaskServiceHandler(srv.Server(), new(handler.TaskService))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

	// Some Wait Time Until The Service Starts !
	time.Sleep(time.Second * 5)

	client := srv.Client()

	req := client.NewRequest("task-service", "TaskService.CreateTodo", &pb.TodoItem{Text: "Read some stuff", Color: "Yellow"})
	res := &pb.Status{}
	if err := client.Call(context.Background(), req, res); err != nil {
		fmt.Println("[error] Something went wrong:", err)
	}

	fmt.Println("End of main()")
}
