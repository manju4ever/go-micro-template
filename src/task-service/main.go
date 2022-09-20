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

	client := pb.NewTaskService("task-service", srv.Client())

	//Create a New Todo
	client.CreateTodo(context.Background(), &pb.TodoItem{Text: "Go Workout in Gym", Color: "Green"})

	// Get All Todos
	res, err := client.GetAllTodos(context.Background(), &pb.Void{})
	if err != nil {
		fmt.Println("[error] Something went wrong:", err)
	}
	fmt.Println("Total Items In Database:", len(res.Items))
}
